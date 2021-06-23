package population

import (
	"github.com/kva3umoda/go-ga/genome"
	"github.com/kva3umoda/go-ga/rand"
)

type orderedPopulation struct {
	genomeSize int
}

func OrderedPopulation(genomeSize int) Creator {
	return &orderedPopulation{
		genomeSize: genomeSize,
	}
}

func (p *orderedPopulation) Create(populationSize int) *Population {
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
