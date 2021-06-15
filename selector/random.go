package selector

import (
	"math/rand"
	"sync"
	"time"

	"github.com/kva3umoda/go-ga/genome"
)

type Random struct {
	rnd  *rand.Rand
	lock sync.Mutex
}

func NewRandom() *Random {
	return &Random{
		rnd: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (r *Random) Select(populationSize int, individuals []genome.Individual) []genome.Individual {
	r.lock.Lock()
	defer r.lock.Unlock()
	// TODO: в оригинале один индивидума может возвращаться два раза. надо проверить так ли надо [random.choice(individuals) for i in xrange(k)]
	chosen := make([]genome.Individual, populationSize)
	copy(chosen, individuals)
	rand.Shuffle(len(chosen), func(i, j int) {
		chosen[i], chosen[j] = chosen[j], chosen[i]
	})

	return chosen[:populationSize]
}
