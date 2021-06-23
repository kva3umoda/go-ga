package fitness

type FitnessFunc func(cost float64) float64

type Fitness interface {
	Fitness(cost float64) float64
}

type function struct {
	fitnessFunc FitnessFunc
}

func Function(fitnessFunc FitnessFunc) Fitness {
	return &function{
		fitnessFunc: fitnessFunc,
	}
}

func (f *function) Fitness(cost float64) float64 {
	return f.fitnessFunc(cost)
}

// max fitness
type max struct{}

func Max() Fitness {
	return &max{}
}

func (m *max) Fitness(cost float64) float64 {
	return cost
}

// min fitness
type min struct{}

func Min() Fitness {
	return &min{}
}

func (m *min) Fitness(cost float64) float64 {
	return -1 * cost
}
