package mutator

import (
	"math"

	"github.com/kva3umoda/go-ga/genome"
	"github.com/kva3umoda/go-ga/helper"
	"github.com/kva3umoda/go-ga/rand"
)

// //    """Polynomial mutation as implemented in original NSGA-II algorithm in
////    C by Deb.
////
////    :param individual: :term:`Sequence <sequence>` individual to be mutated.
////    :param eta: Crowding degree of the mutation. A high eta will produce
////                a mutant resembling its parent, while a small eta will
////                produce a solution much more different.
////    :param low: A value or a :term:`python:sequence` of values that
////                is the lower bound of the search space.
////    :param up: A value or a :term:`python:sequence` of values that
////               is the upper bound of the search space.
////    :returns: A tuple of one individual.
type polynomialBounded struct {
	eta   float64
	low   []float64
	up    []float64
	indpb float64
}

func PolynomialBounded(eta float64, low []float64, up []float64, indpb float64) Mutator {
	return &polynomialBounded{
		eta:   eta,
		low:   low,
		up:    up,
		indpb: indpb,
	}
}

func (p *polynomialBounded) Mutate(individual *genome.Individual) {
	size := helper.Minia(len(individual.Genome), len(p.low), len(p.up))

	for i := 0; i < size; i++ {
		if rand.Float() > p.indpb {
			continue
		}

		x := individual.Genome[i]
		xl := p.low[i]
		xu := p.up[i]

		delta_1 := (x - xl) / (xu - xl)
		delta_2 := (xu - x) / (xu - xl)
		rnd := rand.Float()
		mut_pow := 1.0 / (p.eta + 1.0)

		var xy, val, delta_q float64
		if rnd < 0.5 {
			xy = 1.0 - delta_1
			val = 2.0*rnd + (1.0-2.0*rnd)*math.Pow(xy, p.eta+1.0)
			delta_q = math.Pow(val, mut_pow-1.0)
		} else {
			xy = 1.0 - delta_2
			val = 2.0*(1.0-rnd)+2.0*(rnd-0.5)*math.Pow(xy, p.eta+1.0)
			delta_q = math.Pow(1.0-val, mut_pow)
		}

		x += delta_q * (xu - xl)
		x = helper.Minf(helper.Maxf(x, xl), xu)
		individual.Genome[i] = x
	}
}
