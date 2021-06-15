package mutator

import (
	"github.com/kva3umoda/go-ga/genome"
	"github.com/kva3umoda/go-ga/rand"
)

type Multi struct {
	mutators []Mutator
}

func NewMulti(mutators []Mutator) *Multi {
	return &Multi{
		mutators: mutators,
	}
}

func (m *Multi) Mutate(individual *genome.Individual) {
	r := rand.Int(len(m.mutators))
	m.mutators[r].Mutate(individual)
}
