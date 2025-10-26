package entities

type Cell struct {
	Point
	Walls
}

func NewCell(row int, col int) *Cell {
	return &Cell{
		Point: Point{
			Row: row,
			Col: col,
		},
		Walls: Walls{true, true, true, true},
	}
}

type Walls struct {
	Top, Right, Bottom, Left bool
}

type Point struct {
	Row int 
	Col int
}
