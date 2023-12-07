package aoc2023

import "testing"

func TestDayFive(t *testing.T) {
	tests := []SolverTest{
		{
			day:           5,
			version:       "example",
			solver:        &Garden{},
			expectPartOne: NumberResult{35},
			expectPartTwo: NotImplemented,
		},
		{
			day:           5,
			version:       "full",
			solver:        &Garden{},
			expectPartOne: NumberResult{806029445},
			expectPartTwo: NotImplemented,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name(), tt.Go)
	}
}
