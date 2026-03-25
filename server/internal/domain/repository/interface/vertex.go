package interfaces

import (
	"context"
	"github.com/getLynx-tech/lynx/internal/domain/entity"
)

type VertexRepository interface {
	CreateVertex(ctx context.Context, polygonId string, longitude float64, latitude float64) (*entity.Vertex, error)
}
