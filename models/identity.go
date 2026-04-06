package models

import "time"

type Computer struct {
	ID           uint `gorm:"primaryKey"`
	HostName     string
	MachineID    string
	IPAddress    string
	Manufacturer string
	ModelName    string `gorm:"column:model"`
	UpdatedAt    time.Time
	CreatedAt    time.Time
}

func (Computer) TableName() string {
	return "computers"
}
