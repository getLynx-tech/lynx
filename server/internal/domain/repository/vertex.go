package repository

import (
	"context"
	"github.com/getLynx-tech/lynx/internal/domain/entity"
	interfaces "github.com/getLynx-tech/lynx/internal/domain/repository/interface"
	"github.com/getLynx-tech/lynx/internal/infrastructure/postgres/sqlc"
)

type VertexRepository struct {
	queries *sqlc.Queries
}

var _ interfaces.VertexRepository = (*VertexRepository)(nil)

func NewVertexRepository(queries *sqlc.Queries) interfaces.VertexRepository {
	return &VertexRepository{
		queries: queries,
	}
}

func (vr *VertexRepository) CreateVertex(ctx context.Context, polygonId string, longitude float64, latitude float64) (*entity.Vertex, error) {
	vertex, err := vr.queries.CreateVertex(ctx, sqlc.CreateVertexParams{
		PolygonID: polygonId,
		Longitude: longitude,
		Latitude:  latitude,
	})
	if err != nil {
		return nil, err
	}

	return &entity.Vertex{
		PolygonId: vertex.PolygonID,
		Longitude: vertex.Longitude,
		Latitude:  vertex.Latitude,
	}, nil
}
