package selectors

import (
	"sort"

	"github.com/kva3umoda/go-ga/genome"
)

// Worst - Select the *k* worst individuals among the input *individuals*
type Worst struct {
}

func NewWorst() *Worst {
	return &Worst{}
}

func (b *Worst) Select(populationSize int, individuals []genome.Individual) []genome.Individual {
	chosen := make([]genome.Individual, len(individuals))
	copy(chosen, individuals)
	sort.Slice(chosen, func(i, j int) bool {
		return chosen[i].Fitness < chosen[j].Fitness
	})
	return chosen[:populationSize]
}
