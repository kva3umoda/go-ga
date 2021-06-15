package population

type CreateFunc func(populationSize int) *Population

type Creator interface {
	Create(populationSize int) *Population
}


type Function struct {
	create CreateFunc
}

func NewFunction(createFunc CreateFunc) *Function {
	return &Function{
		create: createFunc,
	}
}

func (f *Function) Create(populationSize int) *Population {
	return f.create(populationSize)
}
