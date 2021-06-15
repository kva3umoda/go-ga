package fitness

type FitnessFunc func(cost float64) float64

type Fitness interface {
	Fitness(cost float64) float64
}

type Function struct {
	fitnessFunc FitnessFunc
}

func NewFunction(fitnessFunc FitnessFunc) *Function {
	return &Function{
		fitnessFunc: fitnessFunc,
	}
}

func (f *Function) Fitness(cost float64) float64 {
	return f.fitnessFunc(cost)
}

type Max struct {
}

func NewMax() *Max {
	return &Max{}
}

func (m *Max) Fitness(cost float64) float64 {
	return cost
}

type Min struct {
}

func NewMin() *Min {
	return &Min{}
}

func (m *Min) Fitness(cost float64) float64 {
	return -1 * cost
}
