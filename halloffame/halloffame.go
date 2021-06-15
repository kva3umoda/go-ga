package halloffame

import (
	"github.com/kva3umoda/go-ga/genome"
	"github.com/kva3umoda/go-ga/population"
)

type HallOfFame interface {
	Update(population *population.Population)
	Individuals() []*genome.Individual
}
