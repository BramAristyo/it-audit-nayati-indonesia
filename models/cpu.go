package models

type CPU struct {
	Model     string  `json:"model"`
	Cores     int     `json:"cores"`
	Threads   int     `json:"threads"`
	Frequency float64 `json:"frequency_mhz"`
}
