package crossover

import (
	"math/rand"
	"sync"
	"time"

	"github.com/kva3umoda/go-ga/genome"
)



// Равномерное скрещивание
type Uniform struct {
	rnd   *rand.Rand
	lock  sync.Mutex
	indpd float64 // Independent probability for each attribute to be exchanged.
}

func NewUniform(indpd float64) *Uniform {
	return &Uniform{
		rnd:   rand.New(rand.NewSource(time.Now().UnixNano())),
		indpd: indpd,
	}
}

func (tp *Uniform) Crossing(ind1, ind2 genome.Individual) {
	tp.lock.Lock()
	defer tp.lock.Unlock()

	size := min(len(ind1.Genome), len(ind2.Genome))
	for i := 0; i < size; i++ {
		if tp.rnd.Float64() < tp.indpd {
			ind1.Genome[i], ind2.Genome[i] = ind2.Genome[i], ind1.Genome[i]
		}
	}
}

func (tp *Uniform) CrossingCopy(ind1, ind2 genome.Individual) (nind1, nind2 genome.Individual) {
	nind1.Genome, nind2.Genome = copy2(ind1.Genome, ind2.Genome)

	tp.Crossing(nind1, nind2)
	return nind1, nind2
}
