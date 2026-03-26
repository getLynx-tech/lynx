package request

type Anchor struct {
	Id string  `json:"id" binding:"required"`
	X  float64 `json:"x" binding:"required"`
	Y  float64 `json:"y" binding:"required"`
}

type AnchorsRequest struct {
	Anchors []Anchor `json:"anchors" binding:"required"`
}
