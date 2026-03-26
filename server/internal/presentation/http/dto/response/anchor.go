package response

import "github.com/getLynx-tech/lynx/internal/domain/value"

type Anchor struct {
	Id string  `json:"id" validate:"required"`
	X  float64 `json:"x" validate:"required"`
	Y  float64 `json:"y" validate:"required"`
}

func NewAnchor(a *value.Anchor) Anchor {
	return Anchor{
		Id: a.Id,
		X:  a.X,
		Y:  a.Y,
	}
}

func NewAnchors(as []*value.Anchor) []Anchor {
	var anchors []Anchor
	for _, a := range as {
		anchors = append(anchors, NewAnchor(a))
	}
	return anchors
}
