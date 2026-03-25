package response

import "github.com/getLynx-tech/lynx/internal/domain/value"

type Scale struct {
	Pixels float64 `json:"pixels" validate:"required"`
	Meters float64 `json:"meters" validate:"required"`
}

func NewScale(s *value.Scale) Scale {
	return Scale{
		Pixels: s.Pixels,
		Meters: s.Meters,
	}
}
