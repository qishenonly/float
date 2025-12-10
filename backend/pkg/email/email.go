package email

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"math/big"
	"net"
	"net/smtp"
	"strings"
	"time"

	"github.com/qiuhaonan/float-backend/pkg/cache"
	"github.com/qiuhaonan/float-backend/pkg/logger"
	"github.com/spf13/viper"
)

// Email 邮件服务接口
type Email interface {
	SendVerificationCode(email string, code string) error
	SendPasswordReset(email string, resetToken string) error
	SendWelcome(email string, displayName string) error
}

type emailService struct {
	smtpHost     string
	smtpPort     string
	fromEmail    string
	fromPassword string
	enabled      bool
}

var (
	defaultService Email
)

// Init 初始化邮件服务
func Init() error {
	service, err := newEmailService()
	if err != nil {
		logger.Warn("邮件服务初始化失败: ", err)
		return err
	}
	defaultService = service

	if defaultService.(*emailService).enabled {
		logger.Info("邮件服务已启用")
	} else {
		logger.Info("邮件服务未启用，使用测试模式（验证码将在控制台输出）")
	}

	return nil
}

// newEmailService 创建邮件服务实例
func newEmailService() (Email, error) {
	enabled := viper.GetBool("email.enabled")
	host := viper.GetString("email.smtp.host")
	port := viper.GetString("email.smtp.port")
	username := viper.GetString("email.smtp.username")
	password := viper.GetString("email.smtp.password")

	if !enabled || host == "" || username == "" || password == "" {
		// 如果未配置邮件服务，返回 mock service
		return &emailService{
			enabled: false,
		}, nil
	}

	return &emailService{
		smtpHost:     host,
		smtpPort:     port,
		fromEmail:    username,
		fromPassword: password,
		enabled:      true,
	}, nil
}

// GetService 获取邮件服务实例
func GetService() Email {
	if defaultService == nil {
		defaultService = &emailService{}
	}
	return defaultService
}

// SendVerificationCode 发送验证码邮件
func (s *emailService) SendVerificationCode(toEmail string, code string) error {
	if !s.enabled {
		// 未配置邮件服务，仅记录日志
		logger.Warn(fmt.Sprintf("[邮件验证码] 邮件服务未启用（测试模式） | 收件人: %s | 验证码: %s", toEmail, code))
		fmt.Printf("【邮件服务未配置】发送验证码到 %s: %s\n", toEmail, code)
		return nil
	}

	subject := "Float Island - 邮箱验证码"
	body := strings.Replace(VerificationCodeTemplate, "{{CODE}}", code, 1)

	logger.Info(fmt.Sprintf("[邮件验证码] 发送验证码-%s | 收件人: %s", code, toEmail))
	return s.sendEmail(toEmail, subject, body)
}

// SendPasswordReset 发送密码重置邮件
func (s *emailService) SendPasswordReset(toEmail string, resetToken string) error {
	if !s.enabled {
		fmt.Printf("【邮件服务未配置】发送密码重置链接到 %s\n", toEmail)
		return nil
	}

	subject := "Float Island - 重置密码"
	resetLink := fmt.Sprintf("https://floatisland.app/reset-password?token=%s", resetToken)
	body := strings.Replace(PasswordResetTemplate, "{{RESET_LINK}}", resetLink, -1)

	logger.Info(fmt.Sprintf("[邮件密码重置] 发送重置链接 | 收件人: %s", toEmail))
	return s.sendEmail(toEmail, subject, body)
}

// SendWelcome 发送欢迎邮件
func (s *emailService) SendWelcome(toEmail string, displayName string) error {
	if !s.enabled {
		fmt.Printf("【邮件服务未配置】发送欢迎邮件到 %s (收件人: %s)\n", toEmail, displayName)
		return nil
	}

	subject := "Float Island - 欢迎加入！"
	body := strings.Replace(WelcomeTemplate, "{{DISPLAY_NAME}}", displayName, 1)

	logger.Info(fmt.Sprintf("[邮件欢迎] 发送欢迎邮件 | 收件人: %s | 用户: %s", toEmail, displayName))
	return s.sendEmail(toEmail, subject, body)
}

// sendEmail 发送邮件 - 使用 TLS 加密连接
func (s *emailService) sendEmail(toEmail, subject, body string) error {
	if !s.enabled {
		return nil
	}

	// SMTP 连接地址
	addr := fmt.Sprintf("%s:%s", s.smtpHost, s.smtpPort)

	// 第一步：建立TCP连接
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		logger.Error(fmt.Sprintf("[邮件服务] TCP连接失败 | 地址: %s | 错误: %v", addr, err))
		return fmt.Errorf("TCP连接失败: %w", err)
	}
	defer conn.Close()

	// 第二步：升级为TLS连接
	tlsConfig := &tls.Config{
		ServerName:         s.smtpHost,
		InsecureSkipVerify: false,
	}

	tlsConn := tls.Client(conn, tlsConfig)
	if err := tlsConn.Handshake(); err != nil {
		logger.Error(fmt.Sprintf("[邮件服务] TLS握手失败 | 错误: %v", err))
		return fmt.Errorf("TLS握手失败: %w", err)
	}
	defer tlsConn.Close()

	// 第三步：创建SMTP客户端
	client, err := smtp.NewClient(tlsConn, s.smtpHost)
	if err != nil {
		logger.Error(fmt.Sprintf("[邮件服务] 创建SMTP客户端失败 | 错误: %v", err))
		return fmt.Errorf("创建SMTP客户端失败: %w", err)
	}
	defer client.Close()

	// 第四步：SMTP认证
	auth := smtp.PlainAuth("", s.fromEmail, s.fromPassword, s.smtpHost)
	if err := client.Auth(auth); err != nil {
		logger.Error(fmt.Sprintf("[邮件服务] SMTP认证失败 | 用户: %s | 错误: %v", s.fromEmail, err))
		return fmt.Errorf("SMTP认证失败: %w", err)
	}

	// 第五步：设置发件人
	if err := client.Mail(s.fromEmail); err != nil {
		logger.Error(fmt.Sprintf("[邮件服务] 设置发件人失败 | 错误: %v", err))
		return fmt.Errorf("设置发件人失败: %w", err)
	}

	// 第六步：设置收件人
	if err := client.Rcpt(toEmail); err != nil {
		logger.Error(fmt.Sprintf("[邮件服务] 设置收件人失败 | 收件人: %s | 错误: %v", toEmail, err))
		return fmt.Errorf("设置收件人失败: %w", err)
	}

	// 第七步：发送邮件内容
	wc, err := client.Data()
	if err != nil {
		logger.Error(fmt.Sprintf("[邮件服务] 获取邮件数据写入器失败 | 错误: %v", err))
		return fmt.Errorf("获取邮件数据写入器失败: %w", err)
	}

	// 生成邮件数据
	emailData := []byte(fmt.Sprintf("Subject: %s\r\nContent-Type: text/html; charset=UTF-8\r\nFrom: %s\r\nTo: %s\r\n\r\n%s",
		subject, s.fromEmail, toEmail, body))

	_, err = wc.Write(emailData)
	if err != nil {
		logger.Error(fmt.Sprintf("[邮件服务] 写入邮件数据失败 | 错误: %v", err))
		return fmt.Errorf("写入邮件数据失败: %w", err)
	}

	if err := wc.Close(); err != nil {
		logger.Error(fmt.Sprintf("[邮件服务] 关闭邮件数据写入器失败 | 错误: %v", err))
		return fmt.Errorf("关闭邮件数据写入器失败: %w", err)
	}

	// 第八步：关闭SMTP连接
	if err := client.Quit(); err != nil {
		logger.Error(fmt.Sprintf("[邮件服务] 关闭SMTP连接失败 | 错误: %v", err))
		return fmt.Errorf("关闭SMTP连接失败: %w", err)
	}

	logger.Info(fmt.Sprintf("[邮件服务] 邮件发送成功 | 收件人: %s | 主题: %s", toEmail, subject))
	return nil
}

// GenerateVerificationCode 生成验证码
func GenerateVerificationCode() (string, error) {
	// 生成 6 位随机数字验证码
	code := ""
	for i := 0; i < 6; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		code += fmt.Sprint(num.Int64())
	}
	return code, nil
}

// StoreVerificationCode 存储验证码到 Redis (TTL: 10分钟)
func StoreVerificationCode(email string, code string) error {
	redisKey := fmt.Sprintf("email_verification:%s", email)
	expiry := GetVerificationCodeExpiry()

	return cache.Set(redisKey, code, expiry)
}

// VerifyCode 验证验证码并删除
func VerifyCode(email string, code string) (bool, error) {
	redisKey := fmt.Sprintf("email_verification:%s", email)
	storedCode, err := cache.Get(redisKey)
	if err != nil {
		// 验证码不存在或已过期
		return false, nil
	}

	if storedCode != code {
		return false, nil
	}

	// 验证成功后删除验证码
	_ = cache.Del(redisKey)
	return true, nil
}

// GetVerificationCodeExpiry 获取验证码过期时间（分钟）
func GetVerificationCodeExpiry() time.Duration {
	expiry := viper.GetInt("email.verification_code_expiry")
	if expiry <= 0 {
		expiry = 10 // 默认 10 分钟
	}
	return time.Duration(expiry) * time.Minute
}
