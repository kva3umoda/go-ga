package main

import (
	"fmt"
	"log"
	"math"

	go_ga "github.com/kva3umoda/go-ga"
	"github.com/kva3umoda/go-ga/crossover"
	"github.com/kva3umoda/go-ga/examples/helper"
	"github.com/kva3umoda/go-ga/fitness"
	"github.com/kva3umoda/go-ga/genome"
	math_helper "github.com/kva3umoda/go-ga/helper"
	"github.com/kva3umoda/go-ga/mutator"
	"github.com/kva3umoda/go-ga/population"
	"github.com/kva3umoda/go-ga/selector"
)

var (
	BOUND_LOW = []float64{-5, -5} // boundaries for all dimensions
	BOUND_UP  = []float64{5, 5}
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

	DISTANCE_THRESHOLD = 0.1
	SHARING_EXTENT     = 5.0
)

func main() {
	builder := go_ga.NewBuilder().
		// создание бинарной популяции с размером генома равный коли
		Population(POPULATION_SIZE).
		Creator(population.FloatPopulation(DIMENSIONS, BOUND_LOW, BOUND_UP)).
		// функция оценки
		CostFunction(himmelblauInverted).
		// необходимо минимальное значение
		Fitness(fitness.Max()).
		// алгоритм отбора
		Selector(TournamentWithSharing(2, SHARING_EXTENT, DISTANCE_THRESHOLD)).
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

	helper.PlotFitness("examples/06_optimize_functions/02_optimize_himmelblau/plot.png", ga.Stat())

	bests := ga.BestIndividuals()
	for i, best := range bests {
		fmt.Printf("#%d, %s\n", i, best)
	}

	globalMinima := [][]float64{{3.0, 2.0}, {-2.805118, 3.131312}, {-3.779310, -3.283186}, {3.584458, -1.848126}}
	for _, minima := range globalMinima {
		fmt.Printf("Best %+v: %v\n", minima, himmelblau(minima))
	}

}

func himmelblau(genome []float64) float64 {
	x := genome[0]
	y := genome[1]
	f := math.Pow((math.Pow(x, 2)+y-11), 2) + math.Pow((x+math.Pow(y, 2)-7), 2)
	return f
}

// Значения функции Химмельблау находятся в диапазоне от 0 до (примерно) 2000
func himmelblauInverted(genome []float64) float64 {
	return 2000 - himmelblau(genome)
}

// реализация механизм разделения внутри оператора разделения, так как в нем применятся приспособленность всех индивидумов
type tournamentWithSharing struct {
	tournament        selector.Selector
	sharingExtent     float64 //
	distanceThreshold float64 // растояние порога
}

func TournamentWithSharing(tournSize int, sharingExtent, distanceThreshold float64) *tournamentWithSharing {
	return &tournamentWithSharing{
		tournament:        selector.Tournament(tournSize),
		sharingExtent:     sharingExtent,
		distanceThreshold: distanceThreshold,
	}
}

func (t *tournamentWithSharing) Select(populationSize int, individuals []*genome.Individual) []*genome.Individual {
	origFitnesses := make([]float64, len(individuals))
	for i, ind := range individuals {
		origFitnesses[i] = ind.Fitness
	}

	for i := range individuals {
		sharingSum := 1.0
		// iterate over all other individuals
		for j := range individuals {
			if i == j {
				continue
			}
			//calculate eucledean distance between individuals:
			distance := math_helper.Distance(individuals[i].Genome, individuals[j].Genome)
			if distance < t.distanceThreshold {
				sharingSum += (1 - distance/(t.sharingExtent*t.distanceThreshold))
			}

		}
		individuals[i].Fitness = origFitnesses[i] / sharingSum
	}
	selected := t.tournament.Select(populationSize, individuals)

	return selected
}

/**
def selTournamentWithSharing(individuals, k, tournsize, fit_attr="fitness"):

    # get orig fitnesses:
    origFitnesses = [ind.fitness.values[0] for ind in individuals]

    # apply sharing to each individual:
    for i in range(len(individuals)):
        sharingSum = 1

        # iterate over all other individuals
        for j in range(len(individuals)):
            if i != j:
                # calculate eucledean distance between individuals:
                distance = math.sqrt(
                    ((individuals[i][0] - individuals[j][0]) ** 2) + ((individuals[i][1] - individuals[j][1]) ** 2))

                if distance < DISTANCE_THRESHOLD:
                    sharingSum += (1 - distance / (SHARING_EXTENT * DISTANCE_THRESHOLD))

        # reduce fitness accordingly:
        individuals[i].fitness.values = origFitnesses[i] / sharingSum,

    # apply original tools.selTournament() using modified fitness:
    selected = tools.selTournament(individuals, k, tournsize, fit_attr)

    # retrieve original fitness:
    for i, ind in enumerate(individuals):
        ind.fitness.values = origFitnesses[i],

    return selected

*/
