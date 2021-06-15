package main

import (
	"fmt"

	"ga-book/examples/helper"
	"ga-book/internal"
	"ga-book/internal/crossovers"
	"ga-book/internal/halloffame"
	"ga-book/internal/mutators"
	"ga-book/internal/population"
	"ga-book/internal/selector"
)

const (
	TSP_NAME = "bayg29.tsp"

	POPULATION_SIZE   = 300
	MAX_GENERATIONS   = 200
	HALL_OF_FAME_SIZE = 30
	P_CROSSOVER       = 0.9 // probability for crossover
	P_MUTATION        = 0.1 // probability for mutating an individual
)

func main() {
	mapCities := NewMapCities(helper.GetCurDir() + "/examples/04_combinatorial_optimization/02_traveling_salesman/" + TSP_NAME)

	conf := internal.NewConfig().
		// создание бинарной популяции с размером генома равный коли
		Population(POPULATION_SIZE, population.NewPopulationOrdered(len(mapCities.cities))).
		// функция оценки
		CostFunction(-1.0, mapCities.TotalDistance).
		// алгоритм отбора
		Selector(selector.NewTournament(2)).
		// алгоритм скрещивания
		Crossover(P_CROSSOVER, crossovers.NewOrdered()).
		// алгоритм мутации
		Mutator(P_MUTATION, mutators.NewShuffleIndexes(1.0/float64(len(mapCities.cities)))).
		// количество эпох
		Generation(MAX_GENERATIONS).
		// Добавляем за славы
		HallOfFame(halloffame.NewBase(HALL_OF_FAME_SIZE)).
		// Элитизм
		Elitism(3)

	ga := internal.NewGA(conf)

	ga.Run()

	helper.PlotFitness(helper.GetCurDir()+"/examples/04_combinatorial_optimization/02_traveling_salesman/plot.png", ga.Stat())
	best := ga.BestIndividuals()[0]
	//printItems(best.Genome)
	fmt.Println(best)
}
