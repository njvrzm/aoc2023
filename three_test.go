package aoc2023

import (
	"testing"
)

func TestDayThree(t *testing.T) {
	tests := []SolverTest{
		{
			day:           3,
			version:       "example",
			solver:        &Schematic{},
			expectPartOne: NumberResult{4361},
			expectPartTwo: NumberResult{467835},
		},
		{
			day:           3,
			version:       "full",
			solver:        &Schematic{},
			expectPartOne: NumberResult{519444},
			expectPartTwo: NumberResult{74528807},
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name(), tt.Go)
	}
}
