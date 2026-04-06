package models

type Printer struct {
	Name       string `json:"name"`
	DriverName string `json:"driver_name"`
	PortName   string `json:"port_name"`
	Status     string `json:"status"`
	Shared     bool   `json:"shared"`
	Network    bool   `json:"network"`
	DeviceID   string `json:"device_id"`
}

type Win32Printer struct {
	Name          string
	DriverName    string
	PortName      string
	PrinterStatus uint16
	Shared        bool
	Network       bool
	DeviceID      string
}
