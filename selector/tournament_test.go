package selector

import (
	"sort"
	"testing"

	"ga-book/internal/genome"
	"github.com/stretchr/testify/assert"
)

func TestTournament_Select(t *testing.T) {
	inds := []genome.Individual{
		{Fitness: 101, Genome: []float64{1}},
		{Fitness: 102, Genome: []float64{2}},
		{Fitness: 103, Genome: []float64{3}},
		{Fitness: 104, Genome: []float64{4}},
		{Fitness: 105, Genome: []float64{5}},
		{Fitness: 106, Genome: []float64{6}},
		{Fitness: 107, Genome: []float64{7}},
		{Fitness: 108, Genome: []float64{8}},
		{Fitness: 109, Genome: []float64{9}},
	}

	selector := NewTournament(2)

	for i := 0; i < 100; i++ {
		indsSel := selector.Select(len(inds), inds)
		sort.Slice(indsSel, func(i, j int) bool {
			return indsSel[i].Fitness < indsSel[j].Fitness
		})

		assert.NotEqualValues(t, inds, indsSel, "select tournament")
	}
}
