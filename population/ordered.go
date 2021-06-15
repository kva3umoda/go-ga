package population

import (
	"github.com/kva3umoda/go-ga/genome"
	"github.com/kva3umoda/go-ga/rand"
)

type OrderedPopulation struct {
	genomeSize int
}

func NewOrderedPopulation(genomeSize int) *OrderedPopulation {
	return &OrderedPopulation{
		genomeSize: genomeSize,
	}
}

func (p *OrderedPopulation) Create(populationSize int) *Population {
	pop := newPopulation(populationSize)
	for i := 0; i < populationSize; i++ {
		ind := genome.NewIndividual(p.genomeSize)
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
