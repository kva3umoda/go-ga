package mutators

type UniformInt struct {

}


func NewUniformInt() *UniformInt {
	return &UniformInt{}
}

func (u *UniformInt) Mutate(individual []float64) {
	panic("implement me")
}


// def mutUniformInt(individual, low, up, indpb):
//    """Mutate an individual by replacing attributes, with probability *indpb*,
//    by a integer uniformly drawn between *low* and *up* inclusively.
//
//    :param individual: :term:`Sequence <sequence>` individual to be mutated.
//    :param low: The lower bound or a :term:`python:sequence` of
//                of lower bounds of the range from which to draw the new
//                integer.
//    :param up: The upper bound or a :term:`python:sequence` of
//               of upper bounds of the range from which to draw the new
//               integer.
//    :param indpb: Independent probability for each attribute to be mutated.
//    :returns: A tuple of one individual.
//    """
//    size = len(individual)
//    if not isinstance(low, Sequence):
//        low = repeat(low, size)
//    elif len(low) < size:
//        raise IndexError("low must be at least the size of individual: %d < %d" % (len(low), size))
//    if not isinstance(up, Sequence):
//        up = repeat(up, size)
//    elif len(up) < size:
//        raise IndexError("up must be at least the size of individual: %d < %d" % (len(up), size))
//
//    for i, xl, xu in zip(xrange(size), low, up):
//        if random.random() < indpb:
//            individual[i] = random.randint(xl, xu)
//
//    return individual,
