package value

type Vertex struct {
	Longitude float64
	Latitude  float64
}

type Polygon struct {
	Id       string
	Vertices []*Vertex
}
