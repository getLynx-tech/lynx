package domain

type AnchorReading struct {
	ID   string  `json:"id"`
	RSSI int     `json:"rssi"`
	Dist float64 `json:"dist"`
}

type DeviceLocationRequest struct {
	DeviceID string          `json:"device_id"`
	Readings []AnchorReading `json:"readings"`
}

type Anchor struct {
	ID  string  `json:"id"`
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type DeviceLocationResponse struct {
	DeviceID string  `json:"device_id"`
	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
}