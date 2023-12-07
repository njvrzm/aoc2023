package aoc2023

import (
	"bufio"
	"cmp"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func AllNumbers(line string) []int {
	ns, _ := Alltoi(regexp.MustCompile(`[0-9]+`).FindAllString(line, -1))
	return ns
}
func Alltoi(nss []string) ([]int, error) {
	// may be over-allocated
	out := make([]int, len(nss))
	for i, ns := range nss {
		if ns != "" {
			n, err := strconv.Atoi(ns)
			if err != nil {
				return nil, err
			}
			out[i] = n
		}
	}
	return out, nil
}

// Set represents a subset of the integers with a list of half-open intervals.
// A single integer n is represented as [[n, n+1]]; an unbroken interval of length
// m by [[n, n+m]]; and, for example, the 20s, 50s and 90s by [[20,30],[50,60],[90,100]].
type Set struct {
	contains []Interval
}

func NewSet(bounds ...int) Set {
	intervals := make([]Interval, len(bounds)/2)
	for i := 0; i < len(bounds); i += 2 {
		intervals[i/2] = Interval{low: bounds[i], high: bounds[i+1]}
	}
	return Set{intervals}
}

func (s Set) Intersect(o Set) Set {
	sIndex := 0
	oIndex := 0
	out := make([]Interval, 0, len(s.contains)+len(o.contains))

	for sIndex < len(s.contains) && oIndex < len(o.contains) {
		sInt := s.contains[sIndex]
		oInt := o.contains[oIndex]
		outInt := sInt.Intersect(oInt)
		if !outInt.IsEmpty() {
			out = append(out, outInt)
		}
		if sInt.high <= oInt.high {
			sIndex += 1
		}
		if oInt.high <= sInt.high {
			oIndex += 1
		}
	}
	return Set{out}
}

var Everything = NewSet(math.MinInt, math.MaxInt)

func (s Set) Inverse() Set {
	bounds := s.bounds()
	if len(bounds) == 0 {
		return Everything
	}
	if s.contains[0].low == math.MinInt {
		bounds = bounds[1 : len(bounds)-1]
	} else {
		wrapped := make([]int, 0, 2*(len(s.contains)+1))
		wrapped = append(wrapped, math.MinInt)
		wrapped = append(wrapped, bounds...)
		wrapped = append(wrapped, math.MaxInt)
		bounds = wrapped
	}
	return NewSet(bounds...)
}

func (s Set) IsEmpty() bool {
	return len(s.contains) == 0
}

func (s Set) Union(o Set) Set {
	return s.Inverse().Intersect(o.Inverse()).Inverse()
}

func (s Set) bounds() []int {
	bounds := make([]int, 2*len(s.contains))
	for i, interval := range s.contains {
		bounds[2*i] = interval.low
		bounds[2*i+1] = interval.high
	}
	return bounds
}
func (s Set) Minus(o Set) Set {
	return s.Intersect(o.Inverse())
}

func (s Set) Equals(o Set) bool {
	if len(s.contains) != len(o.contains) {
		return false
	}
	for i := 0; i < len(s.contains); i++ {
		if s.contains[i] != o.contains[i] {
			return false
		}
	}
	return true
}
func (s Set) Shift(offset int) Set {
	outIntervals := make([]Interval, len(s.contains))
	for i, interval := range s.contains {
		outIntervals[i] = interval.Shift(offset)
	}
	return Set{outIntervals}
}

// Interval represents a half-open range of integers. [n, n] is empty;
// [n, n+1] contains only n; [n, n+m] contains m consecutive integers starting with n
type Interval struct {
	low  int
	high int
}

func (i Interval) Shift(offset int) Interval {
	return Interval{i.low + offset, i.high + offset}
}

func (i Interval) Intersect(o Interval) Interval {
	outLow := Max(i.low, o.low)
	outHigh := Min(i.high, o.high)
	return Interval{low: outLow, high: outHigh}
}

func (i Interval) IsEmpty() bool {
	return i.high <= i.low
}

func Max[T cmp.Ordered](one T, two T) T {
	if one < two {
		return two
	} else {
		return one
	}
}

func Min[T cmp.Ordered](one T, two T) T {
	if one < two {
		return one
	} else {
		return two
	}
}

func ReadBlank(scanner *bufio.Scanner) {
	scanner.Scan()
	if scanner.Text() != "" {
		panic(fmt.Sprintf("Expected empty line, but got: %q", scanner.Text()))
	}
}
