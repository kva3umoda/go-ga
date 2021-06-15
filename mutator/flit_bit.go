package mutator

import (
	"github.com/kva3umoda/go-ga/genome"
	"github.com/kva3umoda/go-ga/rand"
)

// Flip the value of the attributes of the input individual and return the
// mutant. The *individual* is expected to be a :term:`sequence` and the values of the
// attributes shall stay valid after the ``not`` operator is called on them.
// The *indpb* argument is the probability of each attribute to be
// flipped. This mutation is usually applied on boolean individuals.
type FlitBit struct {
	indpb float64 // Independent probability for each attribute to be flipped.
}

func NewFlitBit(indpb float64) *FlitBit {
	return &FlitBit{
		indpb: indpb,
	}
}

func (f *FlitBit) Mutate(individual *genome.Individual) {
	for i := range individual.Genome {
		if rand.Float() < f.indpb {
			if individual.Genome[i] > 0.0 {
				individual.Genome[i] = 0.0
			} else {
				individual.Genome[i] = 1.0
			}
		}
	}
}
