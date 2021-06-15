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
	POPULATION_SIZE   = 200
	P_CROSSOVER       = 0.9
	P_MUTATION        = 0.2
	MAX_GENERATIONS   = 50
	HALL_OF_FAME_SIZE = 1
)

func main() {
	conf := internal.NewConfig().
		// создание бинарной популяции с размером генома равный коли
		Population(POPULATION_SIZE, population.NewPopulationBinary(len(items))).
		// функция оценки
		CostFunction(1.0, calcValue).
		// алгоритм отбора
		Selector(selector.NewTournament(3)).
		// алгоритм скрещивания
		Crossover(P_CROSSOVER, crossovers.NewOrdered()).
		// алгоритм мутации
		Mutator(P_MUTATION, mutators.NewShuffleIndexes(1.0/float64(len(items)))).
		// количество эпох
		Generation(MAX_GENERATIONS).
		// Добавляем за славы
		HallOfFame(halloffame.NewBase(HALL_OF_FAME_SIZE)).
		// Элитизм
		Elitism(3)

	ga := internal.NewGA(conf)

	ga.Run()

	helper.PlotFitness(helper.GetCurDir()+"/examples/04_combinatorial_optimization/01_knapsack/knapsack.png", ga.Stat())
	best := ga.BestIndividuals()[0]
	printItems(best.Genome)
	fmt.Println(best)
}
