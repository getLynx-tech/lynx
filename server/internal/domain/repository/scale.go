package repository

import (
	"context"
	"github.com/getLynx-tech/lynx/internal/domain/entity"
	interfaces "github.com/getLynx-tech/lynx/internal/domain/repository/interface"
	"github.com/getLynx-tech/lynx/internal/domain/value"
	"github.com/getLynx-tech/lynx/internal/infrastructure/postgres/sqlc"
)

type ScaleRepository struct {
	queries *sqlc.Queries
}

func NewScaleRepository(queries *sqlc.Queries) interfaces.ScaleRepository {
	return &ScaleRepository{queries: queries}
}

func (sr *ScaleRepository) CreateScale(ctx context.Context, scale *value.Scale) error {
	_, err := sr.queries.CreateScale(ctx, sqlc.CreateScaleParams{
		Meters: scale.Meters,
		Pixels: scale.Pixels,
	})
	return err
}

func (sr *ScaleRepository) GetScale(ctx context.Context) (*entity.Scale, error) {
	scale, err := sr.queries.GetScale(ctx)
	if err != nil {
		return nil, err
	}
	return &entity.Scale{
		Meters: scale.Meters,
		Pixels: scale.Pixels,
	}, nil
}
