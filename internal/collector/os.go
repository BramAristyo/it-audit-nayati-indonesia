package collector

import (
	"runtime"

	"github.com/Nayati-Indonesia/it-audit-collector/models"
	"github.com/shirou/gopsutil/host"
)

func GetOSInfo() models.OS {
	hostStats, _ := host.Info()

	return models.OS{
		Name:         hostStats.Platform,
		Version:      hostStats.PlatformVersion,
		Architecture: runtime.GOARCH,
	}
}
