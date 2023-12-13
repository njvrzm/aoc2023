package aoc2023

import (
	"testing"
)

func TestGalaxy(t *testing.T) {
	tests := []SolverTest{
		{
			day:           11,
			version:       "example",
			solver:        &Galaxy{},
			expectPartOne: NumberResult{374},
		},
		{
			day:           11,
			version:       "full",
			solver:        &Galaxy{},
			expectPartOne: NumberResult{10154062},
			expectPartTwo: NumberResult{553083047914},
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name(), tt.Go)
	}
}
