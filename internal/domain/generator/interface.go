package generator

import "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"

type Generator interface {
	Generate(m *maze.Maze)
}
