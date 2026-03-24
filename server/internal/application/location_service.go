package application

import (
	"errors"
	"math"

	"server/internal/domain"
)

const earthRadius = 6378137.0

type point struct {
	X float64
	Y float64
}

type anchorWithDistance struct {
	ID   string
	X    float64
	Y    float64
	Dist float64
}

type LocationService struct {
	Anchors []domain.Anchor
}

func NewLocationService() *LocationService {
	// Replace these coordinates with your real anchor coordinates
	return &LocationService{
		Anchors: []domain.Anchor{
			{ID: "ANCRE_01", Lat: 48.137100, Lon: 11.575100}, // top-left
			{ID: "ANCRE_02", Lat: 48.137100, Lon: 11.575300}, // top-right
			{ID: "ANCRE_03", Lat: 48.136950, Lon: 11.575300}, // bottom-right
			{ID: "ANCRE_04", Lat: 48.136950, Lon: 11.575100}, // bottom-left
		},
	}
}

func (s *LocationService) Estimate(req domain.DeviceLocationRequest) (domain.DeviceLocationResponse, error) {
	if len(req.Readings) < 3 {
		return domain.DeviceLocationResponse{}, errors.New("at least 3 readings are required")
	}

	if len(s.Anchors) < 3 {
		return domain.DeviceLocationResponse{}, errors.New("at least 3 anchors are required")
	}

	anchorMap := make(map[string]domain.Anchor)
	for _, a := range s.Anchors {
		anchorMap[a.ID] = a
	}

	origin := s.Anchors[0]

	var data []anchorWithDistance
	for _, r := range req.Readings {
		if r.Dist <= 0 || r.Dist > 1000 {
			continue
		}

		a, ok := anchorMap[r.ID]
		if !ok {
			continue
		}

		p := latLonToXY(a.Lat, a.Lon, origin.Lat, origin.Lon)

		data = append(data, anchorWithDistance{
			ID:   a.ID,
			X:    p.X,
			Y:    p.Y,
			Dist: r.Dist,
		})
	}

	if len(data) < 3 {
		return domain.DeviceLocationResponse{}, errors.New("not enough valid matching anchor readings")
	}

	estimated, err := estimatePosition(data)
	if err != nil {
		return domain.DeviceLocationResponse{}, err
	}

	geoLat, geoLon := xyToLatLon(estimated.X, estimated.Y, origin.Lat, origin.Lon)

	return domain.DeviceLocationResponse{
		DeviceID: req.DeviceID,
		Lat:      geoLat,
		Lon:      geoLon,
		X:        estimated.X,
		Y:        estimated.Y,
	}, nil
}

func latLonToXY(lat, lon, originLat, originLon float64) point {
	latRad := lat * math.Pi / 180.0
	lonRad := lon * math.Pi / 180.0
	originLatRad := originLat * math.Pi / 180.0
	originLonRad := originLon * math.Pi / 180.0

	x := (lonRad - originLonRad) * earthRadius * math.Cos(originLatRad)
	y := (latRad - originLatRad) * earthRadius

	return point{X: x, Y: y}
}

func xyToLatLon(x, y, originLat, originLon float64) (float64, float64) {
	originLatRad := originLat * math.Pi / 180.0
	originLonRad := originLon * math.Pi / 180.0

	latRad := y/earthRadius + originLatRad
	lonRad := x/(earthRadius*math.Cos(originLatRad)) + originLonRad

	return latRad * 180.0 / math.Pi, lonRad * 180.0 / math.Pi
}

func estimatePosition(anchors []anchorWithDistance) (point, error) {
	ref := anchors[0]

	var a11, a12, a22 float64
	var b1, b2 float64

	for i := 1; i < len(anchors); i++ {
		ai := anchors[i]

		Ai1 := 2 * (ai.X - ref.X)
		Ai2 := 2 * (ai.Y - ref.Y)
		bi := ref.Dist*ref.Dist - ai.Dist*ai.Dist +
			ai.X*ai.X - ref.X*ref.X +
			ai.Y*ai.Y - ref.Y*ref.Y

		a11 += Ai1 * Ai1
		a12 += Ai1 * Ai2
		a22 += Ai2 * Ai2

		b1 += Ai1 * bi
		b2 += Ai2 * bi
	}

	det := a11*a22 - a12*a12
	if math.Abs(det) < 1e-9 {
		return point{}, errors.New("anchors form a singular system")
	}

	x := (b1*a22 - b2*a12) / det
	y := (a11*b2 - a12*b1) / det

	return point{X: x, Y: y}, nil
}