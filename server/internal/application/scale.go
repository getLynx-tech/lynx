package application

import (
	"context"
	interfaces "github.com/getLynx-tech/lynx/internal/domain/repository/interface"
	"github.com/getLynx-tech/lynx/internal/domain/value"
)

type ScaleApplication struct {
	scaleRepository interfaces.ScaleRepository
}

func NewScaleApplication(scaleRepository interfaces.ScaleRepository) *ScaleApplication {
	return &ScaleApplication{
		scaleRepository: scaleRepository,
	}
}

func (sa *ScaleApplication) CreateScale(ctx context.Context, scale *value.Scale) error {
	return sa.scaleRepository.CreateScale(ctx, scale)
}

func (sa *ScaleApplication) GetScale(ctx context.Context) (*value.Scale, error) {
	scale, err := sa.scaleRepository.GetScale(ctx)
	if err != nil {
		return nil, err
	}

	return &value.Scale{
		Meters: scale.Meters,
		Pixels: scale.Pixels,
	}, nil
}
