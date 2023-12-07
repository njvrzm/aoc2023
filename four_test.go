package aoc2023

import (
	"testing"
)

func TestDayFour(t *testing.T) {
	tests := []SolverTest{
		{
			day:           4,
			version:       "example",
			solver:        &Lotto{},
			expectPartOne: NumberResult{13},
			expectPartTwo: NumberResult{30},
		},
		{
			day:           4,
			version:       "full",
			solver:        &Lotto{},
			expectPartOne: NumberResult{26443},
			expectPartTwo: NumberResult{6284877},
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name(), tt.Go)
	}
}
