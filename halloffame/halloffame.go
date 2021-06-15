package halloffame

import "ga-book/internal/genome"

type HallOfFame interface {
	Update(population *genome.Population)
	Individuals() []genome.Individual
}
