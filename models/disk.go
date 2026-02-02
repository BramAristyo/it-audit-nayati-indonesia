package models

type Win32DiskDrive struct {
	MediaType string
	Model     string
}

type Disk struct {
	Drive   string  `json:"drive"`
	TotalGB float64 `json:"total_gb"`
	Type    string  `json:"type"`
}
