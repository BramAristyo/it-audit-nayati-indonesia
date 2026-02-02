package models

type Identity struct {
	Hostname     string `json:"host_name"`
	MachineID    string `json:"machine_id"`
	IPAddress    string `json:"ip_address"`
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
}
