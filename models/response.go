package models

import "time"

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
