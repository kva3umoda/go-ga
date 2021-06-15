package mutators

import (
	"math/rand"
	"sync"
	"time"
)

// Flip the value of the attributes of the input individual and return the
// mutant. The *individual* is expected to be a :term:`sequence` and the values of the
// attributes shall stay valid after the ``not`` operator is called on them.
// The *indpb* argument is the probability of each attribute to be
// flipped. This mutation is usually applied on boolean individuals.
type FlitBit struct {
	rnd   *rand.Rand
	indpb float64 // Independent probability for each attribute to be flipped.
	lock  sync.Mutex
}

func NewFlitBit(indpb float64) *FlitBit {
	return &FlitBit{
		rnd:   rand.New(rand.NewSource(time.Now().UnixNano())),
		indpb: indpb,
	}
}

func (f *FlitBit) Mutate(individual []float64) {
	f.lock.Lock()
	defer f.lock.Unlock()
	for i := range individual {
		if f.rnd.Float64() < f.indpb {
			if individual[i] > 0.0 {
				individual[i] = 0.0
			} else {
				individual[i] = 1.0
			}
		}
	}
}
