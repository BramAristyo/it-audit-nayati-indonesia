package collector

import (
	"runtime"

	"github.com/Nayati-Indonesia/spec-collector/models"
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
