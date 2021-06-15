package selector


// def selEpsilonLexicase(individuals, k, epsilon):
//    """
//    Returns an individual that does the best on the fitness cases when
//    considered one at a time in random order. Requires a epsilon parameter.
//    https://push-language.hampshire.edu/uploads/default/original/1X/35c30e47ef6323a0a949402914453f277fb1b5b0.pdf
//    Implemented epsilon_y implementation.
//
//    :param individuals: A list of individuals to select from.
//    :param k: The number of individuals to select.
//    :returns: A list of selected individuals.
//    """
//    selected_individuals = []
//
//    for i in range(k):
//        fit_weights = individuals[0].fitness.weights
//
//        candidates = individuals
//        cases = list(range(len(individuals[0].fitness.values)))
//        random.shuffle(cases)
//
//        while len(cases) > 0 and len(candidates) > 1:
//            if fit_weights[cases[0]] > 0:
//                best_val_for_case = max(map(lambda x: x.fitness.values[cases[0]], candidates))
//                min_val_to_survive_case = best_val_for_case - epsilon
//                candidates = list(filter(lambda x: x.fitness.values[cases[0]] >= min_val_to_survive_case, candidates))
//            else:
//                best_val_for_case = min(map(lambda x: x.fitness.values[cases[0]], candidates))
//                max_val_to_survive_case = best_val_for_case + epsilon
//                candidates = list(filter(lambda x: x.fitness.values[cases[0]] <= max_val_to_survive_case, candidates))
//
//            cases.pop(0)
//
//        selected_individuals.append(random.choice(candidates))
//
//    return selected_individuals
