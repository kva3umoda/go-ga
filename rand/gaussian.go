package rand

import (
	"math/rand"
)
//https://www.alanzucconi.com/2015/09/16/how-to-sample-from-a-gaussian-distribution/
// Gaussian distribution
func Gauss(	stdDev , mean   float64) float64 {
	return rand.NormFloat64()*stdDev + mean
}


//     def gauss(self, mu, sigma):
//        """Gaussian distribution.
//
//        mu is the mean, and sigma is the standard deviation.  This is
//        slightly faster than the normalvariate() function.
//
//        Not thread-safe without a lock around calls.
//
//        """
//        # When x and y are two variables from [0, 1), uniformly
//        # distributed, then
//        #
//        #    cos(2*pi*x)*sqrt(-2*log(1-y))
//        #    sin(2*pi*x)*sqrt(-2*log(1-y))
//        #
//        # are two *independent* variables with normal distribution
//        # (mu = 0, sigma = 1).
//        # (Lambert Meertens)
//        # (corrected version; bug discovered by Mike Miller, fixed by LM)
//
//        # Multithreading note: When two threads call this function
//        # simultaneously, it is possible that they will receive the
//        # same return value.  The window is very small though.  To
//        # avoid this, you have to use a lock around all calls.  (I
//        # didn't want to slow this down in the serial case by using a
//        # lock here.)
//
//        random = self.random
//        z = self.gauss_next
//        self.gauss_next = None
//        if z is None:
//            x2pi = random() * TWOPI
//            g2rad = _sqrt(-2.0 * _log(1.0 - random()))
//            z = _cos(x2pi) * g2rad
//            self.gauss_next = _sin(x2pi) * g2rad
//
//        return mu + z * sigma