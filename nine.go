package aoc2023

import (
	"bufio"
)

type Extrapolator struct {
	sequences []Sequence
}

func (e *Extrapolator) PartOne() Result {
	return NumberResult{e.Extrapolate(Next)}
}

func (e *Extrapolator) PartTwo() Result {
	return NumberResult{e.Extrapolate(Previous)}
}

func (e *Extrapolator) Extrapolate(integrator func(int, []int) int) int {
	total := 0
	for _, seq := range e.sequences {
		total += seq.ExtendWith(integrator)
	}
	return total
}

func (e *Extrapolator) Load(scanner *bufio.Scanner) {
	e.sequences = make([]Sequence, 0)
	for scanner.Scan() {
		e.sequences = append(e.sequences, AllNumbers(scanner.Text()))
	}
}

type Sequence []int

func (s Sequence) ExtendWith(integrator func(int, []int) int) int {
	seq := s
	derivatives := [][]int{seq}
	for Any(seq, NonZero[int]) {
		seq = Derivative(seq)
		derivatives = append(derivatives, seq)
	}
	extension := 0
	for i := len(derivatives) - 1; i > 0; i-- {
		extension = integrator(extension, derivatives[i-1])
	}
	return extension
}
func Next(a int, seq []int) int {
	return a + Last(seq)
}
func Previous(a int, seq []int) int {
	return seq[0] - a
}

func Derivative(seq []int) []int {
	out := make([]int, len(seq)-1)
	for i := 0; i < len(seq)-1; i++ {
		out[i] = seq[i+1] - seq[i]
	}
	return out
}
