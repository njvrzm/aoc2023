package aoc2023

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

type DiceGame struct {
	Rolls []*Roll
	Id    int

	line string
}

func ReadGame(line string) *DiceGame {
	game := &DiceGame{line: line}
	parts := strings.SplitN(line, ": ", 2)
	game.Id = Must(strconv.Atoi(parts[0][5:])) // skip "Game "
	rolls := strings.Split(parts[1], "; ")
	for _, roll := range rolls {
		r := &Roll{}
		game.Rolls = append(game.Rolls, r)
		colorCounts := strings.Split(roll, ", ")
		for _, colorCount := range colorCounts {
			n_c := strings.Split(colorCount, " ")
			color := n_c[1]
			count := Must(strconv.Atoi(n_c[0]))
			switch color {
			case "red":
				r.Red = count
			case "blue":
				r.Blue = count
			case "green":
				r.Green = count
			default:
				log.Fatalf("Unknown color: %s", color)
			}
		}
	}
	return game
}
func (dg *DiceGame) IsValid(r, g, b int) bool {
	for _, roll := range dg.Rolls {
		if roll.Red > r || roll.Green > g || roll.Blue > b {
			return false
		}
	}
	return true
}

type Roll struct {
	Red   int
	Green int
	Blue  int
}

func DayTwoPartOne(r io.Reader) int {
	b := bufio.NewScanner(r)
	good := 0
	for b.Scan() {
		game := ReadGame(b.Text())
		if game.IsValid(12, 13, 14) {
			good += game.Id
		}
	}
	return good
}

func DayTwoPartTwo(r io.Reader) int {
	b := bufio.NewScanner(r)
	powers := 0
	for b.Scan() {
		game := ReadGame(b.Text())
		minimum := &Roll{}
		for _, roll := range game.Rolls {
			if roll.Red > minimum.Red {
				minimum.Red = roll.Red
			}
			if roll.Green > minimum.Green {
				minimum.Green = roll.Green
			}
			if roll.Blue > minimum.Blue {
				minimum.Blue = roll.Blue
			}
		}
		powers += minimum.Red * minimum.Green * minimum.Blue
	}
	return powers
}
