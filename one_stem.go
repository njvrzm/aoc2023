package aoc2023

import (
	"bufio"
	"io"
)

func NewStem() *Stem {
	return &Stem{tree: make(map[uint8]*Stem), value: -1}
}

type Stem struct {
	tree  map[uint8]*Stem
	value int
}

func (st *Stem) Add(word string, value int) {
	root := st
	for i := 0; i < len(word); i++ {
		branch, ok := root.tree[word[i]]
		if !ok {
			branch = NewStem()
			root.tree[word[i]] = branch
		} else if branch.value > 0 {
			return
		}
		root = branch
	}
	root.value = value
}

func (st *Stem) Value(s string) int {
	var ok bool
	root := st
	for i := 0; i < len(s); i++ {
		root, ok = root.tree[s[i]]
		if !ok {
			break
		}
		if root.value > 0 {
			return root.value
		}
	}
	return -1
}

func (st *Stem) Scan(s string) int {
	for i := 0; i < len(s); i++ {
		value := st.Value(s[i:])
		if value > 0 {
			return value
		}
	}
	return -1
}

func DayOnePartTwoStem(r io.Reader) int {
	sum := 0
	br := bufio.NewScanner(r)
	for br.Scan() {
		line := br.Text()
		first := ForeStem.Scan(line)
		last := BackStem.Scan(reverse(line))
		if first < 0 {
			continue
		}
		sum += first*10 + last
	}
	return sum
}

var ForeStem = NewStem()
var BackStem = NewStem()

func init() {
	for word, value := range words {
		ForeStem.Add(word, value)
		BackStem.Add(reverse(word), value)
	}
}

func reverse(ascii string) string {
	r := []rune(ascii)
	for i, j := 0, len(ascii)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
