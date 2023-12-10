package aoc2023

import (
	"testing"
)

func TestExtrapolator(t *testing.T) {
	tests := []SolverTest{
		{
			day:           9,
			version:       "example",
			solver:        &Extrapolator{},
			expectPartOne: NumberResult{114},
			expectPartTwo: NumberResult{2},
		},
		{
			day:           9,
			version:       "full",
			solver:        &Extrapolator{},
			expectPartOne: NumberResult{1702218515},
			expectPartTwo: NumberResult{925},
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name(), tt.Go)
	}
}
