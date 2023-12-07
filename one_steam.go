package aoc2023

import (
	"bufio"
	"io"
)

func NewSteam() *Steam {
	return &Steam{value: -1}
}

type Steam struct {
	tree  [26]*Steam
	value int
}

func (st *Steam) Add(word string, value int) {
	root := st
	for i := 0; i < len(word); i++ {
		branch := root.tree[word[i]-'a']
		if branch == nil {
			branch = NewSteam()
			root.tree[word[i]-'a'] = branch
		}
		root = branch
	}
	root.value = value
}

func (st *Steam) Value(s string) int {
	root := st
	for i := 0; i < len(s); i++ {
		root = root.tree[s[i]]
		if root == nil {
			break
		}
		if root.value > 0 {
			return root.value
		}
	}
	return -1
}

func (st *Steam) Scan(s string) int {
	for i := 0; i < len(s); i++ {
		value := st.Value(s[i:])
		if value > 0 {
			return value
		}
	}
	return -1
}

func DayOnePartTwoSteam(r io.Reader) int {
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

var ForeSteam = NewSteam()
var BackSteam = NewSteam()

func init() {
	for word, value := range words {
		// don't need digits here
		if len(word) == 1 {
			continue
		}
		ForeSteam.Add(word, value)
		BackSteam.Add(reverse(word), value)
	}
}
