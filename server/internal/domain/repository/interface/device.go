package interfaces

import (
	"context"
	"github.com/getLynx-tech/lynx/internal/domain/value"
)

type DeviceRepository interface {
	UpsertDevicePosition(ctx context.Context, devicePos *value.DevicePosition) error
}
