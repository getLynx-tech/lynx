package interfaces

import (
	"context"
	"github.com/getLynx-tech/lynx/internal/domain/entity"
	"github.com/getLynx-tech/lynx/internal/domain/value"
)

type AnchorRepository interface {
	UpsertAnchors(ctx context.Context, anchors []*value.Anchor) ([]*entity.Anchor, error)
	GetAllAnchors(ctx context.Context) ([]*entity.Anchor, error)
}
