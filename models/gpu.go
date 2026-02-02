package models

type GPU struct {
	Name           string `json:"name"`
	AdapterRAM     uint32 `json:"adapter_ram"`
	DriverVersion  string `json:"driver_version"`
	VideoProcessor string `json:"video_processor"`
}
