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

//eta Crowding degree of the crossover. A high eta will produce
//                children resembling to their parents, while a small eta will
//                produce solutions much more different.
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
		xl := s.low[i]
		xu := s.up[i]

		x1 := helper.Minf(ind1.Genome[i], ind2.Genome[i])
		x2 := helper.Maxf(ind1.Genome[i], ind2.Genome[i])
		rnd := rand.Float()

		beta := 1.0 + (2.0 * (x1 - xl) / (x2 - x1))
		alpha := 2.0 - math.Pow(beta, -(s.eta+1.0))
		beta_q := 0.0
		if rnd <= 1.0/alpha {
			beta_q = math.Pow(rnd*alpha, 1.0/(s.eta+1.0))
		} else {
			beta_q = math.Pow(1.0/(2.0-rnd*alpha), 1.0/(s.eta+1))
		}
		c1 := 0.5 * (x1 + x2 - beta_q*(x2-x1))

		beta = 1.0 + (2.0 * (xu - x2) / (x2 - x1))
		alpha = 2.0 - math.Pow(beta, -(s.eta+1))
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
