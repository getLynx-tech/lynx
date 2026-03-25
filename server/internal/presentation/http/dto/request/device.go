package request

type Reading struct {
	Id   string `json:"id" binding:"required"`
	Rssi int    `json:"rssi" binding:"required"`
	Dist int    `json:"dist" binding:"required"`
}

type DeviceRequest struct {
	DeviceId string    `json:"device_id" binding:"required"`
	Readings []Reading `json:"readings" binding:"required"`
}
