package population

import (
	"github.com/kva3umoda/go-ga/genome"
	"github.com/kva3umoda/go-ga/rand"
)

type floatPopulation struct {
	genomeSize int
	up         []float64
	low        []float64
}

func FloatPopulation(genomeSize int, low []float64, up []float64) Creator {
	return &floatPopulation{
		genomeSize: genomeSize,
		up:         up,
		low:        low,
	}
}

func (f *floatPopulation) Create(populationSize int) *Population {
	pop := newPopulation(populationSize)
	for i := 0; i < populationSize; i++ {
		ind := genome.NewIndividual(f.genomeSize)
		for j := range ind.Genome {
			rnd := rand.Float()
			delta := f.up[j] - f.low[j]
			v := rnd * delta
			ind.Genome[j] = v + f.low[j]
		}
		pop.Individuals = append(pop.Individuals, ind)
	}
	return pop
}
