package aoc2023

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
)

type Springy struct {
	reports []Report
}

func (s *Springy) PartOne() Result {
	counts := 0
	oc := make([]int, len(s.reports))
	for j, report := range s.reports {
		degree := len(report.questions)
		count := 0
		record := []byte(report.record)
		for i := 0; i < 1<<degree; i++ {
			factor := 1
			for j := 0; j < degree; j++ {
				if i&(factor<<j) != 0 {
					record[report.questions[j]] = '#'
				} else {
					record[report.questions[j]] = '.'
				}
			}
			foo := Report{record: string(record), blocks: report.blocks, questions: report.questions}
			if foo.Matches() {
				count += 1
			}
		}
		counts += count
		oc[j] = count
	}
	fmt.Println(oc)
	return NumberResult{counts}
}

func (s *Springy) PartTwo() Result {
	return NotImplemented
}

func (s *Springy) Load(scanner *bufio.Scanner) {
	s.reports = make([]Report, 0)
	for scanner.Scan() {
		s.reports = append(s.reports, ReadReport(scanner.Text()))
	}
}

type Report struct {
	record    string
	questions []int
	blocks    []int
}

var brokens = regexp.MustCompile(`#+`)

func (r Report) Matches() bool {
	lengths := Apply(brokens.FindAllString(r.record, -1), func(s string) int { return len(s) })
	return Equal[int](lengths, r.blocks)
}

func ReadReport(line string) Report {
	parts := strings.SplitN(line, " ", 2)
	qs := make([]int, 0)
	for i := 0; i < len(parts[0]); i++ {
		if parts[0][i] == '?' {
			qs = append(qs, i)
		}
	}
	return Report{
		record:    parts[0],
		questions: qs,
		blocks:    AllNumbers(parts[1]),
	}

}
