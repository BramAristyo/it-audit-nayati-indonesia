package models

type Monitor struct {
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	Resolution   string `json:"resolution"`
	ScreenWidth  uint32 `json:"screen_width"`
	ScreenHeight uint32 `json:"screen_height"`
}

type Win32DesktopMonitor struct {
	Name                string
	MonitorManufacturer string
	MonitorType         string
	ScreenWidth         uint32
	ScreenHeight        uint32
}
