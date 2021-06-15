package mutator

import "github.com/kva3umoda/go-ga/genome"

type MutatorFunc func(individual *genome.Individual)

type Mutator interface {
	Mutate(individual *genome.Individual)
}

type Function struct {
	mutatorFunc MutatorFunc
}

func NewFunction(mutatorFunc MutatorFunc) *Function {
	return &Function{
		mutatorFunc: mutatorFunc,
	}
}

func (f *Function) Mutate(individual *genome.Individual) {
	f.mutatorFunc(individual)
}
