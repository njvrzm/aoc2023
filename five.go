package aoc2023

import (
	"bufio"
	"strings"
)

type Projection struct {
	source Set
	offset int
}

func (p Projection) Project(s Set) Set {
	return s.Intersect(p.source).Shift(p.offset)
}

type Almanac struct {
	seeds []int
	maps  []Map
}

func (alm *Almanac) PartOne() Result {
	seeds := NewSet()
	for _, seed := range alm.seeds {
		seeds = seeds.Union(NewSet(seed, seed+1))
	}
	projected := alm.Project(seeds)
	return NumberResult{projected.bounds()[0]}
}

func (alm *Almanac) PartTwo() Result {
	seeds := NewSet()
	for i := 0; i < len(alm.seeds); i += 2 {
		lower, length := alm.seeds[i], alm.seeds[i+1]
		seeds = seeds.Union(NewSet(0, length).Shift(lower))
	}
	projected := alm.Project(seeds)
	return NumberResult{projected.bounds()[0]}
}

func (alm *Almanac) Load(scanner *bufio.Scanner) {
	alm.seeds = ReadSeeds(scanner)
	ReadBlank(scanner)

	alm.maps = make([]Map, 0)
	for scanner.Scan() {
		alm.maps = append(alm.maps, ReadMap(scanner))
	}
}

func ReadSeeds(scanner *bufio.Scanner) []int {
	scanner.Scan()
	return AllNumbers(scanner.Text())
}

func (alm *Almanac) Project(s Set) Set {
	for _, m := range alm.maps {
		s = m.Project(s)
	}
	return s
}

type Map struct {
	category    string
	projections []Projection
}

func ReadMap(scanner *bufio.Scanner) Map {
	category := strings.SplitN(scanner.Text(), "-", 2)[0]
	m := Map{category: category, projections: make([]Projection, 0)}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		m.projections = append(m.projections, ReadProjection(line))
	}
	return m
}
func ReadProjection(line string) Projection {
	numbers := AllNumbers(line)
	low := numbers[1]
	high := numbers[1] + numbers[2]
	offset := numbers[0] - numbers[1]
	return Projection{NewSet(low, high), offset}
}

func (gt *Map) Project(s Set) Set {
	projected := NewSet()
	for _, p := range gt.projections {
		projected = projected.Union(p.Project(s))
		s = s.Minus(p.source)
	}
	return projected.Union(s)
}
