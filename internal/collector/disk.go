package collector

import (
	"strings"

	"github.com/Nayati-Indonesia/spec-collector/models"
	"github.com/shirou/gopsutil/disk"
	"github.com/yusufpapurcu/wmi"
)

func GetDiskInfo() []models.Disk {
	partitions, _ := disk.Partitions(false)

	var drives []models.Win32DiskDrive
	err := wmi.Query("SELECT MediaType, Model FROM Win32_DiskDrive", &drives)
	if err != nil {
		return nil
	}

	var disks []models.Disk
	for _, p := range partitions {
		usage, _ := disk.Usage(p.Mountpoint)

		drive := p.Mountpoint
		totalGB := float64(usage.Total / 1024 / 1024 / 1024)
		diskType := getMediaType(drives)

		disks = append(disks, models.Disk{
			Drive:   drive,
			TotalGB: totalGB,
			Type:    diskType,
		})
	}

	return disks
}

func getMediaType(drives []models.Win32DiskDrive) string {
	diskType := "Unknown"
	if len(drives) > 0 {
		diskType = drives[0].MediaType
	}

	return cleanedType(diskType)
}

func cleanedType(diskType string) string {
	if strings.Contains(strings.ToLower(diskType), "solid state") {
		return "SSD"
	} else if strings.Contains(strings.ToLower(diskType), "external") {
		return "External Drive"
	} else if strings.Contains(strings.ToLower(diskType), "fixed") {
		return "HDD"
	}

	return "HDD"
}
