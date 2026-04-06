package collector

import (
	"fmt"

	"github.com/Nayati-Indonesia/it-audit-collector/models"
	"github.com/yusufpapurcu/wmi"
)

func GetPrinterInfo() ([]models.Printer, error) {
	var printers []models.Printer
	query := "SELECT Name, DriverName, PortName, Shared, Network, DeviceID, Status FROM Win32_Printer"

	if err := wmi.Query(query, &printers); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return printers, nil
}
