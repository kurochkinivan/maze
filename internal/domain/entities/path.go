package entities

// Path is a list of connected cells forming a maze route.
type Path struct {
	Cells []*Cell
}

// BuildPath creates a path from the end cell by tracing back through the 'previous' map.
// The returned path is reversed so it starts at the beginning.
func BuildPath(previous map[*Cell]*Cell, end *Cell) *Path {
	path := []*Cell{}
	current := end

	for current != nil {
		path = append(path, current)
		current = previous[current]
	}

	p := &Path{Cells: path}

	return p.ReversePath()
}

// ReversePath reverses the order of cells in the path.
func (p *Path) ReversePath() *Path {
	for i, j := 0, len(p.Cells)-1; i < j; i, j = i+1, j-1 {
		p.Cells[i], p.Cells[j] = p.Cells[j], p.Cells[i]
	}
	return p
}
