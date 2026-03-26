package repository

import (
	"context"
	"github.com/getLynx-tech/lynx/internal/domain/entity"
	interfaces "github.com/getLynx-tech/lynx/internal/domain/repository/interface"
	"github.com/getLynx-tech/lynx/internal/domain/value"
	"github.com/getLynx-tech/lynx/internal/infrastructure/postgres/sqlc"
)

type AnchorRepository struct {
	queries *sqlc.Queries
}

func NewAnchorRepository(queries *sqlc.Queries) interfaces.AnchorRepository {
	return &AnchorRepository{
		queries: queries,
	}
}

func (ar *AnchorRepository) UpsertAnchors(ctx context.Context, anchors []*value.Anchor) ([]*entity.Anchor, error) {
	err := ar.queries.DeleteAllAnchors(ctx)
	if err != nil {
		return nil, err
	}

	anchorEntities := make([]*entity.Anchor, len(anchors))

	for i, anchor := range anchors {
		res, err := ar.queries.CreateAnchor(ctx, sqlc.CreateAnchorParams{
			AnchorID: anchor.Id,
			X:        anchor.X,
			Y:        anchor.Y,
		})
		if err != nil {
			return nil, err
		}
		anchorEntities[i] = &entity.Anchor{
			Id: res.AnchorID,
			X:  res.X,
			Y:  res.Y,
		}
	}
	return anchorEntities, nil
}

func (ar *AnchorRepository) GetAllAnchors(ctx context.Context) ([]*entity.Anchor, error) {
	res, err := ar.queries.GetAllAnchors(ctx)
	if err != nil {
		return nil, err
	}

	anchors := make([]*entity.Anchor, len(res))
	for i, r := range res {
		anchors[i] = &entity.Anchor{
			Id: r.AnchorID,
			X:  r.X,
			Y:  r.Y,
		}
	}
	return anchors, nil
}
