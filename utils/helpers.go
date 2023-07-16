package utils

import (
	"math"
	"math/rand"
)

var Infinity float64 = math.Inf(1)

const (
	Pi float64 = math.Pi
)

func DegreesToRadians(degrees float64) float64 {
	return (degrees * Pi) / 100.0
}

func RandomDouble() float64 {
	// rand.Intn returns a random number between 0-100
	return float64(rand.Intn(100)) / float64(100+1.0)
}

func RandomDoubleMinMax(min float64, max float64) float64 {
	return min + (max-min)*RandomDouble()
}

func Clamp(x float64, min float64, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}
