package main

import (
	"log"


	"github.com/kva3umoda/go-ga"
	"github.com/kva3umoda/go-ga/crossover"
	"github.com/kva3umoda/go-ga/examples/helper"
	"github.com/kva3umoda/go-ga/mutator"
	"github.com/kva3umoda/go-ga/population"
	"github.com/kva3umoda/go-ga/selector"
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
	builder := go_ga.NewBuilder().
		Population(POPULATION_SIZE).
		Creator(population.NewBinaryPopulation(ONE_MAX_LENGTH)).
		CostFunction(oneMaxFitness).
		Selector(selector.NewTournament(3)).
		CrossoverProb(P_CROSSOVER).
		Crossover(crossover.NewOnePoint()).
		MutatorProb(P_MUTATION).
		Mutator(mutator.NewFlitBit(1.0/ONE_MAX_LENGTH)).
		Generation(MAX_GENERATIONS)

	ga, err := builder.Build()
	if err != nil {
		log.Fatal(err)
	}

	ga.Run()

	helper.PlotFitness("examples/03_onemax/plot.png", ga.Stat())

}

func oneMaxFitness(genome []float64) float64 {
	var sum float64
	for _, v := range genome {
		sum += v
	}
	return sum
}
