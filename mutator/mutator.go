package mutators

type Mutator interface {
	Mutate(individual []float64)
}
