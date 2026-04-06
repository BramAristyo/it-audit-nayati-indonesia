package models

import "time"

type Win32DiskDrive struct {
	MediaType string
	Model     string
}

type Disk struct {
	Drive   string  `json:"drive"`
	TotalGB float64 `json:"total_gb"`
	Type    string  `json:"type"`
}

type DiskModel struct {
	ID         uint `gorm:"primaryKey"`
	ComputerID uint
	Drive      string
	TotalGB    float64 `gorm:"column:total_size_gb"`
	Type       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (DiskModel) TableName() string {
	return "storage_devices"
}
