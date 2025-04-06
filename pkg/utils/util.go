package utils

import "math"

func Mean(slice []float32) float32 {
	var sum float32
	for _, v := range slice {
		sum += v
	}

	return sum / float32(len(slice))
}

func QuadBell(n float32) float32 {
	return float32(-math.Pow(float64(2*n-1), 2) + 1)
}

func GaussBell(n, a, o, w float32) float32 {
	return float32(math.Exp(-math.Pow(float64(n-o), 2)/(2*math.Pow(float64(w), 2)))*2 - 1)
}
