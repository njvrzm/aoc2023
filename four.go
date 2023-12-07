package aoc2023

import (
	"bufio"
	"regexp"
	"slices"
	"strings"
)

type Lotto struct {
	cards []*Card
}

func (l *Lotto) PartOne() Result {
	score := 0
	for _, game := range l.cards {
		gameScore := game.Score()
		if gameScore > 0 {
			score += 1 << (gameScore - 1)
		}
	}
	return NumberResult{value: score}
}

// PartTwoStupid just generates a giant list. It takes about 750ms on my system
func (l *Lotto) PartTwoStupid() Result {
	cards := make([]*Card, len(l.cards))
	copy(cards, l.cards)
	for i := 0; i < len(cards); i++ {
		score := cards[i].Score()
		for j := cards[i].number; j <= cards[i].number+score-1; j++ {
			cards = append(cards, cards[j])
		}
	}
	return NumberResult{len(cards)}
}

// PartTwo is cleverer: we iterate through the cards in order, adding the
// count so far of each one to that each card it wins. We need only go
// through each card once since the cards only win cards below them.
func (l *Lotto) PartTwo() Result {
	counts := make(map[int]int)
	total := 0
	for _, card := range l.cards {
		counts[card.number] = 1
	}
	for _, card := range l.cards {
		for j := card.number + 1; j <= card.number+card.Score(); j++ {
			counts[j] += counts[card.number]
		}
		total += counts[card.number]
	}
	return NumberResult{total}
}

var numberList = regexp.MustCompile(`[0-9 ]+`)

func (l *Lotto) Load(scanner *bufio.Scanner) {
	for row := 1; scanner.Scan(); row++ {
		numberSets := numberList.FindAllString(scanner.Text(), 3)
		game := &Card{number: row}
		l.cards = append(l.cards, game.Load(numberSets[1], numberSets[2]))
	}
}

type Card struct {
	number  int
	winners []int
	gottens []int
}

func (s *Card) Load(win string, got string) *Card {
	ws := strings.Split(win, " ")
	gs := strings.Split(got, " ")
	s.winners = Must(Alltoi(ws))
	s.gottens = Must(Alltoi(gs))
	return s
}
func (s *Card) Score() int {
	score := 0
	for _, w := range s.winners {
		if slices.Contains(s.gottens, w) {
			score += 1
		}
	}
	return score
}
