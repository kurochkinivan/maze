package entities

type Path struct {
	Cells []*Cell
}

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

func (p *Path) ReversePath() *Path {
	for i, j := 0, len(p.Cells)-1; i < j; i, j = i+1, j-1 {
		p.Cells[i], p.Cells[j] = p.Cells[j], p.Cells[i]
	}
	return p
}
