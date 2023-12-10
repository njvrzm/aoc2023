package aoc2023

import "testing"

func TestDayEight(t *testing.T) {
	tests := []SolverTest{
		//{
		//	day:           8,
		//	version:       "example",
		//	subversion:    "1",
		//	solver:        &Network{},
		//	expectPartOne: NumberResult{2},
		//},
		//{
		//	day:           8,
		//	version:       "example",
		//	subversion:    "2",
		//	solver:        &Network{},
		//	expectPartOne: NumberResult{6},
		//},
		//{
		//	day:           8,
		//	version:       "example",
		//	subversion:    "3",
		//	solver:        &Network{},
		//	expectPartTwo: NumberResult{6},
		//},
		{
			day:     8,
			version: "full",
			solver:  &Network{},
			//expectPartOne: NumberResult{20513},
			expectPartTwo: NumberResult{15995167053923},
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name(), tt.Go)
	}
}
