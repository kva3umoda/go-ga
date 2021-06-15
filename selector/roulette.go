package selectors

import (
	"math/rand"
	"sort"
	"sync"

	"github.com/kva3umoda/go-ga/genome"
)

// Roulette - Select *k* individuals from the input *individuals* using *k* spins of a roulette
type Roulette struct {
	rnd  *rand.Rand
	lock sync.Mutex
}

func NewRoulette() *Roulette {
	return &Roulette{
	}
}

func (r *Roulette) Select(populationSize int,individuals []genome.Individual) []genome.Individual {
	r.lock.Lock()
	defer r.lock.Unlock()

	chosen := make([]genome.Individual, 0, populationSize)

	sortInds := make([]genome.Individual, len(individuals))
	copy(sortInds, individuals)
	// Сортировка по возрастанию
	sort.Slice(sortInds, func(i, j int) bool {
		return sortInds[i].Fitness > sortInds[j].Fitness
	})

	// получаем сумму выживаемости
	var sumFits float64
	for _, ind := range individuals {
		sumFits += ind.Fitness
	}

	for i := 0; i < populationSize; i++ {
		u := rand.Float64() * sumFits
		var sum float64
		for _, ind := range sortInds {
			sum += ind.Fitness
			if sum > u {
				chosen = append(chosen, ind)
				break
			}
		}
	}

	return chosen
}
