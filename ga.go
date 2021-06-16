package go_ga

import (
	"github.com/kva3umoda/go-ga/crossover"
	"github.com/kva3umoda/go-ga/fitness"
	"github.com/kva3umoda/go-ga/genome"
	"github.com/kva3umoda/go-ga/halloffame"
	"github.com/kva3umoda/go-ga/mutator"
	"github.com/kva3umoda/go-ga/population"
	"github.com/kva3umoda/go-ga/rand"
	"github.com/kva3umoda/go-ga/selector"
	"github.com/kva3umoda/go-ga/stat"
)

/**
		Начало
      	|
		Создать начальную
		популяцию (поколение 0)
      	|
		Вычеслить приспособленность
		каждого индивидуума в популяции
      	|
 |--->	Отбор
 |     	|
 |		Скрещивание
 |     	|
 |		Мутация
 |     	|
 |		Вычеслить пригодность
 |  	каждого индивидуума
 |		в популяции
 |     	|
 <--Условия остановки
	выполнения
      	|
		Выбрать индивидуума с максимально приспособленностью
      	|
		Конец
*/

type GA struct {
	stat           *stat.Fitness
	generation     int
	populationSize int
	elitismSize    int // размер элитизма
	maxGeneration  int

	selector selector.Selector

	mutatorProb float64         // вероятность мутации
	mutator     mutator.Mutator // алгоритм мутации

	crossoverProb float64             // вероятность скрещивания
	crossover     crossover.Crossover // алгоритм скрещивания

	costFunc fitness.CostFunc // функция оценки
	fitness  fitness.Fitness
	creator  population.Creator // генератор популяции

	hallOfFame halloffame.HallOfFame

	population *population.Population
}

func (ga *GA) Mutate(pop *population.Population) {
	if ga.mutator == nil {
		return
	}
	for _, ind := range pop.Individuals {
		if rand.Float() >= ga.mutatorProb {
			continue
		}
		ga.mutator.Mutate(ind)
	}
}

func (ga *GA) Select(pop *population.Population) *population.Population {
	npop := new(population.Population)
	npop.Individuals = ga.selector.Select(ga.populationSize, pop.Individuals)
	npop.Individuals = npop.Individuals[:ga.populationSize-ga.elitismSize]

	return npop
}

func (ga *GA) Generation() int {
	return ga.generation
}

// возвращает популяцию
func (ga *GA) Population() *population.Population {
	return ga.population
}

func (ga *GA) Crossing(pop *population.Population) {
	size := len(pop.Individuals)
	for i := 1; i < size; i += 2 {
		if rand.Float() < ga.crossoverProb {
			ga.crossover.Crossing(pop.Individuals[i-1], pop.Individuals[i])
			pop.Individuals[i-1].Fitness = 0
			pop.Individuals[i-1].Cost = 0
			pop.Individuals[i].Fitness = 0
			pop.Individuals[i].Cost = 0
		}
	}
}

func (ga *GA) Fitness(pop *population.Population) {
	for _, ind := range pop.Individuals {
		ind.Cost = ga.costFunc(ind.Genome)
		ind.Fitness = ga.fitness.Fitness(ind.Cost)
		ga.stat.Add(ga.generation, ind.Cost)
	}
}

func (ga *GA) HallOfFame(pop *population.Population) {
	if ga.hallOfFame == nil {
		return
	}
	inds := ga.hallOfFame.Individuals()
	for i := 0; i < len(inds); i++ {
		if i >= ga.elitismSize {
			break
		}
		pop.Individuals = append(pop.Individuals, inds[i])
		ga.stat.Add(ga.generation, inds[i].Cost)
	}
	ga.hallOfFame.Update(pop)

}

// создает популяцию
func (ga *GA) Create() *population.Population {
	return ga.creator.Create(ga.populationSize)
}

func (ga *GA) Begin() {
	ga.generation = 0
	ga.population = ga.Create()

	ga.Fitness(ga.population)

	if ga.hallOfFame == nil {
		return
	}
	ga.hallOfFame.Update(ga.population)
}

func (ga *GA) Evolute() {
	ga.generation++
	// Отбор
	offspring := ga.Select(ga.population)
	offspring.Generation = ga.generation
	// скрещивание
	ga.Crossing(offspring)
	// мутация
	ga.Mutate(offspring)
	// расчет приспособленности
	ga.Fitness(offspring)
	// зал славы и
	ga.HallOfFame(offspring)

	ga.population = offspring
}

func (ga *GA) Run() {
	ga.Begin()
	for i := 0; i < ga.maxGeneration; i++ {
		ga.Evolute()
	}
}

func (ga *GA) BestIndividuals() []*genome.Individual {
	return ga.hallOfFame.Individuals()
}

func (ga *GA) Stat() *stat.Fitness {
	return ga.stat
}
