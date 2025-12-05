package service

import (
	"errors"

	"github.com/qiuhaonan/float-backend/internal/dto/response"
	"github.com/qiuhaonan/float-backend/internal/models"
	"github.com/qiuhaonan/float-backend/internal/repository"
	"gorm.io/gorm"
)

type AppUpdateService struct {
	repo *repository.AppUpdateRepository
}

func NewAppUpdateService(repo *repository.AppUpdateRepository) *AppUpdateService {
	return &AppUpdateService{repo: repo}
}

// CheckUpdate 检查更新
func (s *AppUpdateService) CheckUpdate(platform string, currentVersionCode int) (*response.CheckUpdateResponse, error) {
	latest, err := s.repo.GetLatestVersion(platform)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &response.CheckUpdateResponse{HasUpdate: false}, nil
		}
		return nil, err
	}

	if latest.VersionCode > currentVersionCode {
		resp := &response.CheckUpdateResponse{
			HasUpdate: true,
			Latest: &response.AppUpdateResponse{
				ID:                  latest.ID,
				VersionCode:         latest.VersionCode,
				VersionName:         latest.VersionName,
				Platform:            latest.Platform,
				UpdateType:          latest.UpdateType,
				IsForceUpdate:       latest.IsForceUpdate,
				MinSupportedVersion: latest.MinSupportedVersion,
				Title:               latest.Title,
				Description:         latest.Description,
				Changelog:           latest.Changelog,
				DownloadURL:         latest.DownloadURL,
				FileSize:            latest.FileSize,
				FileHash:            latest.FileHash,
				ReleaseNotesURL:     latest.ReleaseNotesURL,
				ReleaseDate:         latest.ReleaseDate,
			},
			ForceUpdate: latest.IsForceUpdate,
		}

		// 如果当前版本低于最低支持版本，强制更新
		// 这里简单处理，实际可能需要解析版本号字符串比较
		if latest.IsForceUpdate {
			resp.UpdateReason = "当前版本过低，请更新到最新版本以继续使用"
		} else {
			resp.UpdateReason = "发现新版本，建议更新"
		}

		return resp, nil
	}

	return &response.CheckUpdateResponse{HasUpdate: false}, nil
}

// GetLatest 获取最新版本
func (s *AppUpdateService) GetLatest(platform string) (*response.AppUpdateResponse, error) {
	latest, err := s.repo.GetLatestVersion(platform)
	if err != nil {
		return nil, err
	}

	return &response.AppUpdateResponse{
		ID:                  latest.ID,
		VersionCode:         latest.VersionCode,
		VersionName:         latest.VersionName,
		Platform:            latest.Platform,
		UpdateType:          latest.UpdateType,
		IsForceUpdate:       latest.IsForceUpdate,
		MinSupportedVersion: latest.MinSupportedVersion,
		Title:               latest.Title,
		Description:         latest.Description,
		Changelog:           latest.Changelog,
		DownloadURL:         latest.DownloadURL,
		FileSize:            latest.FileSize,
		FileHash:            latest.FileHash,
		ReleaseNotesURL:     latest.ReleaseNotesURL,
		ReleaseDate:         latest.ReleaseDate,
	}, nil
}

// GetHistory 获取更新历史
func (s *AppUpdateService) GetHistory(platform string) ([]response.AppUpdateResponse, error) {
	updates, err := s.repo.GetUpdateHistory(platform, 20) // 默认取最近20条
	if err != nil {
		return nil, err
	}

	var res []response.AppUpdateResponse
	for _, u := range updates {
		res = append(res, response.AppUpdateResponse{
			ID:                  u.ID,
			VersionCode:         u.VersionCode,
			VersionName:         u.VersionName,
			Platform:            u.Platform,
			UpdateType:          u.UpdateType,
			IsForceUpdate:       u.IsForceUpdate,
			MinSupportedVersion: u.MinSupportedVersion,
			Title:               u.Title,
			Description:         u.Description,
			Changelog:           u.Changelog,
			DownloadURL:         u.DownloadURL,
			FileSize:            u.FileSize,
			FileHash:            u.FileHash,
			ReleaseNotesURL:     u.ReleaseNotesURL,
			ReleaseDate:         u.ReleaseDate,
		})
	}
	return res, nil
}

// CreateAppUpdate 创建更新
func (s *AppUpdateService) CreateAppUpdate(update *models.AppUpdate) error {
	return s.repo.Create(update)
}
