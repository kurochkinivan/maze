package entities

// Path is a sequence of Points representing a route through the maze.
type Path []Point

// NewPath constructs a Path from the given list of Cells by copying their Points.
func NewPath(cells []*Cell) Path {
	path := make([]Point, len(cells))
	for i, c := range cells {
		path[i] = c.Point
	}
	return path
}
