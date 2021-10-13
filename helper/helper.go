package helper

import "math"

func Mini(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Minf(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

func Minia(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}
	v := vals[0]

	for i := 1; i < len(vals); i++ {
		if v > vals[i] {
			v = vals[0]
		}
	}

	return v
}

func Maxi(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Maxf(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}

func Distance(p1, p2 []float64) float64 {
	size := Mini(len(p1), len(p2))
	sum := 0.0
	for i := 0; i < size; i++ {
		sum += math.Pow(p1[i]-p2[i], 2)
	}
	return math.Sqrt(sum)
}
