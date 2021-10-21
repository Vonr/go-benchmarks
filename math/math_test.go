package fastmath

import (
	"fmt"
	"math"
	"testing"
)

func TestFastFloor(t *testing.T) {
	fmt.Println("Floor Tests 1:", math.Floor(-1.433) == FastFloor(-1.433), "2:", math.Floor(-1.823) == FastFloor(-1.823), "3:", math.Floor(21.2) == FastFloor(21.2))
}

func BenchmarkStdMod(b *testing.B) {
	for x := 0.0; x < 1000.0; x++ {
		for y := 0.0; y < float64(b.N); y++ {
			math.Mod(x, y)
		}
	}
}

func BenchmarkFastMod(b *testing.B) {
	for x := 0.0; x < 1000.0; x++ {
		for y := 0.0; y < float64(b.N); y++ {
			FastMod(x, y)
		}
	}
}

func BenchmarkStdFloor(b *testing.B) {
	for x := 0.0; x < float64(b.N); x++ {
		math.Floor(x)
	}
}

func BenchmarkFastFloor(b *testing.B) {
	for x := 0.0; x < float64(b.N); x++ {
		FastFloor(x)
	}
}

func FastMod(x float64, y float64) float64 {
	return y - math.Floor(x/y)*y
}

func FastFloor(x float64) float64 {
	trunc := math.Trunc(x)
	if x < 0 && x != trunc {
		return trunc - 1
	}
	return trunc
}
