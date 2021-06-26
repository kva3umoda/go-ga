package mutator

import "github.com/kva3umoda/go-ga/genome"

type gaussian struct {
	mu    float64
	sigma float64
	indpd float64
}

func Gaussian(mu, sigma, indpd float64) *gaussian {
	return &gaussian{}
}

func (g gaussian) Mutate(individual *genome.Individual) {
	panic("implement me")
}

// def mutGaussian(individual, mu, sigma, indpb):
//    """This function applies a gaussian mutation of mean *mu* and standard
//    deviation *sigma* on the input individual. This mutation expects a
//    :term:`sequence` individual composed of real valued attributes.
//    The *indpb* argument is the probability of each attribute to be mutated.
//
//    :param individual: Individual to be mutated.
//    :param mu: Mean or :term:`python:sequence` of means for the
//               gaussian addition mutation.
//    :param sigma: Standard deviation or :term:`python:sequence` of
//                  standard deviations for the gaussian addition mutation.
//    :param indpb: Independent probability for each attribute to be mutated.
//    :returns: A tuple of one individual.
//
//    This function uses the :func:`~random.random` and :func:`~random.gauss`
//    functions from the python base :mod:`random` module.
//    """
//    size = len(individual)
//    if not isinstance(mu, Sequence):
//        mu = repeat(mu, size)
//    elif len(mu) < size:
//        raise IndexError("mu must be at least the size of individual: %d < %d" % (len(mu), size))
//    if not isinstance(sigma, Sequence):
//        sigma = repeat(sigma, size)
//    elif len(sigma) < size:
//        raise IndexError("sigma must be at least the size of individual: %d < %d" % (len(sigma), size))
//
//    for i, m, s in zip(xrange(size), mu, sigma):
//        if random.random() < indpb:
//            individual[i] += random.gauss(m, s)
//
//    return individual,
