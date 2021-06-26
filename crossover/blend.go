package crossover

import (
	"github.com/kva3umoda/go-ga/genome"
	"github.com/kva3umoda/go-ga/helper"
	"github.com/kva3umoda/go-ga/rand"
)

// Скрещивание смешением (Blend Crossover – BLX), когда каждый по томок случайным образом выбирается из интервала, образованного родителями:
// [parent1 – α(parent2 – parent1), parent2 + α(parent2 – parent1)].
// Коэффеициент alpha обычно принимается равным 0,5
type blend struct {
	alpha float64
}

func Blend(alpha float64) Crossover {
	return &blend{alpha: alpha}
}

func (b *blend) Crossing(ind1, ind2 *genome.Individual) {
	size := helper.Mini(len(ind1.Genome), len(ind2.Genome))
	for i := 0; i < size; i++ {
		gamma := (1+2*b.alpha)*rand.Float() - b.alpha
		ind1.Genome[i] = (1-gamma)*ind1.Genome[i] + gamma*ind2.Genome[i]
		ind2.Genome[i] = gamma*ind1.Genome[i] + (1-gamma)*ind2.Genome[i]
	}
}