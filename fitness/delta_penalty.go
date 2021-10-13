package fitness

import "github.com/kva3umoda/go-ga/genome"

type FeasibleFunc = func (ind []float64) bool

//     """This decorator returns penalized fitness for invalid individuals and the
//    original fitness value for valid individuals. The penalized fitness is made
//    of a constant factor *delta* added with an (optional) *distance* penalty. The
//    distance function, if provided, shall return a value growing as the
//    individual moves away the valid zone.
//
//    :param feasibility: A function returning the validity status of any
//                        individual.
//    :param delta: Constant or array of constants returned for an invalid individual.
//    :param distance: A function returning the distance between the individual
//                     and a given valid point. The distance function can also return a sequence
//                     of length equal to the number of objectives to affect multi-objective
//                     fitnesses differently (optional, defaults to 0).
//    :returns: A decorator for evaluation function.
type deltaPenalty struct {
	feasible FeasibleFunc
	penalty  float64
	weight   float64
}

func DeltaPenalty(feasible FeasibleFunc, penalty, weight float64) Fitness {
	dp := &deltaPenalty{
		feasible: feasible,
		penalty:  penalty,
		weight:   weight,
	}
	return dp
}

func (f *deltaPenalty) Fitness(ind *genome.Individual) {
	if f.feasible(ind.Genome) {
		ind.Fitness = f.weight * ind.Cost
		return
	}
	ind.Fitness = f.penalty - f.weight*ind.Cost
}
