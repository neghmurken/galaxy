package model

import (
	"math"
)

var (
	G       float32 = 6.6743e-11
	DENSITY float32 = 10
)

func Gravity(m1, m2, d float32) float32 {
	return G * (m2 * m1) / float32(math.Sqrt(float64(d)))
}

func Explosion(m1, m2, d, s float32) float32 {
	return Gravity(m1, m2, d) * s
}

func SizeToMass(size float32) float32 {
	return float32(math.Pi*math.Pow(float64(size), 2)) * DENSITY
}
