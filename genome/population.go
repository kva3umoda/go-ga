package genome

// Популяция
type Population struct {
	Generation  int          // Поколение
	Individuals []Individual // индивидумы
}

func newPopulation(populationSize int) *Population {
	return &Population{
		Generation:  0,
		Individuals: make([]Individual, 0, populationSize),
	}
}

func (p *Population) Clone() *Population {
	pop := new(Population)
	pop.Generation = p.Generation
	pop.Individuals = make([]Individual, len(p.Individuals))
	for i := range p.Individuals {
		pop.Individuals[i] = p.Individuals[i].Clone()
	}
	return pop
}
