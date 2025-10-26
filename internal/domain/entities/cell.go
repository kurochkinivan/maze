package entities

type Cell struct {
	Point
	Walls
}

// NewCell returns a cell with the specified row and column.
// All walls are set to true.
func NewCell(row int, col int) *Cell {
	return &Cell{
		Point: NewPoint(row, col),
		Walls: Walls{true, true, true, true},
	}
}

type Walls struct {
	Top, Right, Bottom, Left bool
}

type Point struct {
	row int
	col int
}

func NewPoint(row int, col int) Point {
	return Point{
		row: row,
		col: col,
	}
}

func (p Point) Row() int {
	return p.row
}

func (p Point) Col() int {
	return p.col
}
