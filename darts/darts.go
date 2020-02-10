package darts

import "math"

func Score(x float64, y float64) int {
	dist := math.Sqrt(x*x + y*y)
	if dist > 10 {
		return 0
	}
	if dist > 5 {
		return 1
	}
	if dist > 1 {
		return 5
	}
	return 10
}
