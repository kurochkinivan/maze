package solver

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
)

type Solver interface {
	Solve(m *maze.Maze, start, end *entities.Cell) (entities.Path, bool)
}
