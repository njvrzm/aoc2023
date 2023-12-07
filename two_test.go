package aoc2023

import (
	"fmt"
	"os"
	"testing"
)

// todo: update test style
func TestDayTwoPartOne(t *testing.T) {
	t.Run("Day Two, Part One (full)", func(t *testing.T) {
		got := DayTwoPartOne(Must(os.Open("inputs/2")))
		fmt.Printf("two.one: %d\n", got)
	})
}

func TestDayTwoPartTwo(t *testing.T) {
	t.Run("Day Two, Part Two (full)", func(t *testing.T) {
		got := DayTwoPartTwo(Must(os.Open("inputs/2")))
		fmt.Printf("two.two: %d\n", got)
	})
}
