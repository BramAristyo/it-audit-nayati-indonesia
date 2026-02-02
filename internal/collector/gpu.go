package collector

import (
	"github.com/Nayati-Indonesia/spec-collector/models"
	"github.com/yusufpapurcu/wmi"
)

func getGPUInfo() []models.GPU {
	var gpus []models.GPU

	err := wmi.Query("SELECT Name, AdapterRam, DriverVersion, VideoProcessor FROM Win32_VideoController", &gpus)
	if err != nil {
		return []models.GPU{}
	}

	return gpus
}
