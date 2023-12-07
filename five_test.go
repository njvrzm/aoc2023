package aoc2023

import "testing"

func TestDayFive(t *testing.T) {
	tests := []SolverTest{
		{
			day:           5,
			version:       "example",
			solver:        &Almanac{},
			expectPartOne: NumberResult{35},
			expectPartTwo: NumberResult{46},
		},
		{
			day:           5,
			version:       "full",
			solver:        &Almanac{},
			expectPartOne: NumberResult{806029445},
			expectPartTwo: NumberResult{59370572},
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name(), tt.Go)
	}
}
