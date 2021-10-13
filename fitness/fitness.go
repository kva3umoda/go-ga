package fitness

import "github.com/kva3umoda/go-ga/genome"

type FitnessFunc func(ind *genome.Individual)

type Fitness interface {
	Fitness(ind *genome.Individual)
}

type function struct {
	fitnessFunc FitnessFunc
}

func Function(fitnessFunc FitnessFunc) Fitness {
	return &function{
		fitnessFunc: fitnessFunc,
	}
}

func (f *function) Fitness(ind *genome.Individual) {
	f.fitnessFunc(ind)
}

// default fitness

type defaultFitness struct {
	weight float64
}

func DefaultFitness(weight float64) Fitness {
	return &defaultFitness{
		weight: weight,
	}
}

func (df *defaultFitness) Fitness(ind *genome.Individual) {
	ind.Fitness = ind.Cost * df.weight
}

// max fitness
func Max() Fitness {
	return DefaultFitness(1.0)
}

// min fitness
func Min() Fitness {
	return DefaultFitness(-1.0)
}
