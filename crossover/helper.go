package crossover

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func copy2(ind1, ind2 []float64) (nind1, nind2 []float64) {
	nind1 = make([]float64, len(ind1))
	copy(nind1, ind1)
	nind2 = make([]float64, len(ind2))
	copy(nind2, ind2)

	return
}
