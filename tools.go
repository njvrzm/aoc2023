package aoc2023

import (
	"bufio"
	"cmp"
	"fmt"
	"golang.org/x/exp/constraints"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func AllNumbers(line string) []int {
	ns, _ := Alltoi(regexp.MustCompile(`-?[0-9]+`).FindAllString(line, -1))
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
	intervals []Interval
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
	out := make([]Interval, 0, len(s.intervals)+len(o.intervals))

	for sIndex < len(s.intervals) && oIndex < len(o.intervals) {
		sInt := s.intervals[sIndex]
		oInt := o.intervals[oIndex]
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
	if s.intervals[0].low == math.MinInt {
		bounds = bounds[1 : len(bounds)-1]
	} else {
		wrapped := make([]int, 0, 2*(len(s.intervals)+1))
		wrapped = append(wrapped, math.MinInt)
		wrapped = append(wrapped, bounds...)
		wrapped = append(wrapped, math.MaxInt)
		bounds = wrapped
	}
	return NewSet(bounds...)
}

func (s Set) IsEmpty() bool {
	return len(s.intervals) == 0
}

func (s Set) Union(o Set) Set {
	return s.Inverse().Intersect(o.Inverse()).Inverse()
}

func (s Set) bounds() []int {
	bounds := make([]int, 2*len(s.intervals))
	for i, interval := range s.intervals {
		bounds[2*i] = interval.low
		bounds[2*i+1] = interval.high
	}
	return bounds
}
func (s Set) Minus(o Set) Set {
	return s.Intersect(o.Inverse())
}

func (s Set) Equals(o Set) bool {
	if len(s.intervals) != len(o.intervals) {
		return false
	}
	for i := 0; i < len(s.intervals); i++ {
		if s.intervals[i] != o.intervals[i] {
			return false
		}
	}
	return true
}
func (s Set) Translate(offset int) Set {
	outIntervals := make([]Interval, len(s.intervals))
	for i, interval := range s.intervals {
		outIntervals[i] = interval.Translate(offset)
	}
	return Set{outIntervals}
}

// Interval represents a half-open range of integers. If high <= low,
// the interval is empty.
type Interval struct {
	low  int
	high int
}

// Translate returns an interval with the same size but with offset added
// to each bound.
func (i Interval) Translate(offset int) Interval {
	return Interval{i.low + offset, i.high + offset}
}

// Intersect returns an interval covering all integers that are in both
// i and o. This may be empty.
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

func GCD(a int, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
func LCM(ns ...int) int {
	out := 1
	for _, n := range ns {
		out = out * n / GCD(out, n)
	}
	return out
}

func ReadBlank(scanner *bufio.Scanner) {
	scanner.Scan()
	if scanner.Text() != "" {
		panic(fmt.Sprintf("Expected empty line, but got: %q", scanner.Text()))
	}
}
func All[T any](seq []T, check func(T) bool) bool {
	for _, it := range seq {
		if !check(it) {
			return false
		}
	}
	return true
}
func Any[T any](seq []T, check func(T) bool) bool {
	for _, it := range seq {
		if check(it) {
			return true
		}
	}
	return false
}

func Last[T any](seq []T) (result T) {
	if len(seq) > 0 {
		result = seq[len(seq)-1]
	}
	return
}

func NonZero[T constraints.Integer | constraints.Float](n T) bool {
	return n != 0
}
func Abs[T constraints.Integer | constraints.Float](n T) T {
	if n >= 0 {
		return n
	}
	return -n
}
func MinMax[T constraints.Integer | constraints.Float](m, n T) (T, T) {
	if m <= n {
		return m, n
	}
	return n, m
}

type Grid struct {
	content map[Place]byte
	width   int
	height  int
}

func (g *Grid) Rows() []string {
	out := make([]string, g.height)
	for y := 0; y < g.width; y++ {
		line := strings.Builder{}
		for x := 0; x < g.width; x++ {
			line.WriteByte(g.content[Place{x, y}])
		}
		out[y] = line.String()
	}
	return out
}
func (g *Grid) Columns() []string {
	out := make([]string, g.height)
	for x := 0; x < g.width; x++ {
		line := strings.Builder{}
		for y := 0; y < g.width; y++ {
			line.WriteByte(g.content[Place{x, y}])
		}
		out[x] = line.String()
	}
	return out
}

func LoadGrid(scanner *bufio.Scanner) *Grid {
	content := make(map[Place]byte)
	var row, col = 0, 0
	for row = 0; scanner.Scan(); row++ {
		line := scanner.Text()
		for col = 0; col < len(line); col++ {
			content[Place{X: col, Y: row}] = line[col]
		}
	}
	return &Grid{content, col, row}
}
func (g *Grid) ScanRows() chan Place {
	out := make(chan Place)
	go func() {
		defer close(out)
		for y := 0; y < g.height; y++ {
			for x := 0; x < g.width; x++ {
				out <- Place{x, y}
			}
		}
	}()
	return out
}
func (g *Grid) ScanColumns() chan Place {
	out := make(chan Place)
	go func() {
		defer close(out)
		for x := 0; x < g.width; x++ {
			for y := 0; y < g.height; y++ {
				out <- Place{x, y}
			}
		}
	}()
	return out
}
func (g *Grid) ScanLines() chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		b := strings.Builder{}
		for place := range g.ScanRows() {
			if place.X == 0 {
				line := b.String()
				if line != "" {
					out <- line
				}
				b.Reset()
			}
			b.WriteByte(g.content[place])
		}
		out <- b.String()
	}()
	return out
}
func In[T comparable](thing T, list []T) bool {
	return Any(list, func(it T) bool { return it == thing })
}

func Apply[T, U any](in []T, f func(T) U) []U {
	out := make([]U, len(in))
	for i, it := range in {
		out[i] = f(it)
	}
	return out
}

func Equal[T comparable](one []T, two []T) bool {
	if len(one) != len(two) {
		return false
	}
	for i := 0; i < len(one); i++ {
		if one[i] != two[i] {
			return false
		}
	}
	return true
}
