package fitness

import "github.com/kva3umoda/go-ga/genome"

type closestValidPenalty struct {
}

func ClosestValidPenalty() Fitness {
	return &closestValidPenalty{}
}

func (c *closestValidPenalty) Fitness(ind *genome.Individual)  {
	panic("implement me")
}
