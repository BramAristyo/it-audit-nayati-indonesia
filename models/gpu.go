package models

import "time"

type GPU struct {
	Name           string `json:"name"`
	AdapterRAM     uint32 `json:"adapter_ram"`
	DriverVersion  string `json:"driver_version"`
	VideoProcessor string `json:"video_processor"`
}

type GPUModel struct {
	ID             uint `gorm:"primaryKey"`
	ComputerID     uint
	Name           string
	AdapterRAM     uint32
	DriverVersion  string
	VideoProcessor string
	GPUIndex       int `gorm:"column:index"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (GPUModel) TableName() string {
	return "gpus"
}
