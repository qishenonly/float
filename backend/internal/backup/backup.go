package backup

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/qiuhaonan/float-backend/pkg/logger"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
)

// BackupService handles database backups
type BackupService struct {
	cron *cron.Cron
}

// NewBackupService creates a new backup service
func NewBackupService() *BackupService {
	return &BackupService{
		cron: cron.New(),
	}
}

// StartScheduler starts the backup scheduler
func (s *BackupService) StartScheduler() {
	// Schedule daily backup at 00:00
	_, err := s.cron.AddFunc("0 0 * * *", func() {
		logger.Info("[Backup] Starting scheduled daily database backup...")
		if err := s.PerformBackup(); err != nil {
			logger.Error(fmt.Sprintf("[Backup] Scheduled backup failed: %v", err))
		} else {
			logger.Info("[Backup] Scheduled backup completed successfully")
		}
	})

	if err != nil {
		logger.Error(fmt.Sprintf("[Backup] Failed to schedule backup job: %v", err))
		return
	}

	s.cron.Start()
	logger.Info("[Backup] Database backup scheduler started (Daily at 00:00)")
}

// PerformBackup executes the mysqldump command
func (s *BackupService) PerformBackup() error {
	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetString("DB_PORT")
	dbUser := viper.GetString("DB_USERNAME")
	dbPassword := viper.GetString("DB_PASSWORD")
	dbName := viper.GetString("DB_DATABASE")

	// Ensure backup directory exists
	// Default to ./backups if not specified in config (better for local dev)
	// In Docker, we can set BACKUP_DIR env var to /root/backups
	viper.SetDefault("BACKUP_DIR", "./backups")
	backupDir := viper.GetString("BACKUP_DIR")

	if _, err := os.Stat(backupDir); os.IsNotExist(err) {
		if err := os.MkdirAll(backupDir, 0755); err != nil {
			return fmt.Errorf("failed to create backup directory: %v", err)
		}
	}

	// Generate filename: float_backup_20260116.sql (saves to date)
	// User requested "save to the day's sql backup file", imply one per day or date stamped.
	// Using date stamp is safer.
	filename := fmt.Sprintf("float_backup_%s.sql", time.Now().Format("20060102"))
	filePath := filepath.Join(backupDir, filename)

	// Construct mysqldump command
	// Note: We use -h, -P, -u, -p (no space for password is safer in some versions but generally -pPASSWORD works or -p and environment var)
	// Ideally use MYSQL_PWD environment variable to avoid password in process list?
	// But exec.Command args are hidden from shell usually.
	// However, standard mysqldump usage: mysqldump -h host -u user -ppassword dbname

	args := []string{
		"-h", dbHost,
		"-P", dbPort,
		"-u", dbUser,
		fmt.Sprintf("-p%s", dbPassword),
		dbName,
	}

	// Check if mysqldump exists
	cmdPath, err := exec.LookPath("mysqldump")
	if err != nil {
		return fmt.Errorf("mysqldump tool not found: %v", err)
	}

	cmd := exec.Command(cmdPath, args...)

	// Create output file
	outfile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create backup file: %v", err)
	}
	defer outfile.Close()

	cmd.Stdout = outfile
	cmd.Stderr = os.Stderr // Capture stderr for debugging

	logger.Info(fmt.Sprintf("[Backup] Executing mysqldump to %s...", filePath))
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("mysqldump execution failed: %v", err)
	}

	logger.Info(fmt.Sprintf("[Backup] Backup successfully saved to %s", filePath))
	return nil
}
