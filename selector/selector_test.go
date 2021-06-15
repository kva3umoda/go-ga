package selectors

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_TestSelect(t *testing.T) {
	ind1 := []float64{0, 1, 2, 3, 4, 5}
	ind2 := []float64{10, 11, 12, 13, 14, 15, 16, 17}

	cxpoint := 5

	tmp := make([]float64, cxpoint)
	copy(tmp, ind1[:cxpoint])

	copy(ind1[:cxpoint], ind2[:cxpoint])
	copy(ind2[:cxpoint], tmp)

	fmt.Println(ind1)
	fmt.Println(ind2)
}

func Test_TestSelect2(t *testing.T) {
	for i := 0; i < 100; i++ {
		cxpoint := rand.Intn(6-1) + 1
		fmt.Println(cxpoint)
	}
}
