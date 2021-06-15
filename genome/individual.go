package genome



// Индивидум с геномом int
type Individual struct {
	Fitness float64   // приспособленность
	Genome  []float64 // геном
}

func NewIndividual(size int) Individual {
	return Individual{
		Fitness: 0,
		Genome:  make([]float64, size),
	}
}

func (ind *Individual) Clone() Individual {
	dst := Individual{
		Fitness: ind.Fitness,
		Genome:  make([]float64, len(ind.Genome)),
	}
	copy(dst.Genome, ind.Genome)

	return dst
}

func (ind *Individual) GetInt(i int) int {
	return int(ind.Genome[i])
}

func (ind *Individual) GetFloat(i int) float64 {
	return ind.Genome[i]
}

func (ind *Individual) Equal(other Individual) bool {
	if len(ind.Genome) != len(other.Genome) {
		return false
	}

	for i := range ind.Genome {
		if ind.Genome[i] != other.Genome[i] {
			return false
		}
	}

	return true
}
