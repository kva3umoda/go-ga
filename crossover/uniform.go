package crossover

import (
	"github.com/kva3umoda/go-ga/genome"
	"github.com/kva3umoda/go-ga/rand"
)

// Равномерное скрещивание
type uniform struct {
	indpd float64 // Independent probability for each attribute to be exchanged.
}

func Uniform(indpd float64) Crossover {
	return &uniform{
		indpd: indpd,
	}
}

func (tp *uniform) Crossing(ind1, ind2 *genome.Individual) {
	size := min(len(ind1.Genome), len(ind2.Genome))
	for i := 0; i < size; i++ {
		if rand.Float() < tp.indpd {
			ind1.Genome[i], ind2.Genome[i] = ind2.Genome[i], ind1.Genome[i]
		}
	}
}
