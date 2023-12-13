package aoc2023

import (
	"bufio"
	"fmt"
	"strings"
)

type Labyrinth struct {
	grid  *Grid
	path  []Place
	start Place
}

var Outlets = map[byte][]Place{
	'|': {Up, Down},
	'-': {Left, Right},
	'J': {Up, Left},
	'7': {Down, Left},
	'L': {Up, Right},
	'F': {Down, Right},
}

func (l *Labyrinth) Outlets(p Place) []Place {
	outlets := make([]Place, 2)
	for i, out := range Outlets[l.grid.content[p]] {
		outlets[i] = p.Plus(out)
	}
	return outlets
}

func (l *Labyrinth) PartOne() Result {
	return NumberResult{value: len(l.path) / 2}
}

func (l *Labyrinth) PartTwo() Result {
	// rewrite path
	onPath := make(map[Place]bool)
	for _, place := range l.path {
		onPath[place] = true
	}
	for place := range l.grid.ScanRows() {
		if !onPath[place] {
			l.grid.content[place] = '.'
		}
	}
	area := 0
	for line := range l.grid.ScanLines() {
		line = strings.ReplaceAll(line, "-", "")
		line = strings.ReplaceAll(line, "LJ", "")
		line = strings.ReplaceAll(line, "F7", "")
		line = strings.ReplaceAll(line, "L7", "|")
		line = strings.ReplaceAll(line, "FJ", "|")
		in := false
		for i := 0; i < len(line); i++ {
			if line[i] == '|' {
				in = !in
			} else if in {
				area += 1
			}
		}
	}
	return NumberResult{value: area}
}

var Infill = map[[4]bool]byte{
	[4]bool{true, false, true, false}: '|',
	[4]bool{true, true, false, false}: 'L',
	[4]bool{true, false, false, true}: 'J',
	[4]bool{false, true, false, true}: '-',
	[4]bool{false, false, true, true}: '7',
	[4]bool{false, true, true, false}: 'F',
}

func (l *Labyrinth) FixStart() {
	up := In(l.start, l.Outlets(l.start.Up()))
	right := In(l.start, l.Outlets(l.start.Right()))
	down := In(l.start, l.Outlets(l.start.Down()))
	left := In(l.start, l.Outlets(l.start.Left()))
	l.grid.content[l.start] = Infill[[4]bool{up, right, down, left}]
}

func (l *Labyrinth) Print() {
	for line := range l.grid.ScanLines() {
		fmt.Println(line)
	}
}
func (l *Labyrinth) Load(scanner *bufio.Scanner) {
	l.grid = LoadGrid(scanner)
	for place := range l.grid.ScanRows() {
		if l.grid.content[place] == 'S' {
			l.start = place
			l.FixStart()
			break
		}
	}
	l.path = []Place{l.start}
	beenTo := map[Place]bool{l.start: true}

	done := false
	for !done {
		done = true
		for _, outlet := range l.Outlets(Last(l.path)) {
			if beenTo[outlet] {
				continue
			}
			l.path = append(l.path, outlet)
			beenTo[outlet] = true
			done = false
		}
	}
}
