package models

import (
	json2 "encoding/json"
	"fmt"
	"time"
)

type SpecResponse struct {
	Hostname     string `json:"host_name"`
	MachineID    string `json:"machine_id"`
	IPAddress    string `json:"ip_address"`
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`

	OS      OS     `json:"os"`
	CPU     CPU    `json:"cpu"`
	Memory  Memory `json:"memory"`
	GPU     []GPU  `json:"gpus"`
	Storage []Disk `json:"storage"`

	CollectedAt time.Time `json:"collected_at"`
}

func (s SpecResponse) ToComputer() Computer {
	return Computer{
		HostName:     s.Hostname,
		MachineID:    s.MachineID,
		IPAddress:    s.IPAddress,
		Manufacturer: s.Manufacturer,
		ModelName:    s.Model,
	}
}

func (s SpecResponse) ToComputerSpec() ComputerSpec {
	return ComputerSpec{
		OSName:         s.OS.Name,
		OSVersion:      s.OS.Version,
		OSArchitecture: s.OS.Architecture,
		CPUModel:       s.CPU.Model,
		CPUCores:       s.CPU.Cores,
		CPUThreads:     s.CPU.Threads,
		Frequency:      s.CPU.Frequency,
		RAMGb:          s.Memory.TotalGB,
		CollectedAt:    time.Time{},
	}
}

func (s SpecResponse) ToGpus() []GPUModel {
	var gpus []GPUModel

	for i, g := range s.GPU {
		gpus = append(gpus, GPUModel{
			Name:           g.Name,
			AdapterRAM:     g.AdapterRAM,
			DriverVersion:  g.DriverVersion,
			VideoProcessor: g.VideoProcessor,
			GPUIndex:       i,
		})
	}

	return gpus
}

func (s SpecResponse) ToDisks() []DiskModel {
	var disks []DiskModel

	for _, d := range s.Storage {
		disks = append(disks, DiskModel{
			Drive:     d.Drive,
			TotalGB:   d.TotalGB,
			Type:      d.Type,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		})
	}

	return disks
}

func (s SpecResponse) PrintCollectedData() {
	json, _ := json2.MarshalIndent(s, "", "  ")
	fmt.Println(string(json))
	fmt.Println()
}

type ComputerSpec struct {
	ID         uint `gorm:"primaryKey"`
	ComputerID uint `gorm:"uniqueIndex"`

	OSName         string
	OSVersion      string
	OSArchitecture string

	CPUModel   string
	CPUCores   int
	CPUThreads int
	Frequency  float64 `gorm:"column:frequency_mhz"`

	RAMGb       float64 `gorm:"column:ram_gb"`
	CollectedAt time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (ComputerSpec) TableName() string {
	return "computer_specs"
}
