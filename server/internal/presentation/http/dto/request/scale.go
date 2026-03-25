package request

type ScaleRequest struct {
	Pixels float64 `json:"pixels" validate:"required"`
	Meters float64 `json:"meters" validate:"required"`
}
