package response

import (
	"encoding/json"
	"time"
)

// AppUpdateResponse 软件更新响应
type AppUpdateResponse struct {
	ID                  int64           `json:"id"`
	VersionCode         int             `json:"version_code"`
	VersionName         string          `json:"version_name"`
	Platform            string          `json:"platform"`
	UpdateType          string          `json:"update_type"`
	IsForceUpdate       bool            `json:"is_force_update"`
	MinSupportedVersion string          `json:"min_supported_version"`
	Title               string          `json:"title"`
	Description         string          `json:"description"`
	Changelog           json.RawMessage `json:"changelog"`
	DownloadURL         string          `json:"download_url"`
	FileSize            int64           `json:"file_size"`
	FileHash            string          `json:"file_hash"`
	ReleaseNotesURL     string          `json:"release_notes_url"`
	ReleaseDate         *time.Time      `json:"release_date"`
}

// CheckUpdateResponse 检查更新响应
type CheckUpdateResponse struct {
	HasUpdate    bool               `json:"has_update"`
	Latest       *AppUpdateResponse `json:"latest,omitempty"`
	ForceUpdate  bool               `json:"force_update"`
	UpdateReason string             `json:"update_reason,omitempty"`
}
