package friedman

import "github.com/pkg/errors"

// Generate the "Friedman #1" regression problem.
//     This dataset is described in Friedman [1] and Breiman [2].
//
//    Inputs `X` are independent features uniformly distributed on the interval
//    [0, 1]. The output `y` is created according to the formula::
//
//        y(X) = 10 * sin(pi * X[:, 0] * X[:, 1]) + 20 * (X[:, 2] - 0.5) ** 2 + 10 * X[:, 3] + 5 * X[:, 4] + noise * N(0, 1).

//	n_samples : int, default=100
//       The number of samples.
//
//   n_features : int, default=10
//       The number of features. Should be at least 5.
//
//   noise : float, default=0.0
//       The standard deviation of the gaussian noise applied to the output.

func MakeFriedman1(nSamples, nFeatures int, noise float64) ([][]float64, error) {
	if nFeatures < 5 {
		return nil, errors.New("n_features must be at least five.")
	}

	return nil, nil
}

/**
def make_friedman1(n_samples=100, n_features=10, *, noise=0.0, random_state=None):
    """Generate the "Friedman #1" regression problem.

    This dataset is described in Friedman [1] and Breiman [2].

    Inputs `X` are independent features uniformly distributed on the interval
    [0, 1]. The output `y` is created according to the formula::

        y(X) = 10 * sin(pi * X[:, 0] * X[:, 1]) + 20 * (X[:, 2] - 0.5) ** 2 \
+ 10 * X[:, 3] + 5 * X[:, 4] + noise * N(0, 1).

    Out of the `n_features` features, only 5 are actually used to compute
    `y`. The remaining features are independent of `y`.

    The number of features has to be >= 5.

    Read more in the :ref:`User Guide <sample_generators>`.

    Parameters
    ----------
    n_samples : int, default=100
        The number of samples.

    n_features : int, default=10
        The number of features. Should be at least 5.

    noise : float, default=0.0
        The standard deviation of the gaussian noise applied to the output.

    random_state : int, RandomState instance or None, default=None
        Determines random number generation for dataset noise. Pass an int
        for reproducible output across multiple function calls.
        See :term:`Glossary <random_state>`.

    Returns
    -------
    X : ndarray of shape (n_samples, n_features)
        The input samples.

    y : ndarray of shape (n_samples,)
        The output values.

    References
    ----------
    .. [1] J. Friedman, "Multivariate adaptive regression splines", The Annals
           of Statistics 19 (1), pages 1-67, 1991.

    .. [2] L. Breiman, "Bagging predictors", Machine Learning 24,
           pages 123-140, 1996.
    """
    if n_features < 5:
        raise ValueError("n_features must be at least five.")

    generator = check_random_state(random_state)

    X = generator.rand(n_samples, n_features)
    y = (
        10 * np.sin(np.pi * X[:, 0] * X[:, 1])
        + 20 * (X[:, 2] - 0.5) ** 2
        + 10 * X[:, 3]
        + 5 * X[:, 4]
        + noise * generator.randn(n_samples)
    )

    return X, y
*/