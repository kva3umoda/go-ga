package mutators

import (
	"math/rand"
	"time"
)

//
// Shuffle the attributes of the input individual and return the mutant.
//    The *individual* is expected to be a :term:`sequence`. The *indpb* argument is the
//    probability of each attribute to be moved. Usually this mutation is applied on
//    vector of indices.
type ShuffleIndexes struct {
	prob float64
	rnd  *rand.Rand
}

func NewShuffleIndexes(prob float64) *ShuffleIndexes {
	return &ShuffleIndexes{
		prob: prob,
		rnd:  rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (s *ShuffleIndexes) Mutate(individual []float64) {
	for i := range individual {
		if s.rnd.Float64() < s.prob {
			swapIndx := s.rnd.Intn(len(individual) - 1)
			if swapIndx >= i {
				swapIndx++
			}
			individual[i], individual[swapIndx] = individual[swapIndx], individual[i]
		}
	}
}

// def mutShuffleIndexes(individual, indpb):
//    """Shuffle the attributes of the input individual and return the mutant.
//    The *individual* is expected to be a :term:`sequence`. The *indpb* argument is the
//    probability of each attribute to be moved. Usually this mutation is applied on
//    vector of indices.
//
//    :param individual: Individual to be mutated.
//    :param indpb: Independent probability for each attribute to be exchanged to
//                  another position.
//    :returns: A tuple of one individual.
//
//    This function uses the :func:`~random.random` and :func:`~random.randint`
//    functions from the python base :mod:`random` module.
//    """
//    size = len(individual)
//    for i in xrange(size):
//        if random.random() < indpb:
//            swap_indx = random.randint(0, size - 2)
//            if swap_indx >= i:
//                swap_indx += 1
//            individual[i], individual[swap_indx] = \
//                individual[swap_indx], individual[i]
//
//    return individual,
