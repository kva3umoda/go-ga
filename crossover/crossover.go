package crossover

import "github.com/kva3umoda/go-ga/genome"

type Crossover interface {
	Crossing(ind1, ind2 genome.Individual)
}
