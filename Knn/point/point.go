package point

import "math"

type Point struct {
	X int
	Y int
}

// Calculate distances of one point to multiple points
func (p *Point) Distance(p2 []Point) []float64 {

	first := 0.0
	second := 0.0
	var distances []float64

	for _, val := range p2 {
		first = math.Pow(float64(p.X-val.X), 2)
		second = math.Pow(float64(p.Y-val.Y), 2)
		distances = append(distances, math.Sqrt(first+second))
	}

	return distances
}
