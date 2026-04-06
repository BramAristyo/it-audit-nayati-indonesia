package collector

import (
	"github.com/Nayati-Indonesia/it-audit-collector/models"
	"github.com/shirou/gopsutil/mem"
)

func GetMemoryInfo() models.Memory {
	memStats, _ := mem.VirtualMemory()
	ramGb := memStats.Total / 1024 / 1024 / 1024

	return models.Memory{
		TotalGB: float64(ramGb),
	}
}
