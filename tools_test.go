package aoc2023

import (
	"math"
	"testing"
)

func TestSet_Intersection(t *testing.T) {
	tests := []struct {
		name string
		one  Set
		two  Set
		want Set
	}{
		{
			name: "single intervals, no overlap",
			one:  NewSet(1, 5),
			two:  NewSet(6, 10),
			want: NewSet(),
		},
		{
			name: "same interval, same out",
			one:  NewSet(1, 5),
			two:  NewSet(1, 5),
			want: NewSet(1, 5),
		},
		{
			name: "XX...XX meets .XXXXX.",
			one:  NewSet(1, 3, 6, 8),
			two:  NewSet(2, 7),
			want: NewSet(2, 3, 6, 7),
		},
		{
			name: "matching upper bounds",
			one:  NewSet(1, 8, 12, 20),
			two:  NewSet(3, 8, 15, 20),
			want: NewSet(3, 8, 15, 20),
		},
		{
			name: "matching offset bounds",
			one:  NewSet(8, 12, 20, 21),
			two:  NewSet(3, 8, 15, 20),
			want: NewSet(),
		},
		{
			name: "more complex",
			one:  NewSet(1, 3, 10, 20, 35, 55),
			two:  NewSet(15, 40, 50, 60),
			want: NewSet(15, 20, 35, 40, 50, 55),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.one.Intersect(tt.two)
			if !got.Equals(tt.want) {
				t.Errorf("Intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Minus(t *testing.T) {
	tests := []struct {
		name     string
		original Set
		subtract Set
		want     Set
	}{
		{
			name:     "no overlap, no change",
			original: NewSet(1, 10),
			subtract: NewSet(20, 30),
			want:     NewSet(1, 10),
		},
		{
			name:     "bite out of middle",
			original: NewSet(1, 10),
			subtract: NewSet(3, 6),
			want:     NewSet(1, 3, 6, 10),
		},
		{
			name:     "all gone",
			original: NewSet(1, 10),
			subtract: NewSet(0, 11),
			want:     NewSet(),
		},
		{
			name:     "one side",
			original: NewSet(1, 10),
			subtract: NewSet(5, 20),
			want:     NewSet(1, 5),
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.original.Minus(tt.subtract); !tt.want.Equals(got) {
				t.Errorf("Minus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Inverse(t *testing.T) {
	tests := []struct {
		name     string
		original Set
		inverted Set
	}{
		{
			"two to three",
			NewSet(1, 3, 5, 10),
			NewSet(math.MinInt, 1, 3, 5, 10, math.MaxInt),
		},
		{
			"and back again",
			NewSet(math.MinInt, 1, 3, 5, 10, math.MaxInt),
			NewSet(1, 3, 5, 10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.original.Inverse(); !got.Equals(tt.inverted) {
				t.Errorf("Inverse() = %v, want %v", got, tt.inverted)
			}
		})
	}
}

func TestSet_Union(t *testing.T) {
	tests := []struct {
		name     string
		original Set
		added    Set
		expected Set
	}{
		{
			"same in, same out",
			NewSet(1, 5),
			NewSet(1, 5),
			NewSet(1, 5),
		},
		{
			"overlap",
			NewSet(1, 5),
			NewSet(4, 10),
			NewSet(1, 10),
		},
		{
			"distinct",
			NewSet(1, 5),
			NewSet(7, 10),
			NewSet(1, 5, 7, 10),
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.original.Union(tt.added); !got.Equals(tt.expected) {
				t.Errorf("Union() = %v, want %v", got, tt.expected)
			}
		})
	}
}
