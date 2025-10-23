package maze

import "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"

type Neighbor struct {
	Cell      *entities.Cell
	Direction Direction
}

type Direction struct {
	DRow, DCol int
	removeWall func(from, to *entities.Cell)
}

func (d Direction) RemoveWall(from, to *entities.Cell) {
	d.removeWall(from, to)
}

var (
	DirectionDown = Direction{
		DRow: 1, DCol: 0,
		removeWall: func(a, b *entities.Cell) { a.Walls.Bottom, b.Walls.Top = false, false },
	}
	DirectionUp = Direction{
		DRow: -1, DCol: 0,
		removeWall: func(a, b *entities.Cell) { a.Walls.Top, b.Walls.Bottom = false, false },
	}
	DirectionRight = Direction{
		DRow: 0, DCol: 1,
		removeWall: func(a, b *entities.Cell) { a.Walls.Right, b.Walls.Left = false, false },
	}
	DirectionLeft = Direction{
		DRow: 0, DCol: -1,
		removeWall: func(a, b *entities.Cell) { a.Walls.Left, b.Walls.Right = false, false },
	}

	allDirections = []Direction{DirectionDown, DirectionUp, DirectionRight, DirectionLeft}
)
