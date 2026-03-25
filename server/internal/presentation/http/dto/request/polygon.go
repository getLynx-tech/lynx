package request

type Vertex struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type CreatePolygon struct {
	Id       string   `json:"id" binding:"required"`
	Vertices []Vertex `json:"vertices" binding:"required"`
}
