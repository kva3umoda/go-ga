package mutator

import (
	"github.com/kva3umoda/go-ga/genome"
	"github.com/kva3umoda/go-ga/rand"
)

//
// Shuffle the attributes of the input individual and return the mutant.
//    The *individual* is expected to be a :term:`sequence`. The *indpb* argument is the
//    probability of each attribute to be moved. Usually this mutation is applied on
//    vector of indices.
type shuffleIndexes struct {
	prob float64
}

func ShuffleIndexes(prob float64) Mutator {
	return &shuffleIndexes{
		prob: prob,
	}
}

func (s *shuffleIndexes) Mutate(individual *genome.Individual) {
	for i := range individual.Genome {
		if rand.Float() < s.prob {
			swapIndx := rand.Int(len(individual.Genome) - 1)
			if swapIndx >= i {
				swapIndx++
			}
			individual.Genome[i], individual.Genome[swapIndx] = individual.Genome[swapIndx], individual.Genome[i]
		}
	}
}
