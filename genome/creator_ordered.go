package genome

import (
	"math/rand"
	"time"
)

type PopulationOrdered struct {
	rnd            *rand.Rand
	genomeSize     int
}

func NewPopulationOrdered(genomeSize int) *PopulationOrdered {
	return &PopulationOrdered{
		rnd:            rand.New(rand.NewSource(time.Now().UnixNano())),
		genomeSize:     genomeSize,
	}
}

func (p *PopulationOrdered) Create(populationSize int) *Population {
	pop := newPopulation(populationSize)
	for i := 0; i < populationSize; i++ {
		ind := NewIndividual(p.genomeSize)
		begin := 0
		for i := range ind.Genome {
			ind.Genome[i] = float64(begin)
			begin++
		}

		rand.Shuffle(p.genomeSize, func(i, j int) {
			ind.Genome[i], ind.Genome[j] = ind.Genome[j], ind.Genome[i]
		})
		pop.Individuals = append(pop.Individuals, ind)
	}
	return pop
}
