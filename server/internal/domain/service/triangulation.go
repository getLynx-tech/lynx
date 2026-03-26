package service

import (
	"errors"
	"github.com/getLynx-tech/lynx/internal/domain/value"
)

type TriangulationService struct {
}

func NewTriangulationService() *TriangulationService {
	return &TriangulationService{}
}

func (ts *TriangulationService) TriangulatePosition(
	device *value.Device,
	anchors []*value.Anchor,
	scale *value.Scale,
) (*value.DevicePosition, error) {
	if len(anchors) < 3 {
		return nil, errors.New("at least 3 anchors required for triangulation")
	}

	// Build a lookup from anchor ID → distance (pixels) from device readings
	distMap := make(map[string]float64, len(device.Readings))
	for _, r := range device.Readings {
		if r.Dist > 0 {
			distMap[r.Id] = mmToPixels(float64(r.Dist), scale)
		}
	}

	// Collect anchors that have a matching reading
	type anchorDist struct {
		anchor *value.Anchor
		dist   float64
	}
	var matched []anchorDist
	for _, a := range anchors {
		if d, ok := distMap[a.Id]; ok {
			matched = append(matched, anchorDist{a, d})
		}
	}
	if len(matched) < 3 {
		return nil, errors.New("fewer than 3 anchors matched device readings")
	}

	// Weighted least-squares triangulation.
	//
	// Use the last anchor as the reference point and subtract it from every
	// other row to linearize the system (standard Chan–Ho reduction):
	//
	//   2(xᵢ−xₙ)X + 2(yᵢ−yₙ)Y = dₙ²−dᵢ²+xᵢ²−xₙ²+yᵢ²−yₙ²
	//
	// Weight each equation by 1/(dᵢ²) so closer, more reliable anchors
	// contribute more strongly to the solution.

	ref := matched[len(matched)-1]
	xn, yn, dn := ref.anchor.X, ref.anchor.Y, ref.dist

	// Accumulators for the 2×2 normal equations  (AᵀWA) p = AᵀWb
	var (
		ata00, ata01, ata11 float64 // symmetric 2×2
		atb0, atb1          float64 // right-hand side
	)

	for _, m := range matched[:len(matched)-1] {
		xi, yi, di := m.anchor.X, m.anchor.Y, m.dist

		// Row coefficients
		a0 := 2 * (xi - xn)
		a1 := 2 * (yi - yn)
		b := dn*dn - di*di + xi*xi - xn*xn + yi*yi - yn*yn

		// Weight: favor closer anchors; guard against zero distance
		w := 1.0
		if di > 0 {
			w = 1.0 / (di * di)
		}

		ata00 += w * a0 * a0
		ata01 += w * a0 * a1
		ata11 += w * a1 * a1
		atb0 += w * a0 * b
		atb1 += w * a1 * b
	}

	// Solve 2×2 system via Cramer's rule
	det := ata00*ata11 - ata01*ata01
	//if math.Abs(det) < 1e-10 {
	//	return nil, errors.New("anchors are collinear or too close together; cannot triangulate")
	//}

	x := (atb0*ata11 - atb1*ata01) / det
	y := (ata00*atb1 - ata01*atb0) / det

	return &value.DevicePosition{
		DeviceId: device.DeviceId,
		X:        x,
		Y:        y,
	}, nil
}

// mmToPixels converts a distance in millimeters to pixels using the provided scale.
func mmToPixels(mm float64, scale *value.Scale) float64 {
	if scale == nil || scale.Meters == 0 {
		return mm
	}
	return (mm / 1000) * (scale.Pixels / scale.Meters)
}
