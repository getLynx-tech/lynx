package value

type Reading struct {
	Id   string
	Rssi int
	Dist int
}

type Device struct {
	DeviceId string
	IsActive bool
	Readings []*Reading
}

type DevicePosition struct {
	DeviceId string
	IsActive bool
	X        float64
	Y        float64
}
