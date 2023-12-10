package aoc2023

import (
	"testing"
)

func TestLabyrinth(t *testing.T) {
	tests := []SolverTest{
		{
			day:           10,
			version:       "example",
			solver:        &Labyrinth{},
			expectPartOne: NumberResult{8},
		},
		{
			day:           10,
			version:       "example",
			subversion:    "1",
			solver:        &Labyrinth{},
			expectPartTwo: NumberResult{8},
		},
		{
			day:           10,
			version:       "full",
			solver:        &Labyrinth{},
			expectPartOne: NumberResult{7093},
			expectPartTwo: NumberResult{1432},
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name(), tt.Go)
	}
}
