package aoc2023

import "math"

type Place struct {
	X int
	Y int
}

var Nowhere = Place{math.MaxInt, math.MaxInt}
var Up = Place{0, -1}
var Right = Place{1, 0}
var Down = Place{0, 1}
var Left = Place{-1, 0}

func (p Place) To(o Place) Place {
	return Place{o.X - p.X, o.Y - p.Y}
}
func (p Place) Plus(o Place) Place {
	return Place{p.X + o.X, p.Y + o.Y}
}
func (p Place) Up() Place {
	return p.Plus(Place{0, -1})
}
func (p Place) Down() Place {
	return p.Plus(Place{0, 1})
}
func (p Place) Left() Place {
	return p.Plus(Place{-1, 0})
}
func (p Place) Right() Place {
	return p.Plus(Place{1, 0})
}

func (p Place) Neighbors() chan Place {
	return p.RectangleNeighbors(1, 1)
}

func (p Place) RectangleNeighbors(width int, height int) chan Place {
	out := make(chan Place)
	go func() {
		defer close(out)
		for row := p.Y - 1; row <= p.Y+height; row++ {
			for col := p.X - 1; col <= p.X+width; col++ {
				if row > p.Y-1 && row < p.Y+height && col == p.X {
					col = p.X + width
				}
				out <- Place{X: col, Y: row}
			}
		}
	}()
	return out
}
