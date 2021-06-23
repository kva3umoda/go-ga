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
	// problem constants:

	// Genetic Algorithm constants:
	POPULATION_SIZE   = 300
	P_CROSSOVER       = 0.9 // probability for crossover
	P_MUTATION        = 0.1 // probability for mutating an individual
	MAX_GENERATIONS   = 200
	HALL_OF_FAME_SIZE = 30
)

func main() {
	ln := getNumberShiftsShedule()
	builder := go_ga.NewBuilder().
		// создание бинарной популяции с размером генома равный коли
		Population(POPULATION_SIZE).
		Creator(population.BinaryPopulation(ln)).
		// функция оценки
		CostFunction(getCost).
		// необходимо минимальное значение
		Fitness(fitness.Min()).
		// алгоритм отбора
		Selector(selector.Tournament(2)).
		// алгоритм скрещивания
		CrossoverProb(P_CROSSOVER).
		Crossover(crossover.TwoPoint()).
		// алгоритм мутации
		MutatorProb(P_MUTATION).
		Mutator(mutator.FlitBit(1.0 / float64(ln))).
		// количество эпох
		Generation(MAX_GENERATIONS).
		// Добавляем за славы
		HallOfFame(HALL_OF_FAME_SIZE).
		// Элитизм
		Elitism(HALL_OF_FAME_SIZE)

	ga, err := builder.Build()
	if err != nil {
		log.Fatal(err)
	}

	ga.Run()

	helper.PlotFitness("examples/05_restricted_tasks/02_sheduler/plot.png", ga.Stat())

	bests := ga.BestIndividuals()

	fmt.Printf("-- Best Individual = %s\n", bests[0].String())
	fmt.Print("-- Schedule = \n")

	printScheduleInfo(bests[0].Genome)
}
