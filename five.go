package aoc2023

import (
	"bufio"
	"math"
	"strings"
)

type Remap struct {
	target int
	source int
	length int
}

func (r Remap) Map(n int) (int, bool) {
	if r.source <= n && n < r.source+r.length {
		return r.target + n - r.source, true
	}
	return 0, false
}

type Garden struct {
	seeds []int
	tree  *GardenTree
}

func (g *Garden) PartOne() Result {
	lowestLocation := math.MaxInt
	for _, seed := range g.seeds {
		location := g.RemapFully(seed)
		if location < lowestLocation {
			lowestLocation = location
		}
	}
	return NumberResult{value: lowestLocation}
}

func (g *Garden) PartTwo() Result {
	return NotImplemented
}

func (g *Garden) Load(scanner *bufio.Scanner) Solver {
	// read seed line
	scanner.Scan()
	g.seeds = AllNumbers(scanner.Text())

	tree := &GardenTree{category: "seed", remappers: make([]Remap, 0)}
	g.tree = tree

	for scanner.Scan() {
		tree.ReadLine(scanner.Text())
		if tree.nextStep != nil {
			tree = tree.nextStep
		}
	}
	return g
}

func (g *Garden) RemapFully(n int) int {
	tree := g.tree
	for tree != nil {
		n = tree.Remap(n)
		tree = tree.nextStep
	}
	return n
}

type GardenTree struct {
	category  string
	nextStep  *GardenTree
	remappers []Remap
}

func NewGardenTree(category string) *GardenTree {
	return &GardenTree{category: category, remappers: make([]Remap, 0)}
}
func (gt *GardenTree) ReadLine(line string) {
	if line == "" {
		return
	}
	numbers := AllNumbers(line)
	if len(numbers) > 0 {
		gt.remappers = append(gt.remappers, Remap{numbers[0], numbers[1], numbers[2]})
		return
	}
	if strings.HasSuffix(line, " map:") {
		category := strings.SplitN(line, "-", 2)[0]
		gt.nextStep = NewGardenTree(category)
	}
}

func (gt *GardenTree) Remap(n int) int {
	for _, rm := range gt.remappers {
		out, yes := rm.Map(n)
		if yes {
			return out
		}
	}
	return n
}
