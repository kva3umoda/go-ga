package crossover

import (
	"testing"

	"github.com/kva3umoda/go-ga/genome"
	"github.com/stretchr/testify/assert"
)

func TestTwoPoint_Crossing(t *testing.T) {
	ind1 := &genome.Individual{Genome: []float64{0, 1, 2, 3, 4, 5}}
	ind2 := &genome.Individual{Genome: []float64{5, 4, 3, 2, 1, 0}}

	cross := TwoPoint()
	for i := 0; i < 100; i++ {
		nind1 := ind1.Clone()
		nind2 := ind2.Clone()
		cross.Crossing(nind1, nind2)
		assert.NotEqualValues(t, ind1.Genome, nind1.Genome, "not equal ind1")
		assert.NotEqualValues(t, ind2.Genome, nind2.Genome, "not equal ind2")
	}
}
