package crossover

import (
	"github.com/kva3umoda/go-ga/genome"
	"github.com/kva3umoda/go-ga/helper"
	"github.com/kva3umoda/go-ga/rand"
)

// Executes a uniform partially matched crossover (UPMX) on the input
//    individuals. The two individuals are modified in place. This crossover
//    expects :term:`sequence` individuals of indices, the result for any other
//    type of individuals is unpredictable
type uniformPartialyMatched struct {
	prob float64
}

func UniformPartialMatched(prob float64) Crossover {
	return &uniformPartialyMatched{
		prob: prob,
	}
}

func (tp *uniformPartialyMatched) Crossing(ind1, ind2 *genome.Individual) {
	size := helper.Mini(len(ind1.Genome), len(ind2.Genome))
	p1 := make([]float64, size)
	p2 := make([]float64, size)
	// Initialize the position of each indices in the individuals
	for i := 0; i < size; i++ {
		p1[int(ind1.Genome[i])] = float64(i)
		p2[int(ind2.Genome[i])] = float64(i)
	}

	for i := 0; i < size; i++ {
		if rand.Float() < tp.prob {
			// Keep track of the selected values
			temp1 := ind1.Genome[i]
			temp2 := ind2.Genome[i]
			// Swap the matched value
			ind1.Genome[i], ind1.Genome[int(p1[int(temp2)])] = temp2, temp1
			ind2.Genome[i], ind2.Genome[int(p2[int(temp1)])] = temp1, temp2
			// Position bookkeeping
			p1[int(temp1)], p1[int(temp2)] = p1[int(temp2)], p1[int(temp1)]
			p2[int(temp1)], p2[int(temp2)] = p2[int(temp2)], p2[int(temp1)]
		}
	}

}

// def cxUniformPartialyMatched(ind1, ind2, indpb):
//    """Executes a uniform partially matched crossover (UPMX) on the input
//    individuals. The two individuals are modified in place. This crossover
//    expects :term:`sequence` individuals of indices, the result for any other
//    type of individuals is unpredictable.
//
//    :param ind1: The first individual participating in the crossover.
//    :param ind2: The second individual participating in the crossover.
//    :returns: A tuple of two individuals.
//
//    Moreover, this crossover generates two children by matching
//    pairs of values chosen at random with a probability of *indpb* in the two
//    parents and swapping the values of those indexes. For more details see
//    [Cicirello2000]_.
//
//    This function uses the :func:`~random.random` and :func:`~random.randint`
//    functions from the python base :mod:`random` module.
//
//    .. [Cicirello2000] Cicirello and Smith, "Modeling GA performance for
//       control parameter optimization", 2000.
//    """
//    size = Mini(len(ind1), len(ind2))
//    p1, p2 = [0] * size, [0] * size
//
//    # Initialize the position of each indices in the individuals
//    for i in xrange(size):
//        p1[ind1[i]] = i
//        p2[ind2[i]] = i
//
//    for i in xrange(size):
//        if random.random() < indpb:
//            # Keep track of the selected values
//            temp1 = ind1[i]
//            temp2 = ind2[i]
//            # Swap the matched value
//            ind1[i], ind1[p1[temp2]] = temp2, temp1
//            ind2[i], ind2[p2[temp1]] = temp1, temp2
//            # Position bookkeeping
//            p1[temp1], p1[temp2] = p1[temp2], p1[temp1]
//            p2[temp1], p2[temp2] = p2[temp2], p2[temp1]
//
//    return ind1, ind2
