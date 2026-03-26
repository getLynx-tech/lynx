package interfaces

import (
	"context"
	"github.com/getLynx-tech/lynx/internal/domain/entity"
	"github.com/getLynx-tech/lynx/internal/domain/value"
)

type DeviceRepository interface {
	GetAllDevices(ctx context.Context) ([]*entity.Device, error)
	UpsertDevicePosition(ctx context.Context, devicePos *value.DevicePosition) error
}
