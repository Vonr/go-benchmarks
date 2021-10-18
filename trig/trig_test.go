package main

import (
	"fmt"
	"math"
	"testing"
)

const (
	Pi     = float64(math.Pi)
	Tau    = Pi * 2
	InvPi  = 1 / Pi
	InvTau = InvPi * 0.5
)

func BenchmarkStdSine(b *testing.B) {
	for i := 0; i < 10000000; i++ {
		math.Sin(float64(b.N) / 1000)
	}
}

func BenchmarkFastSine(b *testing.B) {
	for i := 0; i < 10000000; i++ {
		FastSin(float64(b.N) / 1000)
	}
}

func BenchmarkStdCos(b *testing.B) {
	for i := 0; i < 10000000; i++ {
		math.Cos(float64(b.N) / 1000)
	}
}

func BenchmarkFastCos(b *testing.B) {
	for i := 0; i < 10000000; i++ {
		FastCos(float64(b.N) / 1000)
	}
}

func TestFastSin(t *testing.T) {
	maxError := 0.0
	for i := 0.0; i < Tau; i += 0.0000001 {
		maxError = math.Max(maxError, math.Abs(FastSin(i)-math.Sin(i)))
	}
	fmt.Println("FastSin Max Error:", maxError)
}

func TestFastCos(t *testing.T) {
	maxError := 0.0
	for i := 0.0; i < Tau; i += 0.0000001 {
		maxError = math.Max(maxError, math.Abs(FastCos(i)-math.Cos(i)))
	}
	fmt.Println("FastCos Max Error:", maxError)
}
func FastSin(radians float64) float64 {
	var x float64 = radians - math.Floor((radians+Pi)*InvTau)*Tau

	x = 4 * x * InvPi * (math.FMA(-math.Abs(x), InvPi, 1))
	return x * math.FMA(0.224, math.Abs(x), 0.776)
}

func FastCos(radians float64) float64 {
	return FastSin(radians + Pi*0.5)
}
