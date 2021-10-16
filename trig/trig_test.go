package main

import (
	"math"
	"testing"
)

var TAU float64 = float64(math.Pi) * 2
var INVPI float64 = 1 / math.Pi
var INVTAU float64 = INVPI * 0.5

func BenchmarkStdSine(b *testing.B) {
	for i := 0; i < 10000000; i++ {
		math.Sin(float64(b.N) / 1000)
	}
}

func BenchmarkFastSine(b *testing.B) {
	for i := 0; i < 10000000; i++ {
		f_sin(float64(b.N) / 1000)
	}
}

func f_sin(radians float64) float64 {
	var x float64 = radians - math.Floor((radians+float64(math.Pi))*INVTAU)*TAU

	x = 4 * x * INVPI * (math.FMA(-math.Abs(x), INVPI, 1))
	return x * math.FMA(0.224, math.Abs(x), 0.776)
}
