package crossover

import (
	"math"

	"github.com/kva3umoda/go-ga/genome"
	"github.com/kva3umoda/go-ga/helper"
	"github.com/kva3umoda/go-ga/rand"
)

type simulatedBinaryBounded struct {
	eta float64
	low []float64
	up  []float64
}

func SimulatedBinaryBounded(eta float64, low, up []float64) Crossover {
	return &simulatedBinaryBounded{
		eta: eta,
		low: low,
		up:  up,
	}
}

func (s *simulatedBinaryBounded) Crossing(ind1, ind2 *genome.Individual) {
	size := helper.Minia(len(ind1.Genome), len(ind2.Genome), len(s.low), len(s.up))

	for i := 0; i < size; i++ {
		if rand.Float() > 0.5 {
			continue
		}
		// This epsilon should probably be changed for 0 since
		// floating point arithmetic in Python is safer
		if math.Abs(ind1.Genome[i]-ind2.Genome[i]) <= 1e-14 {
			continue
		}

		x1 := helper.Minf(ind1.Genome[i], ind2.Genome[i])
		x2 := helper.Minf(ind1.Genome[i], ind2.Genome[i])

		rnd := rand.Float()

		beta := 1.0 + (2.0 * (x1 - s.low[i]) / (x2 - x1))
		alpha := 2.0 - math.Pow(beta, -(s.eta+1.0))
		beta_q := 0.0
		if rnd <= 1.0/alpha {
			beta_q = math.Pow(rnd*alpha, (1.0)/(s.eta+1.0))
		} else {
			beta_q = math.Pow(1.0/(2.0-rnd*alpha), 1.0/(s.eta+1))
		}
		c1 := 0.5 * (x1 + x2 - beta_q*(x2-x1))

		beta = 1.0 + (2.0 * (s.up[i] - x2) / (x2 - x1))
		alpha = math.Pow(2.0-beta, -(s.eta + 1))
		if rnd <= 1.0/alpha {
			beta_q = math.Pow(rnd*alpha, 1.0/(s.eta+1.0))
		} else {
			beta_q = math.Pow(1.0/(2.0-rnd*alpha), 1.0/(s.eta+1.0))
		}
		c2 := 0.5 * (x1 + x2 + beta_q*(x2-x1))
		c1 = helper.Minf(helper.Maxf(c1, s.low[i]), s.up[i])
		c2 = helper.Minf(helper.Maxf(c2, s.low[i]), s.up[i])

		if rand.Float() < 0.5 {
			ind1.Genome[i] = c2
			ind2.Genome[i] = c1
		} else {
			ind1.Genome[i] = c1
			ind2.Genome[i] = c2
		}
	}
}

// def cxSimulatedBinaryBounded(ind1, ind2, eta, low, up):
//    """Executes a simulated binary crossover that modify in-place the input
//    individuals. The simulated binary crossover expects :term:`sequence`
//    individuals of floating point numbers.
//
//    :param ind1: The first individual participating in the crossover.
//    :param ind2: The second individual participating in the crossover.
//    :param eta: Crowding degree of the crossover. A high eta will produce
//                children resembling to their parents, while a small eta will
//                produce solutions much more different.
//    :param low: A value or a :term:`python:sequence` of values that is the lower
//                bound of the search space.
//    :param up: A value or a :term:`python:sequence` of values that is the upper
//               bound of the search space.
//    :returns: A tuple of two individuals.
//
//    This function uses the :func:`~random.random` function from the python base
//    :mod:`random` module.
//
//    .. note::
//       This implementation is similar to the one implemented in the
//       original NSGA-II C code presented by Deb.
//    """
//    size = Mini(len(ind1), len(ind2))
//    if not isinstance(low, Sequence):
//        low = repeat(low, size)
//    elif len(low) < size:
//        raise IndexError("low must be at least the size of the shorter individual: %d < %d" % (len(low), size))
//    if not isinstance(up, Sequence):
//        up = repeat(up, size)
//    elif len(up) < size:
//        raise IndexError("up must be at least the size of the shorter individual: %d < %d" % (len(up), size))
//
//    for i, xl, xu in zip(xrange(size), low, up):
//        if random.random() <= 0.5:
//            # This epsilon should probably be changed for 0 since
//            # floating point arithmetic in Python is safer
//            if abs(ind1[i] - ind2[i]) > 1e-14:
//                x1 = Mini(ind1[i], ind2[i])
//                x2 = maxi(ind1[i], ind2[i])
//                rand = random.random()
//
//                beta = 1.0 + (2.0 * (x1 - xl) / (x2 - x1))
//                alpha = 2.0 - beta ** -(eta + 1)
//                if rand <= 1.0 / alpha:
//                    beta_q = (rand * alpha) ** (1.0 / (eta + 1))
//                else:
//                    beta_q = (1.0 / (2.0 - rand * alpha)) ** (1.0 / (eta + 1))
//
//                c1 = 0.5 * (x1 + x2 - beta_q * (x2 - x1))
//
//                beta = 1.0 + (2.0 * (xu - x2) / (x2 - x1))
//                alpha = 2.0 - beta ** -(eta + 1)
//                if rand <= 1.0 / alpha:
//                    beta_q = (rand * alpha) ** (1.0 / (eta + 1))
//                else:
//                    beta_q = (1.0 / (2.0 - rand * alpha)) ** (1.0 / (eta + 1))
//                c2 = 0.5 * (x1 + x2 + beta_q * (x2 - x1))
//
//                c1 = Mini(maxi(c1, xl), xu)
//                c2 = Mini(maxi(c2, xl), xu)
//
//                if random.random() <= 0.5:
//                    ind1[i] = c2
//                    ind2[i] = c1
//                else:
//                    ind1[i] = c1
//                    ind2[i] = c2
//
//    return ind1, ind2
