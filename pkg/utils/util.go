package utils

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	G float32 = 6.6743e-11
)

func Gravity(m1, m2, d float32) float32 {
	return G * (m2 * m1) / float32(math.Sqrt(float64(d)))
}

func Explosion(m1, m2, d, s float32) float32 {
	return Gravity(m1, m2, d) * s
}

func SizeToMass(size float32) float32 {
	if size == 0 {
		return 0
	}

	return rl.Remap(size, 1, 10, 100, 10000)
}
