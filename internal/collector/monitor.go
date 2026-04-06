package collector

import (
	"fmt"

	"github.com/Nayati-Indonesia/it-audit-collector/models"
	"github.com/yusufpapurcu/wmi"
)

func GetMonitorInfo() ([]models.Monitor, error) {
	var wmiMonitors []models.Win32DesktopMonitor

	query := "SELECT Name, MonitorManufacturer, MonitorType, ScreenWidth, ScreenHeight FROM Win32_DesktopMonitor"
	err := wmi.Query(query, &wmiMonitors)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var monitors []models.Monitor
	for _, m := range wmiMonitors {
		resolution := fmt.Sprintf("%dx%d", m.ScreenWidth, m.ScreenHeight)

		monitors = append(monitors, models.Monitor{
			Name:         m.Name,
			Manufacturer: m.MonitorManufacturer,
			Resolution:   resolution,
			ScreenWidth:  m.ScreenWidth,
			ScreenHeight: m.ScreenHeight,
		})
	}

	return monitors, nil
}
