package go_ga

import (
	"errors"

	"github.com/kva3umoda/go-ga/crossover"
	"github.com/kva3umoda/go-ga/fitness"
	"github.com/kva3umoda/go-ga/halloffame"
	"github.com/kva3umoda/go-ga/mutator"
	"github.com/kva3umoda/go-ga/population"
	"github.com/kva3umoda/go-ga/selector"
	"github.com/kva3umoda/go-ga/stat"
)

type Builder struct {
	populationSize int // размер популяции
	hallOfFameSize int // размер зала славы
	elitismSize    int // размер элитизма

	maxGeneration int // максимальное число поколений

	creator population.Creator // генератор популяции

	mutatorProb float64         // вероятность мутации
	mutator     mutator.Mutator // алгоритм мутации

	selector selector.Selector // алгорим отбора популяции

	crossoverProb float64             // вероятность скрещивания
	crossover     crossover.Crossover // алгоритм скрещивания

	costFunc fitness.CostFunc // функция оценки
	fitness  fitness.Fitness

	hallOfFame halloffame.HallOfFame
	elitism    halloffame.HallOfFame
}

func NewBuilder() *Builder {
	return &Builder{}
}

// Population - указание размера популяции.
func (b *Builder) Population(size int) *Builder {
	b.populationSize = size
	return b
}

// HallOfFame - установка размер зала славы. минимальное значение 1
func (b *Builder) HallOfFame(size int) *Builder {
	b.hallOfFameSize = size
	return b
}

// Elitism - размер элитизма.
func (b *Builder) Elitism(size int) *Builder {
	b.elitismSize = size
	return b
}

// Creator - создатель популяции.
func (b *Builder) Creator(creator population.Creator) *Builder {
	b.creator = creator
	return b
}

// CreatorFunc - создатель популции.
func (b *Builder) CreatorFunc(createFunc population.CreateFunc) *Builder {
	b.creator = population.NewFunction(createFunc)
	return b
}

// MutatorProb - Вероятность мутации.
func (b *Builder) MutatorProb(prob float64) *Builder {
	b.mutatorProb = prob
	return b
}

// Mutator - алгоритм мутации.
func (b *Builder) Mutator(mutator mutator.Mutator) *Builder {
	b.mutator = mutator
	return b
}

// MutatorFunc - алгоритм мутации.
func (b *Builder) MutatorFunc(mutatorFunc mutator.MutatorFunc) *Builder {
	b.mutator = mutator.NewFunction(mutatorFunc)
	return b
}

// CrossoverProb - вероятность скрещивания.
func (b *Builder) CrossoverProb(prob float64) *Builder {
	b.crossoverProb = prob
	return b
}

// Crossover - алгоритм скрещивания.
func (b *Builder) Crossover(crossover crossover.Crossover) *Builder {
	b.crossover = crossover
	return b
}

// CrossoverFunc - алгоритм скрещивания.
func (b *Builder) CrossoverFunc(crossoverFunc crossover.CrossoverFunc) *Builder {
	b.crossover = crossover.NewFunction(crossoverFunc)
	return b
}

func (b *Builder) Selector(selector selector.Selector) *Builder {
	b.selector = selector
	return b
}

func (b *Builder) SelectorFunc(selectorFunc selector.SelectFunc) *Builder {
	b.selector = selector.NewFunction(selectorFunc)
	return b
}

// Generation - установка максимальное числа поколений.
func (b *Builder) Generation(maxGeneration int) *Builder {
	b.maxGeneration = maxGeneration
	return b
}

func (b *Builder) CostFunction(costFunc fitness.CostFunc) *Builder {
	b.costFunc = costFunc
	return b
}

func (b *Builder) Fitness(fitness fitness.Fitness) *Builder {
	b.fitness = fitness
	return b
}

// Build - build GA.
func (b *Builder) Build() (*GA, error) {
	ga := new(GA)

	// Population
	if b.populationSize <= 0 {
		return nil, errors.New("unsupported 'Population' parameter value")
	}
	ga.populationSize = b.populationSize

	// Elitism
	if b.elitismSize < 0 {
		return nil, errors.New("unsupported 'Elitism' parameter value")
	}
	ga.elitismSize = b.elitismSize

	// HallOfFame
	switch {
	case b.hallOfFameSize < b.elitismSize:
		b.hallOfFameSize = b.elitismSize
	case b.hallOfFameSize == 0:
		b.hallOfFameSize = 1
	case b.hallOfFameSize < 0:
		return nil, errors.New("unsupported 'HallOfFame' parameter value")
	}
	ga.hallOfFame = halloffame.NewBase(b.hallOfFameSize)

	// Creator
	if b.creator == nil {
		return nil, errors.New("unsupported 'Creator' parameter value")
	}
	ga.creator = b.creator
	// Mutator
	ga.mutatorProb = b.mutatorProb
	ga.mutator = b.mutator

	// Crossover
	if b.crossover == nil {
		return nil, errors.New("unsupported 'Crossover' parameter value")
	}
	ga.crossover = b.crossover

	// CrossoverProb
	if b.crossoverProb <= 0.0 {
		return nil, errors.New("unsupported 'CrossoverProb' parameter value")
	}
	ga.crossoverProb = b.crossoverProb

	// Selector
	if b.selector == nil {
		return nil, errors.New("unsupported 'Selector' parameter value")
	}
	ga.selector = b.selector

	if b.maxGeneration <= 0 {
		return nil, errors.New("unsupported 'Generation' parameter value")
	}
	ga.maxGeneration = b.maxGeneration

	if b.costFunc == nil {
		return nil, errors.New("unsupported 'CostFunction' parameter value")
	}
	ga.costFunc = b.costFunc

	ga.fitness = b.fitness
	if ga.fitness == nil {
		ga.fitness = fitness.Max()
	}

	ga.stat = stat.NewFitness(ga.maxGeneration + 1)

	return ga, nil

}
