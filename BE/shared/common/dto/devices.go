package dto

type DevicesOutput struct {
	ID          int    `json:"id"`
	DeviceID    string `json:"device_id"`
	Location    string `json:"location"`
	InstalledAt string `json:"installed_at"`
	Active      bool   `json:"ative"`
	Total       int    `json:"total"`
}

type DevicesInput struct {
	DeviceID string `json:"device_id"`
	Location string `json:"location"`
}

type DeviceOutput struct {
	ID        int    `json:"id"`
	DeviceID  string `json:"device_id"`
	Location  string `json:"location"`
	Installed string `json:"installed"`
	Active    bool   `json:"active"`
}
