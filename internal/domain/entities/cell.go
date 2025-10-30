package entities

// Cell represents a single cell in the maze.
// It contains its position (Point) and the state of its four surrounding walls.
type Cell struct {
	Point
	Walls
}

// NewCell creates and returns a new Cell with the given row and column coordinates.
// By default, all walls (Top, Right, Bottom, Left) are initialized as true.
func NewCell(row int, col int) *Cell {
	return &Cell{
		Point: NewPoint(row, col),
		Walls: Walls{true, true, true, true},
	}
}

// Walls defines the presence or absence of walls around a cell.
// A value of true indicates that the wall exists.
type Walls struct {
	Top, Right, Bottom, Left bool
}

// Point represents the coordinates of a cell in the maze grid.
type Point struct {
	row int
	col int
}

// NewPoint returns a new Point with the specified row and column value
func NewPoint(row int, col int) Point {
	return Point{
		row: row,
		col: col,
	}
}

// Row returns the row index of the point.
func (p Point) Row() int {
	return p.row
}

// Col returns the column index of the point.
func (p Point) Col() int {
	return p.col
}
