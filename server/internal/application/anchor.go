package application

import (
	"context"
	interfaces "github.com/getLynx-tech/lynx/internal/domain/repository/interface"
	"github.com/getLynx-tech/lynx/internal/domain/value"
)

type AnchorApplication struct {
	anchorRepository interfaces.AnchorRepository
}

func NewAnchorApplication(anchorRepository interfaces.AnchorRepository) *AnchorApplication {
	return &AnchorApplication{
		anchorRepository: anchorRepository,
	}
}

func (aa *AnchorApplication) UpsertAnchors(ctx context.Context, anchors []*value.Anchor) error {
	_, err := aa.anchorRepository.UpsertAnchors(ctx, anchors)
	return err
}

func (aa *AnchorApplication) GetAllAnchors(ctx context.Context) ([]*value.Anchor, error) {
	anchors, err := aa.anchorRepository.GetAllAnchors(ctx)
	if err != nil {
		return nil, err
	}

	anchorValues := make([]*value.Anchor, len(anchors))
	for i, anchor := range anchors {
		anchorValues[i] = &value.Anchor{
			Id: anchor.Id,
			X:  anchor.X,
			Y:  anchor.Y,
		}
	}
	return anchorValues, nil
}
