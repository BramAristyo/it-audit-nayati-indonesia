package collector

import (
	"net"

	"github.com/Nayati-Indonesia/spec-collector/models"
	"github.com/shirou/gopsutil/host"
	"github.com/yusufpapurcu/wmi"
)

// TODO: Get Primary IP, Just get the First data [0]
//func getIPAddress() []string {
//	var ips []string
//
//	addresses, err := net.InterfaceAddrs()
//	if err != nil {
//		return ips
//	}
//
//	for _, addr := range addresses {
//		if ip, ok := addr.(*net.IPNet); ok && !ip.IP.IsLoopback() {
//			if ip.IP.To4() != nil {
//				ips = append(ips, ip.IP.String())
//			}
//		}
//	}
//
//	return ips
//}

func getIPAddress() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return ""
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

func getHardwareInfo() models.Win32ComputerSystem {
	var result []models.Win32ComputerSystem

	err := wmi.Query("SELECT Manufacturer, Model FROM Win32_ComputerSystem", &result)
	if err != nil || len(result) == 0 {
		return models.Win32ComputerSystem{}
	}

	return result[0]
}

func GetSystemInfo() *models.System {
	hostStats, _ := host.Info()
	hardwareInfo := getHardwareInfo()
	IP := getIPAddress()

	return &models.System{
		Hostname:     hostStats.Hostname,
		MachineID:    hostStats.HostID,
		Manufacturer: hardwareInfo.Manufacturer,
		Model:        hardwareInfo.Model,
		IPAddress:    IP,
	}

}
