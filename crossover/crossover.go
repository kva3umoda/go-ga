package crossover

import "github.com/kva3umoda/go-ga/genome"

type CrossoverFunc func(ind1, ind2 *genome.Individual)

type Crossover interface {
	Crossing(ind1, ind2 *genome.Individual)
}

type Function struct {
	crossoverFunc CrossoverFunc
}

func NewFunction(crossoverFunc CrossoverFunc) *Function {
	return &Function{
		crossoverFunc: crossoverFunc,
	}
}

func (f *Function) Crossing(ind1, ind2 *genome.Individual) {
	f.crossoverFunc(ind1, ind2)
}
