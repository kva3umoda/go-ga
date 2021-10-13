package main

import (
	"fmt"
	"log"
	"math"

	go_ga "github.com/kva3umoda/go-ga"
	"github.com/kva3umoda/go-ga/crossover"
	"github.com/kva3umoda/go-ga/examples/helper"
	"github.com/kva3umoda/go-ga/fitness"
	"github.com/kva3umoda/go-ga/mutator"
	"github.com/kva3umoda/go-ga/population"
	"github.com/kva3umoda/go-ga/selector"
)

var (
	BOUND_LOW = []float64{-1.25, -1.25} // boundaries for all dimensions
	BOUND_UP  = []float64{1.25, 1.25}
)

const (
	DIMENSIONS = 2 // number of dimensions
	// Genetic Algorithm constants:
	POPULATION_SIZE   = 300
	P_CROSSOVER       = 0.9 // probability for crossover
	P_MUTATION        = 0.5 // (try also 0.5 ) probability for mutating an individual
	MAX_GENERATIONS   = 300
	HALL_OF_FAME_SIZE = 30
	CROWDING_FACTOR   = 20.0 // crowding factor for crossover and mutation
	PENALTY_VALUE     = 10.0 // фиксированный штраф за на рушение ограничения
)

func main() {
	builder := go_ga.NewBuilder().
		// создание бинарной популяции с размером генома равный коли
		Population(POPULATION_SIZE).
		Creator(population.FloatPopulation(DIMENSIONS, BOUND_LOW, BOUND_UP)).
		// функция оценки
		CostFunction(simionescu).
		// необходимо минимальное значение
		Fitness(fitness.DeltaPenalty(feasible, PENALTY_VALUE, 1.0)).
		// алгоритм отбора
		Selector(selector.Tournament(2)).
		// алгоритм скрещивания
		CrossoverProb(P_CROSSOVER).
		Crossover(crossover.SimulatedBinaryBounded(CROWDING_FACTOR, BOUND_LOW, BOUND_UP)).
		// алгоритм мутации
		MutatorProb(P_MUTATION).
		Mutator(mutator.PolynomialBounded(CROWDING_FACTOR, BOUND_LOW, BOUND_UP, 1.0/DIMENSIONS)).
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

	helper.PlotFitness("examples/06_optimize_functions/01_optimize_eggholder/plot.png", ga.Stat())

	bests := ga.BestIndividuals()
	for i, best := range bests[0:2] {
		fmt.Printf("#%d, %s\n", i, best)
	}

	/*
		x = 0.84852813, y = –0.84852813
		x = –0.84852813, y = 0.84852813
	*/

	bestPoints := [][]float64{{0.84852813, -0.84852813}, {-0.84852813, 0.84852813}}
	for _, best := range bestPoints {
		fmt.Printf("Best %+v: %v\n", best, simionescu(best))
	}

}

// функцией Симионеску
func simionescu(genome []float64) float64 {
	x := genome[0]
	y := genome[1]
	f := 0.1 * x * y
	return f
}

// функция допустимой области
func feasible(genome []float64) bool {
	x := genome[0]
	y := genome[1]
	return math.Pow(x, 2)+math.Pow(y, 2) <= math.Pow(1.0+0.2*math.Cos(8.0*math.Atan2(x, y)), 2)
}
