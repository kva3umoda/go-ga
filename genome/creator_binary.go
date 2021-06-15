package genome

import (
	"math/rand"
	"time"
)

type PopulationBinary struct {
	rnd        *rand.Rand
	genomeSize int
}

func NewPopulationBinary(genomeSize int) *PopulationBinary {
	return &PopulationBinary{
		rnd:        rand.New(rand.NewSource(time.Now().UnixNano())),
		genomeSize: genomeSize,
	}
}

func (p *PopulationBinary) Create(populationSize int) *Population {
	pop := newPopulation(populationSize)
	for i := 0; i < populationSize; i++ {
		ind := NewIndividual(p.genomeSize)
		for j := range ind.Genome {
			if rand.Float64() >= 0.5 {
				ind.Genome[j] = 1.0
			}
		}
		pop.Individuals = append(pop.Individuals, ind)
	}
	return pop
}
