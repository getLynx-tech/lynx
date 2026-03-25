package interfaces

import (
	"context"
	"github.com/getLynx-tech/lynx/internal/domain/entity"
	"github.com/getLynx-tech/lynx/internal/domain/value"
)

type ScaleRepository interface {
	CreateScale(ctx context.Context, scale *value.Scale) error
	GetScale(ctx context.Context) (*entity.Scale, error)
}
