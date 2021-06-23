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
	NUM_OF_QUEENS = 16

	// Genetic Algorithm constants:
	POPULATION_SIZE   = 300
	MAX_GENERATIONS   = 100
	HALL_OF_FAME_SIZE = 30
	P_CROSSOVER       = 0.9 // probability for crossover
	P_MUTATION        = 0.1 // probability for mutating an individual
)

func main() {
	builder := go_ga.NewBuilder().
		// создание бинарной популяции с размером генома равный коли
		Population(POPULATION_SIZE).
		Creator(population.OrderedPopulation(NUM_OF_QUEENS)).
		// функция оценки
		CostFunction(violationsCount).
		// необходимо минимальное значение
		Fitness(fitness.Min()).
		// алгоритм отбора
		Selector(selector.Tournament(2)).
		// алгоритм скрещивания
		CrossoverProb(P_CROSSOVER).
		Crossover(crossover.UniformPartialMatched(2.0 / float64(NUM_OF_QUEENS))).
		// алгоритм мутации
		MutatorProb(P_MUTATION).
		Mutator(mutator.ShuffleIndexes(1.0 / float64(NUM_OF_QUEENS))).
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

	helper.PlotFitness("examples/05_restricted_tasks/01_n_queens/plot.png", ga.Stat())

	bests := ga.BestIndividuals()
	for i, best := range bests {
		fmt.Printf("#%d, %s\n", i, best)
	}
}

// Возвращает количество нарушений в решение
func violationsCount(positions []float64) float64 {
	violations := 0.0

	// iterate over every pair of queens and find if they are on the same diagonal:
	for i := range positions {
		for j := i + 1; j < len(positions); j++ {
			// first queen in pair:
			column1 := i
			row1 := int(positions[i])

			// second queen in pair:
			column2 := j
			row2 := int(positions[j])

			//# look for diagonal threat for th ecurrent pair:
			if abs(column1-column2) == abs(row1-row2) {
				violations += 1
			}
		}
	}
	return violations
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
