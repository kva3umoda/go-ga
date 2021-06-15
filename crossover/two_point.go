package crossover

import (
	"math/rand"
	"sync"
	"time"

	"github.com/kva3umoda/go-ga/genome"
)

// Executes a two-point crossover on the input :term:`sequence`
//   individuals. The two individuals are modified in place and both keep
//   their original length
type TwoPoint struct {
	rnd  *rand.Rand
	lock sync.Mutex
}

func NewTwoPoint() *TwoPoint {
	return &TwoPoint{
		rnd: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (tp *TwoPoint) Crossing(ind1, ind2 genome.Individual) {
	tp.lock.Lock()
	defer tp.lock.Unlock()

	size := min(len(ind1.Genome), len(ind2.Genome))
	cxpoint1 := tp.rnd.Intn(size)       // random.randint(1, size)
	cxpoint2 := tp.rnd.Intn(size-1) + 1 //  random.randint(1, size - 1)
	//fmt.Printf("#1 cxpoint1: %d, cxpoint2: %d\n", cxpoint1, cxpoint2)
	if cxpoint2 >= cxpoint1 {
		cxpoint2++
	} else {
		cxpoint1, cxpoint2 = cxpoint2, cxpoint1
	}
	//fmt.Printf("#2 cxpoint1: %d, cxpoint2: %d\n", cxpoint1, cxpoint2)

	tmp := make([]float64, cxpoint2-cxpoint1)
	copy(tmp, ind1.Genome[cxpoint1:cxpoint2])

	copy(ind1.Genome[cxpoint1:cxpoint2], ind2.Genome[cxpoint1:cxpoint2])
	copy(ind2.Genome[cxpoint1:cxpoint2], tmp)
}

// def cxTwoPoint(ind1, ind2):
//    """Executes a two-point crossover on the input :term:`sequence`
//    individuals. The two individuals are modified in place and both keep
//    their original length.
//
//    :param ind1: The first individual participating in the crossover.
//    :param ind2: The second individual participating in the crossover.
//    :returns: A tuple of two individuals.
//
//    This function uses the :func:`~random.randint` function from the Python
//    base :mod:`random` module.
//    """
//    size = min(len(ind1), len(ind2))
//    cxpoint1 = random.randint(1, size)
//    cxpoint2 = random.randint(1, size - 1)
//    if cxpoint2 >= cxpoint1:
//        cxpoint2 += 1
//    else:  # Swap the two cx points
//        cxpoint1, cxpoint2 = cxpoint2, cxpoint1
//
//    ind1[cxpoint1:cxpoint2], ind2[cxpoint1:cxpoint2] \
//        = ind2[cxpoint1:cxpoint2], ind1[cxpoint1:cxpoint2]
//
//    return ind1, ind2
