package collector

import (
	"time"

	"github.com/Nayati-Indonesia/it-audit-collector/models"
)

func CollectAll() models.SpecResponse {
	var info models.SpecResponse

	info.Hostname = GetSystemInfo().Hostname
	info.MachineID = GetSystemInfo().MachineID
	info.Manufacturer = GetSystemInfo().Manufacturer
	info.IPAddress = GetSystemInfo().IPAddress
	info.Model = GetSystemInfo().Model

	info.OS = GetOSInfo()
	info.CPU = GetCPUInfo()
	info.Memory = GetMemoryInfo()
	info.GPU = getGPUInfo()
	info.Storage = GetDiskInfo()
	info.CollectedAt = time.Now()

	return info
}
