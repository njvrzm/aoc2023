package aoc2023

import (
	"bufio"
	"cmp"
	"slices"
	"strconv"
	"strings"
)

type Game struct {
	hands []Hand
}

func (g *Game) PartOne() Result {
	slices.SortFunc(g.hands, CompareHands)
	total := 0
	for i, hand := range g.hands {
		total += hand.bid * (i + 1)
	}
	return NumberResult{total}
}

func (g *Game) PartTwo() Result {
	hands := make([]Hand, len(g.hands))
	for i, hand := range g.hands {
		for j := 0; j < 5; j++ {
			if hand.cards[j] == 11 {
				hand.cards[j] = 0
			}
		}
		hand.counts = hand.spectrum()
		hands[i] = hand
	}
	slices.SortFunc(hands, CompareHands)
	total := 0
	for i, hand := range hands {
		total += hand.bid * (i + 1)
	}
	return NumberResult{total}
}

func (g *Game) Load(scanner *bufio.Scanner) {
	for scanner.Scan() {
		g.hands = append(g.hands, ReadHand(scanner.Text()))
	}
	return
}

type Hand struct {
	cards  [5]int
	counts [5]int
	bid    int
}

func (h Hand) spectrum() [5]int {
	counts := make(map[int]int)
	jokers := 0
	for _, card := range h.cards {
		if card == 0 {
			jokers += 1
		} else {
			counts[card] += 1
		}
	}
	spect := make([]int, 5)
	i := 0
	for _, count := range counts {
		spect[i] = count
		i += 1
	}
	SortIntSliceDescending(spect)
	spect[0] += jokers
	return [5]int(spect)
}

var faces = map[uint8]int{
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

func ReadHand(line string) Hand {
	parts := strings.SplitN(line, " ", 2)
	hand, bid := parts[0], parts[1]
	cc := make(map[int]int)
	cards := make([]int, 5)
	for i := 0; i < 5; i++ {
		if '2' <= hand[i] && hand[i] <= '9' {
			cards[i] = Must(strconv.Atoi(string(hand[i])))
		} else {
			cards[i] = faces[hand[i]]
		}
		cc[cards[i]] += 1
	}
	out := Hand{cards: [5]int(cards), bid: Must(strconv.Atoi(bid))}
	out.counts = out.spectrum()
	return out
}
func SortIntSliceDescending(is []int) {
	slices.SortFunc(is, func(a int, b int) int { return -cmp.Compare(a, b) })
}

func CompareArrays(one, two [5]int) int {
	for i := 0; i < 5; i++ {
		if one[i] > two[i] {
			return 1
		} else if one[i] < two[i] {
			return -1
		}
	}
	return 0
}
func CompareHands(one Hand, two Hand) int {
	if ca := CompareArrays(one.spectrum(), two.spectrum()); ca != 0 {
		return ca
	}
	return CompareArrays(one.cards, two.cards)
}
