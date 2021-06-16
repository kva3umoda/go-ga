package main

import (
	"fmt"
	"log"

	go_ga "github.com/kva3umoda/go-ga"
	"github.com/kva3umoda/go-ga/crossover"
	"github.com/kva3umoda/go-ga/examples/helper"
	"github.com/kva3umoda/go-ga/mutator"
	"github.com/kva3umoda/go-ga/population"
	"github.com/kva3umoda/go-ga/selector"
)

const (
	POPULATION_SIZE   = 200
	P_CROSSOVER       = 0.9
	P_MUTATION        = 0.2
	MAX_GENERATIONS   = 50
	HALL_OF_FAME_SIZE = 1
)

func main() {
	builder := go_ga.NewBuilder().
		// создание бинарной популяции с размером генома равный коли
		Population(POPULATION_SIZE).
		Creator(population.NewBinaryPopulation(len(items))).
		// функция оценки
		CostFunction(calcValue).
		// алгоритм отбора
		Selector(selector.NewTournament(3)).
		// алгоритм скрещивания
		CrossoverProb(P_CROSSOVER).
		Crossover(crossover.NewOrdered()).
		// алгоритм мутации
		MutatorProb(P_MUTATION).
		Mutator(mutator.NewShuffleIndexes(1.0 / float64(len(items)))).
		// количество эпох
		Generation(MAX_GENERATIONS).
		// Добавляем за славы
		HallOfFame(HALL_OF_FAME_SIZE).
		// Элитизм
		Elitism(3)

	ga, err := builder.Build()
	if err != nil {
		log.Fatal(err)
	}

	ga.Run()

	helper.PlotFitness("examples/04_combinatorial_optimization/01_knapsack/knapsack.png", ga.Stat())
	best := ga.BestIndividuals()[0]
	printItems(best.Genome)
	fmt.Println(best)
}
