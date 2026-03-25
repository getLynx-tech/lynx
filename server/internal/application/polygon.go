package application

import (
	"context"
	interfaces "github.com/getLynx-tech/lynx/internal/domain/repository/interface"
	"github.com/getLynx-tech/lynx/internal/domain/value"
)

type PolygonApplication struct {
	vertexRepository interfaces.VertexRepository
}

func NewVertexApplication(
	vertexRepository interfaces.VertexRepository,
) *PolygonApplication {
	return &PolygonApplication{
		vertexRepository: vertexRepository,
	}
}

func (va *PolygonApplication) CreatePolygon(ctx context.Context, polygon *value.Polygon) error {
	for _, vertex := range polygon.Vertices {
		_, err := va.vertexRepository.CreateVertex(ctx, polygon.Id, vertex.Longitude, vertex.Latitude)
		if err != nil {
			return err
		}
	}
	return nil
}
