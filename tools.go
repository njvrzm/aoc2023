package aoc2023

import (
	"regexp"
	"strconv"
)

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func AllNumbers(line string) []int {
	ns, _ := Alltoi(regexp.MustCompile(`[0-9]+`).FindAllString(line, -1))
	return ns
}
func Alltoi(nss []string) ([]int, error) {
	// may be over-allocated
	out := make([]int, len(nss))
	for i, ns := range nss {
		if ns != "" {
			n, err := strconv.Atoi(ns)
			if err != nil {
				return nil, err
			}
			out[i] = n
		}
	}
	return out, nil
}
