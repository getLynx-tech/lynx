package repository

import (
	"context"
	interfaces "github.com/getLynx-tech/lynx/internal/domain/repository/interface"
	"github.com/getLynx-tech/lynx/internal/domain/value"
	"github.com/getLynx-tech/lynx/internal/infrastructure/postgres/sqlc"
)

type DeviceRepository struct {
	queries *sqlc.Queries
}

func NewDeviceRepository(queries *sqlc.Queries) interfaces.DeviceRepository {
	return &DeviceRepository{queries: queries}
}
func (dr *DeviceRepository) UpsertDevicePosition(ctx context.Context, devicePos *value.DevicePosition) error {
	status := "inactive"
	if devicePos.IsActive {
		status = "active"
	}

	_, err := dr.queries.UpsertDevice(ctx, sqlc.UpsertDeviceParams{
		DeviceID: devicePos.DeviceId,
		Status:   status,
		X:        devicePos.X,
		Y:        devicePos.Y,
	})
	return err
}
