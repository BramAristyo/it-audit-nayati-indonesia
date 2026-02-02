package collector

import (
	"github.com/Nayati-Indonesia/spec-collector/models"
	"github.com/shirou/gopsutil/cpu"
)

func GetCPUInfo() models.CPU {
	cpuStats, _ := cpu.Info()
	cores, _ := cpu.Counts(false)
	threads, _ := cpu.Counts(true)

	modelName := "Unknown"
	frequency := 0.0

	if len(cpuStats) > 0 {
		modelName = cpuStats[0].ModelName
		frequency = cpuStats[0].Mhz
	}

	return models.CPU{
		Model:     modelName,
		Cores:     cores,
		Threads:   threads,
		Frequency: frequency,
	}
}
