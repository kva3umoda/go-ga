package crossover

import (
	"fmt"
	"testing"

	"github.com/kva3umoda/go-ga/genome"
)

func TestSimulatedBinaryBounded_Crossing(t *testing.T) {
	c := SimulatedBinaryBounded(20, []float64{-20}, []float64{20})
	ind1 := &genome.Individual{Genome: []float64{10}}
	ind2 := &genome.Individual{Genome: []float64{-15}}

	c.Crossing(ind1, ind2)
	fmt.Println(*ind1)
	fmt.Println(*ind2)


}
