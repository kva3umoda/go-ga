package selectors

import "github.com/kva3umoda/go-ga/genome"

// алгоритмы отбора
type Selector interface {
	Select(populationSize int, individuals []genome.Individual) []genome.Individual
}
