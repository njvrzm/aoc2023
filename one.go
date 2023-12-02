package aoc2023

import (
	"bufio"
	"io"
	"strings"
)

const digits = "0123456789"

func digit(b uint8) int {
	return int(b - '0')
}
func DigitAt(s string, f func(string, string) int) int {
	return digit(s[f(s, digits)])
}
func DayOnePartOne(r io.Reader) int {
	br := bufio.NewScanner(r)
	sum := 0
	for br.Scan() {
		line := br.Text()
		sum += 10 * DigitAt(line, strings.IndexAny)
		sum += DigitAt(line, strings.LastIndexAny)
	}
	return sum
}

var words = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func DayOnePartTwo(r io.Reader) int {
	br := bufio.NewScanner(r)
	sum := 0
	for br.Scan() {
		line := br.Text()
		sum += 10 * LeftmostNumber(line)
		sum += RightmostNumber(line)
	}
	return sum
}

func LeftmostNumber(s string) int {
	left := len(s)
	value := 0
	for w, i := range words {
		index := strings.Index(s, w)
		if index >= 0 && index < left {
			value = i
			left = index
		}
		if left == 0 {
			break
		}
	}
	return value
}

func RightmostNumber(s string) int {
	right := -1
	value := 0
	top := len(s) - 1
	for w, i := range words {
		index := strings.LastIndex(s, w)
		if index > right {
			value = i
			right = index
		}
		if right == top {
			break
		}
	}
	return value
}
