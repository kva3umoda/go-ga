package population

import "github.com/kva3umoda/go-ga/genome"

// Популяция
type Population struct {
	Generation  int                  // Поколение
	Individuals []*genome.Individual // индивидумы
}

func newPopulation(populationSize int) *Population {
	return &Population{
		Generation:  0,
		Individuals: make([]*genome.Individual, 0, populationSize),
	}
}

func (p *Population) Clone() *Population {
	pop := new(Population)
	pop.Generation = p.Generation
	pop.Individuals = make([]*genome.Individual, len(p.Individuals))
	for i := range p.Individuals {
		pop.Individuals[i] = p.Individuals[i].Clone()
	}
	return pop
}
