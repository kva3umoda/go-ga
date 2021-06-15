package crossover

import (
	"math/rand"
	"sync"
	"time"

	"github.com/kva3umoda/go-ga/genome"
)

// Executes a one point crossover on the input :term:`sequence` individuals.
//    The two individuals are modified in place. The resulting individuals will
//    respectively have the length of the other
type OnePoint struct {
	rnd  *rand.Rand
	lock sync.Mutex
}

func NewOnePoint() *OnePoint {
	return &OnePoint{
		rnd: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (op *OnePoint) Crossing(ind1, ind2 genome.Individual) {
	op.lock.Lock()
	defer op.lock.Unlock()

	size := min(len(ind1.Genome), len(ind2.Genome))

	cxpoint := op.rnd.Intn(size-1) + 1 //  random.randint(1, size - 1)

	tmp := make([]float64, cxpoint)
	copy(tmp, ind1.Genome[:cxpoint])

	copy(ind1.Genome[:cxpoint], ind2.Genome[:cxpoint])
	copy(ind2.Genome[:cxpoint], tmp)
}

func (op *OnePoint) CrossingCopy(ind1, ind2 genome.Individual) (nind1, nind2 genome.Individual) {
	nind1.Genome, nind2.Genome = copy2(ind1.Genome, ind2.Genome)

	op.Crossing(nind1, nind2)
	return
}

// def cxOnePoint(ind1, ind2):
//    """Executes a one point crossover on the input :term:`sequence` individuals.
//    The two individuals are modified in place. The resulting individuals will
//    respectively have the length of the other.
//
//    :param ind1: The first individual participating in the crossover.
//    :param ind2: The second individual participating in the crossover.
//    :returns: A tuple of two individuals.
//
//    This function uses the :func:`~random.randint` function from the
//    python base :mod:`random` module.
//    """
//    size = min(len(ind1), len(ind2))
//    cxpoint = random.randint(1, size - 1)
//    ind1[cxpoint:], ind2[cxpoint:] = ind2[cxpoint:], ind1[cxpoint:]
//
//    return ind1, ind2
