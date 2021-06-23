package crossover

import (
	"fmt"
	"testing"

	"github.com/kva3umoda/go-ga/genome"
	"github.com/stretchr/testify/assert"
)

func TestUniformPartialMatched_Crossing(t *testing.T) {
	ind1 := genome.Individual{Genome: []float64{0, 1, 2, 3, 4, 5, 6, 7, 8}}
	ind2 := genome.Individual{Genome: []float64{8, 7, 6, 5, 4, 3, 2, 1, 0}}
	ordered := UniformPartialMatched(0.5)

	for i := 0; i < 100; i++ {
		nind1 := ind1.Clone()
		nind2 := ind2.Clone()
		ordered.Crossing(nind1, nind2)
		sum := 0.0
		for _, n := range nind1.Genome {
			sum += n
		}
		assert.Equal(t, 36.0, sum, "nind1")
		fmt.Printf("nind1 : %v\n", nind1)

		sum = 0.0
		for _, n := range nind2.Genome {
			sum += n
		}
		assert.Equal(t, 36.0, sum, "nind2")
		fmt.Printf("nind2 : %v\n", nind2)
	}

}
