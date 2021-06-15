package main

import (
	"ga-book/examples/helper"
	"ga-book/internal"
	"ga-book/internal/crossovers"
	"ga-book/internal/mutators"
	"ga-book/internal/population"
	"ga-book/internal/selector"
)

const (
	ONE_MAX_LENGTH = 100.0 // Длина подлежащей оптимизации битовой строки

	// Константы генетического алгоритма
	POPULATION_SIZE = 200 // количество индивидуумов в популяции
	P_CROSSOVER     = 0.9 // вероятность скрещивания
	P_MUTATION      = 0.1 // вероятность мутации индивидуумов
	MAX_GENERATIONS = 50  // максимальное количество поколений
)

func main() {
	conf := internal.NewConfig().
		Population(POPULATION_SIZE, population.NewPopulationBinary(ONE_MAX_LENGTH)).
		CostFunction(1.0, oneMaxFitness).
		Selector(selector.NewTournament(3)).
		Crossover(P_CROSSOVER, crossovers.NewOnePoint()).
		Mutator(P_MUTATION, mutators.NewFlitBit(1.0/ONE_MAX_LENGTH)).
		Generation(MAX_GENERATIONS)

	ga := internal.NewGA(conf)

	ga.Run()

	helper.PlotFitness("plot.png", ga.Stat())

}

func oneMaxFitness(genome []float64) float64 {
	var sum float64
	for _, v := range genome {
		sum += v
	}
	return sum
}
