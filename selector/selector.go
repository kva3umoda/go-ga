package selector

import "github.com/kva3umoda/go-ga/genome"

type SelectFunc func(populationSize int, individuals []*genome.Individual) []*genome.Individual

// алгоритмы отбора
type Selector interface {
	Select(populationSize int, individuals []*genome.Individual) []*genome.Individual
}

type Function struct {
	selectFunc SelectFunc
}

func NewFunction(selectFunc SelectFunc) *Function {
	return &Function{
		selectFunc: selectFunc,
	}
}

func (f *Function) Select(populationSize int, individuals []*genome.Individual) []*genome.Individual {
	return f.selectFunc(populationSize, individuals)
}