package aoc2023

import (
	"bufio"
	"strings"
)

type Galaxy struct {
	grid         *Grid
	stars        []Place
	emptyRows    map[int]bool
	emptyColumns map[int]bool
}

func (g *Galaxy) PartOne() Result {
	total := 0
	for o := 0; o < len(g.stars); o++ {
		for t := o + 1; t < len(g.stars); t++ {
			total += g.Distance(o, t, 1)
		}
	}
	return NumberResult{total}
}

func (g *Galaxy) PartTwo() Result {
	total := 0
	for o := 0; o < len(g.stars); o++ {
		for t := o + 1; t < len(g.stars); t++ {
			total += g.Distance(o, t, 1000000-1)
		}
	}
	return NumberResult{total}
}

func (g *Galaxy) Distance(o int, t int, expansion int) int {
	one := g.stars[o]
	two := g.stars[t]
	lx, hx := MinMax[int](one.X, two.X)
	ly, hy := MinMax[int](one.Y, two.Y)
	base := (hx - lx) + (hy - ly)
	for r := lx + 1; r < hx; r++ {
		if g.emptyColumns[r] {
			base += expansion
		}
	}
	for c := ly + 1; c < hy; c++ {
		if g.emptyRows[c] {
			base += expansion
		}
	}
	return base

}

func (g *Galaxy) Load(scanner *bufio.Scanner) {
	g.emptyRows = make(map[int]bool)
	g.emptyColumns = make(map[int]bool)
	g.grid = LoadGrid(scanner)
	for y, row := range g.grid.Rows() {
		stars := false
		for x := 0; x < len(row); x++ {
			if row[x] == '#' {
				stars = true
				g.stars = append(g.stars, Place{x, y})
			}
		}
		if !stars {
			g.emptyRows[y] = true
		}
	}
	for x, col := range g.grid.Columns() {
		if !strings.ContainsRune(col, '#') {
			g.emptyColumns[x] = true
		}
	}
}
