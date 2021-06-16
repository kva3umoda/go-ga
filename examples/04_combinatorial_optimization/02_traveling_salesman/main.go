package main

import (
	"fmt"
	"log"

	go_ga "github.com/kva3umoda/go-ga"
	"github.com/kva3umoda/go-ga/crossover"
	"github.com/kva3umoda/go-ga/examples/helper"
	"github.com/kva3umoda/go-ga/fitness"
	"github.com/kva3umoda/go-ga/mutator"
	"github.com/kva3umoda/go-ga/population"
	"github.com/kva3umoda/go-ga/selector"
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
	mapCities := NewMapCities("examples/04_combinatorial_optimization/02_traveling_salesman/" + TSP_NAME)

	builder := go_ga.NewBuilder().
		// создание бинарной популяции с размером генома равный коли
		Population(POPULATION_SIZE).
		Creator(population.NewOrderedPopulation(len(mapCities.cities))).
		// функция оценки
		CostFunction(mapCities.TotalDistance).
		Fitness(fitness.NewMin()).
		// алгоритм отбора
		Selector(selector.NewTournament(2)).
		// алгоритм скрещивания
		CrossoverProb(P_CROSSOVER).
		Crossover(crossover.NewOrdered()).
		// алгоритм мутации
		MutatorProb(P_MUTATION).
		Mutator(mutator.NewShuffleIndexes(1.0 / float64(len(mapCities.cities)))).
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

	helper.PlotFitness("examples/04_combinatorial_optimization/02_traveling_salesman/plot.png", ga.Stat())
	best := ga.BestIndividuals()[0]
	//printItems(best.Genome)
	fmt.Println(best)
}
