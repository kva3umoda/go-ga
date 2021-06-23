package population

import (
	"github.com/kva3umoda/go-ga/genome"
	"github.com/kva3umoda/go-ga/rand"
)

type binaryPopulation struct {
	genomeSize int
}

func BinaryPopulation(genomeSize int) Creator {
	return &binaryPopulation{
		genomeSize: genomeSize,
	}
}

func (p *binaryPopulation) Create(populationSize int) *Population {
	pop := newPopulation(populationSize)
	for i := 0; i < populationSize; i++ {
		ind := genome.NewIndividual(p.genomeSize)
		for j := range ind.Genome {
			if rand.Float() >= 0.5 {
				ind.Genome[j] = 1.0
			}
		}
		pop.Individuals = append(pop.Individuals, ind)
	}
	return pop
}
