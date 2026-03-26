package application

import (
	"context"
	interfaces "github.com/getLynx-tech/lynx/internal/domain/repository/interface"
	"github.com/getLynx-tech/lynx/internal/domain/service"
	"github.com/getLynx-tech/lynx/internal/domain/value"
)

type DeviceApplication struct {
	deviceRepository     interfaces.DeviceRepository
	anchorRepository     interfaces.AnchorRepository
	scaleRepository      interfaces.ScaleRepository
	triangulationService *service.TriangulationService
}

func NewDeviceApplication(
	deviceRepository interfaces.DeviceRepository,
	anchorRepository interfaces.AnchorRepository,
	scaleRepository interfaces.ScaleRepository,
	triangulationService *service.TriangulationService,

) *DeviceApplication {
	return &DeviceApplication{
		deviceRepository:     deviceRepository,
		anchorRepository:     anchorRepository,
		scaleRepository:      scaleRepository,
		triangulationService: triangulationService,
	}
}

func (da *DeviceApplication) GetAllDevices(ctx context.Context) ([]*value.PersistedDevice, error) {
	devices, err := da.deviceRepository.GetAllDevices(ctx)
	if err != nil {
		return nil, err
	}

	deviceValues := make([]*value.PersistedDevice, len(devices))
	for i, device := range devices {
		deviceValues[i] = &value.PersistedDevice{
			DeviceId: device.Id,
			Status:   device.Status,
			X:        device.X,
			Y:        device.Y,
		}
	}

	return deviceValues, nil
}

func (da *DeviceApplication) UpsertDevicePosition(ctx context.Context, device *value.Device) error {
	scale, err := da.scaleRepository.GetScale(ctx)
	if err != nil {
		return err
	}

	anchors, err := da.anchorRepository.GetAllAnchors(ctx)
	if err != nil {
		return err
	}

	anchorValues := make([]*value.Anchor, len(anchors))
	for i, anchor := range anchors {
		anchorValues[i] = &value.Anchor{
			Id: anchor.Id,
			X:  anchor.X,
			Y:  anchor.Y,
		}
	}

	devicePos, err := da.triangulationService.TriangulatePosition(device, anchorValues, &value.Scale{
		Meters: scale.Meters,
		Pixels: scale.Pixels,
	})
	if err != nil {
		return err
	}

	err = da.deviceRepository.UpsertDevicePosition(ctx, devicePos)
	if err != nil {
		return err
	}
	return nil
}
