package aoc2023

import (
	"bufio"
	"regexp"
	"strconv"
)

type Part struct {
	place   Place
	value   string
	touches map[*Part]bool
}

func (c *Part) TouchesSymbol() bool {
	for other := range c.touches {
		if other.IsSymbol() {
			return true
		}
	}
	return false
}
func (c *Part) IsSymbol() bool {
	return c.value[0] < '0' || c.value[0] > '9'
}
func (c *Part) Number() int {
	if c.IsSymbol() {
		return 0
	}
	return Must(strconv.Atoi(c.value))
}
func (c *Part) Touch(o *Part) {
	if o != nil {
		c.touches[o] = true
	}
}
func (c *Part) Neighbors() chan Place {
	return c.place.RectangleNeighbors(len(c.value), 1)
}

var reNonempty = regexp.MustCompile(`[0-9]+|[^.0-9]`)

type Schematic struct {
	diagram map[Place]*Part
	parts   []*Part
}

func (s *Schematic) AddPart(x, y int, value string) {
	where := Place{x, y}
	part := &Part{place: where, value: value, touches: make(map[*Part]bool)}
	s.parts = append(s.parts, part)
	for col := x; col < x+len(value); col++ {
		s.diagram[Place{col, y}] = part
	}
}
func (s *Schematic) TouchEverything() {
	for _, part := range s.parts {
		for neighbor := range part.Neighbors() {
			part.Touch(s.diagram[neighbor])
		}
	}
}
func (s *Schematic) ReadLine(row int, line string) {
	for _, bounds := range reNonempty.FindAllStringIndex(line, -1) {
		start, end := bounds[0], bounds[1]
		s.AddPart(start, row, line[start:end])
	}
}
func (s *Schematic) Load(scanner *bufio.Scanner) {
	s.diagram = make(map[Place]*Part)
	s.parts = make([]*Part, 0)
	for row := 0; scanner.Scan(); row++ {
		s.ReadLine(row, scanner.Text())
	}
	s.TouchEverything()
}
func (s *Schematic) PartOne() Result {
	sum := 0
	for _, part := range s.parts {
		if !part.IsSymbol() && part.TouchesSymbol() {
			sum += part.Number()
		}
	}
	return NumberResult{sum}
}
func (s *Schematic) PartTwo() Result {
	sum := 0
	touched := make([]*Part, 0)
	for _, part := range s.parts {
		if part.value == "*" {
			for other := range part.touches {
				if !other.IsSymbol() {
					touched = append(touched, other)
				}
			}
			if len(touched) == 2 {
				sum += touched[0].Number() * touched[1].Number()
			}
		}
		touched = touched[:0]
	}
	return NumberResult{sum}
}
