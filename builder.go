package go_ga

import (
	"github.com/kva3umoda/go-ga/crossover"
	"github.com/kva3umoda/go-ga/fitness"
	"github.com/kva3umoda/go-ga/mutator"
	"github.com/kva3umoda/go-ga/population"
	"github.com/kva3umoda/go-ga/selector"
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
}

func NewBuilder() *Builder {
	return &Builder{}
}

// Population - указание размера популяции
func (b *Builder) Population(size int) *Builder {
	b.populationSize = size
	return b
}

// HallOfFame - установка размер зала славы
func (b *Builder) HallOfFame(size int) *Builder {
	b.hallOfFameSize = size
	return b
}

// Elitism - размер элитизма
func (b *Builder) Elitism(size int) *Builder {
	b.elitismSize = size
	return b
}

// Creator - создатель популяции
func (b *Builder) Creator(creator population.Creator) *Builder {
	b.creator = creator
	return b
}

// CreatorFunc - создатель популции
func (b *Builder) CreatorFunc(createFunc population.CreateFunc) *Builder {
	b.creator = population.NewFunction(createFunc)
	return b
}

// MutatorProb - Вероятность мутации
func (b *Builder) MutatorProb(prob float64) *Builder {
	b.mutatorProb = prob
	return b
}

// Mutator - алгоритм мутации
func (b *Builder) Mutator(mutator mutator.Mutator) *Builder {
	b.mutator = mutator
	return b
}

// MutatorFunc - алгоритм мутации
func (b *Builder) MutatorFunc(mutatorFunc mutator.MutatorFunc) *Builder {
	b.mutator = mutator.NewFunction(mutatorFunc)
	return b
}

// CrossoverProb - вероятность скрещивания
func (b *Builder) CrossoverProb(prob float64) *Builder {
	b.crossoverProb = prob
	return b
}

// Crossover - алгоритм скрещивания
func (b *Builder) Crossover(crossover crossover.Crossover) *Builder {
	b.crossover = crossover
	return b
}

// CrossoverFunc - алгоритм скрещивания
func (b *Builder) CrossoverFunc(crossoverFunc crossover.CrossoverFunc) *Builder {
	b.crossover = crossover.NewFunction(crossoverFunc)
	return b
}

// Generation - установка максимальное числа поколений
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

// Создание генетического алгоритма
func (b *Builder) Build() (*GA, error) {

	return nil, nil
}
