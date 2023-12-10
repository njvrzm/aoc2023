package aoc2023

import (
	"bufio"
	"fmt"
	"strings"
)

type Labyrinth struct {
	content map[Place]byte
	path    []Place
	been    map[Place]bool
	width   int
	height  int

	inside map[Place]bool
}

func (l *Labyrinth) Outlets(p Place) []Place {
	switch l.content[p] {
	case '|':
		return []Place{p.Up(), p.Down()}
	case '-':
		return []Place{p.Left(), p.Right()}
	case '7':
		return []Place{p.Left(), p.Down()}
	case 'L':
		return []Place{p.Up(), p.Right()}
	case 'J':
		return []Place{p.Left(), p.Up()}
	case 'F':
		return []Place{p.Right(), p.Down()}
	case 'S':
		return l.Inlets(p)
	default:
		return nil
	}
}
func (l *Labyrinth) Inlets(p Place) []Place {
	inlets := make([]Place, 0)
	for neighbor := range p.Neighbors() {
		outlets := l.Outlets(neighbor)
		if Any(outlets, func(o Place) bool { return o == p }) {
			inlets = append(inlets, neighbor)
		}
	}
	return inlets
}

func (l *Labyrinth) PartOne() Result {
	return NumberResult{value: len(l.path) / 2}
}

func (l *Labyrinth) PartTwo() Result {
	area := 0
	for y := 0; y < l.height; y++ {
		in := false
		last := byte(' ')
		for x := 0; x < l.width; x++ {
			where := Place{x, y}
			if !l.been[where] {
				if in {
					l.inside[where] = true
					area += 1
				}
			} else {
				what := l.content[where]
				switch what {
				case '|', '-':
					if what == '|' {
						in = !in
					}
					continue
				case '7':
					if last != 'F' {
						in = !in
					}
				case 'J':
					if last != 'L' {
						in = !in
					}
				}
				last = what
			}
		}
	}
	l.Print()
	return NumberResult{value: area}
}

func (l *Labyrinth) Print() {
	for y := 0; y < l.width; y++ {
		line := strings.Builder{}
		for x := 0; x < l.width; x++ {
			where := Place{x, y}
			if l.inside[where] {
				line.WriteByte('.')
			} else if l.been[where] {
				line.WriteByte(l.content[where])
			} else {
				line.WriteByte(' ')
			}
		}
		fmt.Println(line.String())
	}
}
func (l *Labyrinth) Load(scanner *bufio.Scanner) {
	var start Place
	l.content = make(map[Place]byte)
	l.inside = make(map[Place]bool)
	var row, col int
	for row = 0; scanner.Scan(); row++ {
		var char byte
		for col, char = range []byte(scanner.Text()) {
			l.content[Place{X: col, Y: row}] = char
			if char == 'S' {
				start = Place{X: col, Y: row}
			}
		}
	}
	possible := map[byte]bool{'|': true, '-': true, 'F': true, 'J': true, '7': true, 'L': true}
	if up := l.content[start.Up()]; up == 'J' || up == '-' || up == 'L' {
		delete(possible, '|')
		delete(possible, 'J')
		delete(possible, 'L')
	}
	if right := l.content[start.Right()]; right == 'L' || right == '|' || right == 'F' {
		delete(possible, 'L')
		delete(possible, '-')
		delete(possible, 'F')
	}
	if down := l.content[start.Down()]; down == '-' || down == 'F' || down == '7' {
		delete(possible, '|')
		delete(possible, 'F')
		delete(possible, '7')
	}
	if left := l.content[start.Left()]; left == '|' || left == '7' || left == 'J' {
		delete(possible, '-')
		delete(possible, '7')
		delete(possible, 'J')
	}
	for k := range possible {
		l.content[start] = k
		break
	}

	l.width = col + 1
	l.height = row + 1
	l.path = []Place{start}
	l.been = map[Place]bool{start: true}
	for {
		var next = Nowhere
		for _, outlet := range l.Outlets(Last(l.path)) {
			if l.been[outlet] {
				continue
			}
			next = outlet
		}
		if next == Nowhere {
			return
		}
		l.path = append(l.path, next)
		l.been[next] = true
	}
}
