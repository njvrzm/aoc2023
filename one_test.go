package aoc2023

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
func TestDayOnePartOne_Full(t *testing.T) {
	t.Run("Day one, part one", func(t *testing.T) {
		expected := 56049
		got := DayOnePartOne(Must(os.Open("inputs/1")))
		if expected != got {
			t.Errorf("one.one: %d != %d", expected, got)
		}
	})
}

func bench[T any](f func() T) T {
	before := time.Now()
	defer fmt.Printf("Duration: %s\n", time.Since(before))
	return f()
}
func TestDayOnePartTwo_Full(t *testing.T) {
	t.Run("Day one, part two (full)", func(t *testing.T) {
		expected := 54530
		got := bench(func() int {
			return DayOnePartTwo(Must(os.Open("inputs/1")))
		})
		if expected != got {
			t.Errorf("one.one: %d != %d", expected, got)
		}
	})
}
func TestDayOnePartTwo_Stem_Full(t *testing.T) {
	t.Run("Day one, part two (full, stem method)", func(t *testing.T) {
		expected := 54530
		got := DayOnePartTwoStem(Must(os.Open("inputs/1")))
		if expected != got {
			t.Errorf("one.one: %d != %d", expected, got)
		}
	})
}
func TestDayOnePartTwo_Steam_Full(t *testing.T) {
	t.Run("Day one, part two (full, stem method)", func(t *testing.T) {
		expected := 54530
		got := DayOnePartTwoSteam(Must(os.Open("inputs/1")))
		if expected != got {
			t.Errorf("one.one: %d != %d", expected, got)
		}
	})
}
func BenchmarkDayOnePartTwo(b *testing.B) {
	b.Run("Day one, part two", func(b *testing.B) {
		for i := 1; i <= b.N; i++ {
			DayOnePartTwo(Must(os.Open(fmt.Sprintf("testdata/dayone/%d", i))))
		}
	})
}
func BenchmarkDayOnePartTwo_Stem(b *testing.B) {
	b.Run("Day one, part two", func(b *testing.B) {
		for i := 1; i <= b.N; i++ {
			DayOnePartTwoStem(Must(os.Open(fmt.Sprintf("testdata/dayone/%d", i))))
		}
	})
}
func BenchmarkDayOnePartTwo_Steam(b *testing.B) {
	b.Run("Day one, part two", func(b *testing.B) {
		for i := 1; i <= b.N; i++ {
			DayOnePartTwoSteam(Must(os.Open(fmt.Sprintf("testdata/dayone/%d", i))))
		}
	})
}

func TestDayOnePartOne_Unit(t *testing.T) {
	tests := []struct {
		name string
		text string
		want int
	}{
		{
			"two lines",
			"a91b5c\n44444x0000x1x\n",
			95 + 41,
		},
		{
			"only one digit",
			"a9c\n",
			99,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DayOnePartOne(strings.NewReader(tt.text)); got != tt.want {
				t.Errorf("DayOnePartOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDayOnePartTwo_Unit(t *testing.T) {
	tests := []struct {
		name string
		text string
		want int
	}{
		{
			"two lines",
			"onetwo3\nnonsenseven\n",
			13 + 77,
		},
		{
			"only one digit",
			"afivec\n",
			55,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DayOnePartTwo(strings.NewReader(tt.text)); got != tt.want {
				t.Errorf("DayOnePartTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
