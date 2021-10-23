package fastmath

import (
	"fmt"
	"math"
	"testing"
)

func TestFastFloor(t *testing.T) {
	fmt.Println("Floor Tests 1:", math.Floor(-1.433) == FastFloor(-1.433), "2:", math.Floor(-1.823) == FastFloor(-1.823), "3:", math.Floor(21.2) == FastFloor(21.2))
}

func TestFastMod(t *testing.T) {
	fmt.Println("Mod Tests 1:", math.Mod(-1.433, -1) == FastMod(-1.433, -1), "2:", math.Mod(-7.823, 2) == FastMod(-7.823, 2), "3:", math.Mod(21.2, -3) == FastMod(21.2, -3))
}

func TestFastMax(t *testing.T) {
	fmt.Println("Max Tests 1:", math.Max(-1.433, -2.238) == FastMax(-1.433, -2.238), "2:", math.Max(-1.823, 7.484) == FastMax(-1.823, 7.484), "3:", math.Max(21.2, 3.744) == FastMax(21.2, 3.744))
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

func BenchmarkStdMax(b *testing.B) {
	for x := 0.0; x < float64(b.N); x++ {
		math.Max(x, 0.0)
		math.Max(x, -0.0)
		math.Max(x, 1)
		math.Max(x, -1)
	}
}

func BenchmarkFastMax(b *testing.B) {
	for x := 0.0; x < float64(b.N); x++ {
		FastMax(x, 0)
		FastMax(x, -0)
		FastMax(x, 1)
		FastMax(x, -1)
	}
}

func BenchmarkStdAbs(b *testing.B) {
	for x := 0.0; x < float64(b.N); x++ {
		math.Abs(x)
	}
}

func BenchmarkFastAbs(b *testing.B) {
	for x := 0.0; x < float64(b.N); x++ {
		FastAbs(x)
	}
}

func FastMod(x, y float64) float64 {
	s := 0.0
	if (x < 0 && y > 0) || (y < 0 && x > 0) {
		s = y
	}
	return x - math.Floor(x/y)*y - s
}

func FastFloor(x float64) float64 {
	trunc := math.Trunc(x)
	if x < 0 && x != trunc {
		return trunc - 1
	}
	return trunc
}

func FastMax(x, y float64) float64 {
	if x == y && x == -0.0 {
		return -0.0
	}
	if y > x {
		return y
	}
	return x
}

func FastAbs(x float64) float64 {
	return math.Copysign(x, +0)
}
