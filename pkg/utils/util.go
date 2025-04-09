package utils

import "math"

var (
	G float32 = 6.6743e-11
)

func Gravity(m1, m2, d float32) float32 {
	return G * (m2 * m1) / float32(math.Sqrt(float64(d)))
}
