package helper

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
