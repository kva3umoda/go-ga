package mutator

import (
	"math"

	"github.com/kva3umoda/go-ga/genome"
	"github.com/kva3umoda/go-ga/helper"
	"github.com/kva3umoda/go-ga/rand"
)

// Polynomial mutation as implemented in original NSGA-II algorithm in
//    C by Deb
type polynomialBounded struct {
	eta   float64
	low   []float64
	up    []float64
	indpb float64
}

func PolynomialBounded(eta float64, low []float64, up []float64, indpb float64) Mutator {
	return &polynomialBounded{}
}

func (p *polynomialBounded) Mutate(individual *genome.Individual) {
	size := helper.Minia(len(individual.Genome), len(p.low), len(p.up))

	for i := 0; i < size; i++ {
		if rand.Float() > p.indpb {
			continue
		}

		x := individual.Genome[i]
		delta_1 := (x - p.low[i]) / (p.up[i] - p.low[i])
		delta_2 := (p.up[i] - x) / (p.up[i] - p.low[i])
		rnd := rand.Float()
		mut_pow := 1.0 / (p.eta + 1.0)

		var xy, val, delta_q float64
		if rnd < 0.5 {
			xy = 1.0 - delta_1
			val = math.Pow(2.0*rnd+(1.0-2.0*rnd)*xy, p.eta+1.0)
			delta_q = math.Pow(val, mut_pow-1.0)
		} else {
			xy = 1.0 - delta_2
			val = math.Pow(2.0*(1.0-rnd)+2.0*(rnd-0.5)*xy, p.eta+1.0)
			delta_q = math.Pow(1.0-val, mut_pow)
		}

		x += delta_q * (p.up[i] - p.low[i])
		x = helper.Minf(helper.Maxf(x, p.low[i]), p.up[i])
		individual.Genome[i] = x
	}
}

// def mutPolynomialBounded(individual, eta, low, up, indpb):
//    """Polynomial mutation as implemented in original NSGA-II algorithm in
//    C by Deb.
//
//    :param individual: :term:`Sequence <sequence>` individual to be mutated.
//    :param eta: Crowding degree of the mutation. A high eta will produce
//                a mutant resembling its parent, while a small eta will
//                produce a solution much more different.
//    :param low: A value or a :term:`python:sequence` of values that
//                is the lower bound of the search space.
//    :param up: A value or a :term:`python:sequence` of values that
//               is the upper bound of the search space.
//    :returns: A tuple of one individual.
//    """
//    size = len(individual)
//    if not isinstance(low, Sequence):
//        low = repeat(low, size)
//    elif len(low) < size:
//        raise IndexError("low must be at least the size of individual: %d < %d" % (len(low), size))
//    if not isinstance(up, Sequence):
//        up = repeat(up, size)
//    elif len(up) < size:
//        raise IndexError("up must be at least the size of individual: %d < %d" % (len(up), size))
//
//    for i, xl, xu in zip(xrange(size), low, up):
//        if random.random() <= indpb:
//            x = individual[i]
//            delta_1 = (x - xl) / (xu - xl)
//            delta_2 = (xu - x) / (xu - xl)
//            rand = random.random()
//            mut_pow = 1.0 / (eta + 1.)
//
//            if rand < 0.5:
//                xy = 1.0 - delta_1
//                val = 2.0 * rand + (1.0 - 2.0 * rand) * xy ** (eta + 1)
//                delta_q = val ** mut_pow - 1.0
//            else:
//                xy = 1.0 - delta_2
//                val = 2.0 * (1.0 - rand) + 2.0 * (rand - 0.5) * xy ** (eta + 1)
//                delta_q = 1.0 - val ** mut_pow
//
//            x = x + delta_q * (xu - xl)
//            x = min(max(x, xl), xu)
//            individual[i] = x
//    return individual,
