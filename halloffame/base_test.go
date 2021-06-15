package halloffame

import (
	"testing"

	"github.com/kva3umoda/go-ga/genome"
	"github.com/kva3umoda/go-ga/population"
	"github.com/stretchr/testify/assert"
)

func TestBaseOne_Update(t *testing.T) {
	halloffame := NewBase(1)

	tests := []struct {
		name   string
		input  []*genome.Individual
		output []*genome.Individual
	}{
		{
			name: "first",
			input: []*genome.Individual{
				{Fitness: 5, Genome: []float64{5000}},
				{Fitness: 6, Genome: []float64{6000}},
				{Fitness: 7, Genome: []float64{7000}},
				{Fitness: 3, Genome: []float64{3000}},
				{Fitness: 1, Genome: []float64{1000}},
				{Fitness: 2, Genome: []float64{2000}},
			},
			output: []*genome.Individual{
				{Fitness: 7, Genome: []float64{7000}},
			},
		},

		{
			name: "second",
			input: []*genome.Individual{
				{Fitness: 1, Genome: []float64{1001}},
				{Fitness: 7, Genome: []float64{7001}},
				{Fitness: 5, Genome: []float64{5001}},
				{Fitness: 6, Genome: []float64{6001}},
				{Fitness: 3, Genome: []float64{3001}},
				{Fitness: 2, Genome: []float64{2001}},
			},
			output: []*genome.Individual{
				{Fitness: 7, Genome: []float64{7001}},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			halloffame.Update(&population.Population{Individuals: test.input})
			assert.Equal(t, test.output, halloffame.Individuals(), test.name)
		})
	}

}

func TestBaseThree_Update(t *testing.T) {

	halloffame := NewBase(3)

	tests := []struct {
		name   string
		input  []*genome.Individual
		output []*genome.Individual
	}{
		{
			name: "first",
			input: []*genome.Individual{
				{Fitness: 5, Genome: []float64{5000}},
				{Fitness: 6, Genome: []float64{6000}},
				{Fitness: 7, Genome: []float64{7000}},
				{Fitness: 3, Genome: []float64{3000}},
				{Fitness: 1, Genome: []float64{1000}},
				{Fitness: 2, Genome: []float64{2000}},
			},
			output: []*genome.Individual{
				{Fitness: 7, Genome: []float64{7000}},
				{Fitness: 6, Genome: []float64{6000}},
				{Fitness: 5, Genome: []float64{5000}},
			},
		},
		{
			name: "second",
			input: []*genome.Individual{
				{Fitness: 1, Genome: []float64{1001}},
				{Fitness: 7, Genome: []float64{7001}},
				{Fitness: 3, Genome: []float64{3001}},
				{Fitness: 2, Genome: []float64{2001}},
			},
			output: []*genome.Individual{
				{Fitness: 7, Genome: []float64{7001}},
				{Fitness: 7, Genome: []float64{7000}},
				{Fitness: 6, Genome: []float64{6000}},
			},
		},
		{
			name: "three",
			input: []*genome.Individual{
				{Fitness: 5, Genome: []float64{5002}},
				{Fitness: 6, Genome: []float64{6002}},
				{Fitness: 3, Genome: []float64{3002}},
				{Fitness: 1, Genome: []float64{1002}},
				{Fitness: 2, Genome: []float64{2002}},
			},
			output: []*genome.Individual{
				{Fitness: 7, Genome: []float64{7001}},
				{Fitness: 7, Genome: []float64{7000}},
				{Fitness: 6, Genome: []float64{6002}},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			halloffame.Update(&population.Population{Individuals: test.input})
			assert.Equal(t, test.output, halloffame.Individuals(), test.name)
		})
	}

}
