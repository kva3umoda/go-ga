package genome


type PopulationCreator interface {
	Create(populationSize int) *Population
}
