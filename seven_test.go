package aoc2023

import "testing"

func TestDaySeven(t *testing.T) {
	tests := []SolverTest{
		{
			day:           7,
			version:       "example",
			solver:        &Game{},
			expectPartOne: NumberResult{6440},
			expectPartTwo: NumberResult{5905},
		},
		{
			day:           7,
			version:       "full",
			solver:        &Game{},
			expectPartOne: NumberResult{249204891},
			expectPartTwo: NumberResult{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name(), tt.Go)
	}
}
