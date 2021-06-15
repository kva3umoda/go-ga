package crossover

import (
	"math/rand"
	"sync"
	"time"

	"github.com/kva3umoda/go-ga/genome"
)

// https://proproprogs.ru/ga/ga-obzor-metodov-otbora-skreshchivaniya-i-mutacii
// Executes an ordered crossover (OX) on the input
//    individuals. The two individuals are modified in place. This crossover
//    expects :term:`sequence` individuals of indices, the result for any other
//    type of individuals is unpredictable.
type Ordered struct {
	rnd  *rand.Rand
	lock sync.Mutex
}

func NewOrdered() *Ordered {
	return &Ordered{
		rnd: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

//  TODO: Надо оптимизировать
func (tp *Ordered) Crossing(ind1, ind2 genome.Individual) {
	tp.lock.Lock()
	defer tp.lock.Unlock()

	size := min(len(ind1.Genome), len(ind2.Genome))
	a := tp.rnd.Intn(size)
	b := tp.rnd.Intn(size)
	if a > b {
		a, b = b, a
	}
	holes1 := make([]bool, size)
	holes2 := make([]bool, size)
	for i := a; i <= b; i++ {
		holes1[ind2.GetInt(i)] = true
		holes2[ind1.GetInt(i)] = true
	}

	tmp1 := ind1.Clone()
	tmp2 := ind2.Clone()
	copy(tmp1.Genome[a:b+1], ind2.Genome[a:b+1])
	copy(tmp2.Genome[a:b+1], ind1.Genome[a:b+1])

	k := b + 1
	for j := 0; j < len(tmp1.Genome); j++ {
		if holes1[ind1.GetInt(j)] {
			continue
		}
		tmp1.Genome[k%len(tmp1.Genome)] = ind1.Genome[j]
		k++
	}
	k = b + 1
	for j := 0; j < len(tmp2.Genome); j++ {
		if holes2[ind2.GetInt(j)] {
			continue
		}
		tmp2.Genome[k%len(tmp2.Genome)] = ind2.Genome[j]
		k++
	}

	copy(ind1.Genome, tmp1.Genome)
	copy(ind2.Genome, tmp2.Genome)
}
