package aoc2023

import (
	"testing"
)

func TestSpringy(t *testing.T) {
	tests := []SolverTest{
		{
			day:           12,
			version:       "example",
			solver:        &Springy{},
			expectPartOne: NumberResult{21},
		},
		{
			day:           12,
			version:       "full",
			solver:        &Springy{},
			expectPartOne: NumberResult{21071},
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name(), tt.Go)
	}
}
