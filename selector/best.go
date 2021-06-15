package selector

import (
	"sort"

	"github.com/kva3umoda/go-ga/genome"
)

// Best - Select the *k* best individuals among the input *individuals*
type Best struct {
}

func NewBest() *Best {
	return &Best{
	}
}

func (b *Best) Select(populationSize int, individuals []genome.Individual) []genome.Individual {
	chosen := make([]genome.Individual, len(individuals))
	copy(chosen, individuals)
	sort.Slice(chosen, func(i, j int) bool {
		return chosen[i].Fitness > chosen[j].Fitness
	})
	return chosen[:populationSize]
}
