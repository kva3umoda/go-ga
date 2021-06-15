package mutators

import (
	"math/rand"
)

type Multi struct {
	mutators []Mutator
}

func NewMulti(mutators []Mutator) *Multi {
	return &Multi{mutators: mutators}
}

func (m *Multi) Mutate(individual []float64) {
	r := rand.Intn(len(individual))
	m.mutators[r].Mutate(individual)
}
