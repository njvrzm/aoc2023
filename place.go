package aoc2023

type Place struct {
	X int
	Y int
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
