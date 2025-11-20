package entities

// DirectionType represents the possible movement directions in the maze.
type DirectionType int

const (
	DirUp DirectionType = iota + 1
	DirDown
	DirLeft
	DirRight
)

// Direction defines a movement vector and its type.
type Direction struct {
	Type DirectionType
	DRow int
	DCol int
}

// Neighbor represents an adjacent cell and the direction to reach it.
type Neighbor struct {
	Cell      *Cell
	Direction Direction
}

// HasWall returns true if there's a wall in this direction.
func (d Direction) HasWall(cell *Cell) bool {
	switch d.Type {
	case DirUp:
		return cell.Walls.Top
	case DirDown:
		return cell.Walls.Bottom
	case DirLeft:
		return cell.Walls.Left
	case DirRight:
		return cell.Walls.Right
	default:
		return true
	}
}

// RemoveWall removes walls between two cells in this direction.
func (d Direction) RemoveWall(from, to *Cell) {
	switch d.Type {
	case DirUp:
		from.Walls.Top = false
		to.Walls.Bottom = false
	case DirDown:
		from.Walls.Bottom = false
		to.Walls.Top = false
	case DirLeft:
		from.Walls.Left = false
		to.Walls.Right = false
	case DirRight:
		from.Walls.Right = false
		to.Walls.Left = false
	}
}

var (
	DirectionUp    = Direction{Type: DirUp, DRow: -1, DCol: 0}
	DirectionDown  = Direction{Type: DirDown, DRow: 1, DCol: 0}
	DirectionLeft  = Direction{Type: DirLeft, DRow: 0, DCol: -1}
	DirectionRight = Direction{Type: DirRight, DRow: 0, DCol: 1}

	// AllDirections contains all possible movement directions.
	AllDirections = []Direction{DirectionUp, DirectionDown, DirectionLeft, DirectionRight}
)
