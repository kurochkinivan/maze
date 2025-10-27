package maze

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"
)

type DirectionType int

const (
	DirUp DirectionType = iota + 1
	DirDown
	DirLeft
	DirRight
)

type Direction struct {
	Type DirectionType
	DRow int
	DCol int
}

type Neighbor struct {
	Cell      *entities.Cell
	Direction Direction
}

// HasWall returns true if there's a wall in this direction.
func (d Direction) HasWall(cell *entities.Cell) bool {
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
func (d Direction) RemoveWall(from, to *entities.Cell) {
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

	allDirections = []Direction{DirectionUp, DirectionDown, DirectionLeft, DirectionRight}
)
