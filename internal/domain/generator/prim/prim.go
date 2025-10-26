package prim

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/generator"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/pkg/bag"
)

type Generator struct {
	generator.BaseGenerator
}

func New(opts ...generator.GeneratorOption) *Generator {
	return &Generator{
		BaseGenerator: generator.NewBaseGenerator(opts...),
	}
}

// Generate generates maze using Prim algorithm.
func (g *Generator) Generate(m *maze.Maze) {
	if m.Size() <= 1 {
		return
	}

	visited := make(map[*entities.Cell]bool, m.Size())
	frontier := bag.New[*entities.Cell](4)

	start := m.Cell(g.Rand().IntN(m.Rows), g.Rand().IntN(m.Cols))
	visited[start] = true

	for _, cell := range m.UnvisitedNeighbors(start, visited) {
		frontier.Add(cell.Cell)
	}

	for frontier.Len() > 0 {
		frontierCell := frontier.RandomItemAndDelete()

		visitedCells := m.VisitedNeighbors(frontierCell, visited)
		neighbor := visitedCells[g.Rand().IntN(len(visitedCells))]

		neighbor.Direction.RemoveWall(frontierCell, neighbor.Cell)

		visited[frontierCell] = true

		for _, cell := range m.UnvisitedNeighbors(frontierCell, visited) {
			frontier.Add(cell.Cell)
		}
	}
}
