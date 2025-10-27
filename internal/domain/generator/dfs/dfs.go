package dfs

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/generator"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
)

type Generator struct {
	generator.BaseGenerator
}

func New(opts ...generator.Option) *Generator {
	return &Generator{
		BaseGenerator: generator.NewBaseGenerator(opts...),
	}
}

// Generate generates maze using DFS algorithm.
func (g *Generator) Generate(m *maze.Maze) {
	if m.Size() <= 1 {
		return
	}

	visited := make(map[*entities.Cell]bool, m.Size())

	start := m.Cell(g.Rand().IntN(m.Rows()), g.Rand().IntN(m.Cols()))
	visited[start] = true

	stack := []*entities.Cell{start}

	for len(stack) > 0 {
		current := stack[len(stack)-1]

		unvisited := m.UnvisitedNeighbors(current, visited)

		if len(unvisited) == 0 {
			stack = stack[:len(stack)-1]
			continue
		}

		neighbor := unvisited[g.Rand().IntN(len(unvisited))]
		next := neighbor.Cell
		neighbor.Direction.RemoveWall(current, next)

		visited[next] = true
		stack = append(stack, next)
	}
}
