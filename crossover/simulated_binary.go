package crossover

import (
	"math"

	"github.com/kva3umoda/go-ga/genome"
	"github.com/kva3umoda/go-ga/helper"
	"github.com/kva3umoda/go-ga/rand"
)

// Имитация двоичного скрещивания (Simulated Binary Crossover SBX),
// когда два потомка порождаются по следующей формуле, гарантирующей, что среднее значение потомков равно среднему значению
//родителей:
//offsping1 = ½ [(1 + β]parent1 + (1 – β)parent2]);
//offsping2 = ½ [(1 – β]parent1 + (1 + β)parent2]).
//Коэффициент β, называемый коэффициентом распределения, вычисляется в виде комбинации случайно выбранного значения и зара-
// нее заданного параметра η, называемого индексом распределения, или выбирается в диапазаон от 10 до 20
type simulatedBinary struct {
	eta float64
}

func SimulatedBinary(eta float64) Crossover {
	return &simulatedBinary{
		eta: eta,
	}
}

func (s *simulatedBinary) Crossing(ind1, ind2 *genome.Individual) {
	size := helper.Mini(len(ind1.Genome), len(ind2.Genome))
	for i := 0; i < size; i++ {
		rnd := rand.Float()
		beta := 0.0
		if rnd <= 0.5 {
			beta = 2.0 * rnd
		} else {
			beta = 1 / (2 * (1 - rnd))
		}
		beta = math.Pow(beta, 1.0/(2.0*(1.0-rnd)))
		ind1.Genome[i] = 0.5 * (((1.0 + beta) * ind1.Genome[i]) + ((1.0 - beta) * ind2.Genome[i]))
		ind2.Genome[i] = 0.5 * (((1.0 - beta) * ind1.Genome[i]) + ((1.0 + beta) * ind2.Genome[i]))
	}
}

// def cxSimulatedBinary(ind1, ind2, eta):
//    """Executes a simulated binary crossover that modify in-place the input
//    individuals. The simulated binary crossover expects :term:`sequence`
//    individuals of floating point numbers.
//
//    :param ind1: The first individual participating in the crossover.
//    :param ind2: The second individual participating in the crossover.
//    :param eta: Crowding degree of the crossover. A high eta will produce
//                children resembling to their parents, while a small eta will
//                produce solutions much more different.
//    :returns: A tuple of two individuals.
//
//    This function uses the :func:`~random.random` function from the python base
//    :mod:`random` module.
//    """
//    for i, (x1, x2) in enumerate(zip(ind1, ind2)):
//        rand = random.random()
//        if rand <= 0.5:
//            beta = 2. * rand
//        else:
//            beta = 1. / (2. * (1. - rand))
//        beta **= 1. / (eta + 1.)
//        ind1[i] = 0.5 * (((1 + beta) * x1) + ((1 - beta) * x2))
//        ind2[i] = 0.5 * (((1 - beta) * x1) + ((1 + beta) * x2))
//
//    return ind1, ind2
