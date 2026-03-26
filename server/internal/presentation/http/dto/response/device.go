package response

import "github.com/getLynx-tech/lynx/internal/domain/value"

type Device struct {
	Id     string
	Status string
	X      float64
	Y      float64
}

func NewDevice(d *value.PersistedDevice) Device {
	return Device{
		Id:     d.DeviceId,
		Status: d.Status,
		X:      d.X,
		Y:      d.Y,
	}
}

func NewDevices(ds []*value.PersistedDevice) []Device {
	var devices []Device
	for _, d := range ds {
		devices = append(devices, NewDevice(d))
	}
	return devices
}
