package halloffame

import "github.com/kva3umoda/go-ga/genome"

// TODO : покрыть тестами
type Base struct {
	maxSize     int
	individuals []genome.Individual
}

func NewBase(maxSize int) *Base {
	return &Base{
		maxSize:     maxSize,
		individuals: make([]genome.Individual, 0, maxSize+1),
	}
}

func (hof *Base) Update(population *genome.Population) {
	if hof.maxSize == 0 {
		return
	}

	for _, ind := range population.Individuals {
		if len(hof.individuals) == 0 {
			hof.insert(ind)
			continue
		}
		if hof.exist(ind) {
			continue
		}
		hof.insert(ind)
	}
}

func (hof *Base) Individuals() []genome.Individual {
	return hof.individuals
}

func (hof *Base) insert(item genome.Individual) {
	item = item.Clone()

	var n int

	for i, ind := range hof.individuals {
		n = i
		if item.Fitness >= ind.Fitness {
			break
		}
	}

	if n+1 == len(hof.individuals) && item.Fitness < hof.individuals[len(hof.Individuals())-1].Fitness {
		return
	}

	hof.individuals = append(hof.individuals, genome.Individual{})
	copy(hof.individuals[n+1:], hof.individuals[n:])
	hof.individuals[n] = item

	if len(hof.individuals) > hof.maxSize {
		hof.individuals = hof.individuals[:hof.maxSize]
	}
}

// Прверяет на существование индивидума
func (hof *Base) exist(cmp genome.Individual) bool {
	for _, ind := range hof.individuals {
		if cmp.Equal(ind) {
			return true
		}
	}

	return false
}
