package crossover

import (
	"github.com/kva3umoda/go-ga/genome"
	"github.com/kva3umoda/go-ga/rand"
)

// Executes a one point crossover on the input :term:`sequence` individuals.
//    The two individuals are modified in place. The resulting individuals will
//    respectively have the length of the other
type onePoint struct {}

func OnePoint() Crossover {
	return &onePoint{
	}
}

func (op *onePoint) Crossing(ind1, ind2 *genome.Individual) {

	size := min(len(ind1.Genome), len(ind2.Genome))

	cxpoint := rand.Int(size-1) + 1 //  random.randint(1, size - 1)

	tmp := make([]float64, cxpoint)
	copy(tmp, ind1.Genome[:cxpoint])

	copy(ind1.Genome[:cxpoint], ind2.Genome[:cxpoint])
	copy(ind2.Genome[:cxpoint], tmp)
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
