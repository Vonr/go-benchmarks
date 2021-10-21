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
	for i := 0.0; i < float64(b.N); i++ {
		math.Sin(i)
	}
}

func BenchmarkFastSine(b *testing.B) {
	for i := 0.0; i < float64(b.N); i++ {
		FastSin(i)
	}
}

func BenchmarkStdCos(b *testing.B) {
	for i := 0.0; i < float64(b.N); i++ {
		math.Cos(i)
	}
}

func BenchmarkFastCos(b *testing.B) {
	for i := 0.0; i < float64(b.N); i++ {
		FastCos(i)
	}
}

func BenchmarkStdAtan2(b *testing.B) {
	for i := 0.0; i < float64(b.N); i++ {
		math.Atan2(i, i)
	}
}

func BenchmarkFastAtan2(b *testing.B) {
	for i := 0.0; i < float64(b.N); i++ {
		FastAtan2(i, i)
	}
}

func BenchmarkStdAtan(b *testing.B) {
	for i := 0.0; i < float64(b.N); i++ {
		math.Atan(i)
	}
}

func BenchmarkFastAtan(b *testing.B) {
	for i := 0.0; i < float64(b.N); i++ {
		FastAtan(i)
	}
}

func TestFastSinCos(t *testing.T) {
	maxError := 0.0
	for i := 0.0; i < Tau; i += 0.0000001 {
		maxError = math.Max(maxError, math.Abs(FastSin(i)-math.Sin(i)))
	}
	fmt.Println("FastSin/Cos Max Error:", maxError)
}

func TestFastAtan2(t *testing.T) {
	maxError := 0.0
	for x := 0.00001; x < Tau; x += 0.00001 {
		for y := 0.01; y < 1; y += 0.01 {
			maxError = math.Max(maxError, math.Abs(FastAtan2(y, x)-math.Atan2(y, x)))
		}
	}
	fmt.Println("FastAtan2 Max Error:", maxError)
}

func TestFastAtan(t *testing.T) {
	maxError := 0.0
	for i := 0.0; i < Tau; i += 0.0000001 {
		maxError = math.Max(maxError, math.Abs(FastAtan(i)-math.Atan(i)))
	}
	fmt.Println("FastAtan", maxError)
}

func FastSin(radians float64) float64 {
	var x float64 = radians - math.Floor((radians+Pi)*InvTau)*Tau

	x = 4 * x * InvPi * (math.FMA(-math.Abs(x), InvPi, 1))
	return x * math.FMA(0.224, math.Abs(x), 0.776)
}

func FastCos(radians float64) float64 {
	return FastSin(radians + Pi*0.5)
}

func FastAtan2(y float64, x float64) float64 {
	t3 := math.Abs(x)
	t1 := math.Abs(y)
	t0 := math.Max(t3, t1)
	t1 = math.Min(t3, t1)
	t3 = 1 / t0
	t3 = t1 * t3

	t4 := t3 * t3
	t0 = -0.013480470
	t0 = math.FMA(t0, t4, 0.057477314)
	t0 = math.FMA(t0, t4, -0.121239071)
	t0 = math.FMA(t0, t4, 0.195635925)
	t0 = math.FMA(t0, t4, -0.332994597)
	t0 = math.FMA(t0, t4, 0.999995630)
	t3 = t0 * t3

	if math.Abs(y) > math.Abs(x) {
		t3 = math.FMA(Pi, 0.5, -t3)
	}
	if x < 0 {
		t3 = Pi - t3
	}

	if y < 0 {
		return -t3
	}
	return t3
}

func FastAtan(y float64) float64 {
	t1 := math.Abs(y)
	t0 := math.Max(1, t1)
	t1 = math.Min(1, t1)
	t1 = t1 / t0

	t4 := t1 * t1
	t0 = -0.013480470
	t0 = math.FMA(t0, t4, 0.057477314)
	t0 = math.FMA(t0, t4, -0.121239071)
	t0 = math.FMA(t0, t4, 0.195635925)
	t0 = math.FMA(t0, t4, -0.332994597)
	t0 = math.FMA(t0, t4, 0.999995630)
	t1 = t0 * t1

	if math.Abs(y) > 1.0 {
		t1 = math.FMA(Pi, 0.5, -t1)
	}

	if y < 0 {
		return -t1
	}
	return t1
}
