package prim

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/generator"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/pkg/bag"
)

// Generator creates mazes using the Prim algorithm.
type Generator struct {
	generator.RandomGenerator
}

// New creates a new Prim-based maze generator with optional settings.
func New() *Generator {
	return &Generator{
		RandomGenerator: generator.NewRandomGenerator(),
	}
}

// Generate builds a maze using Prim's algorithm.
func (g *Generator) Generate(m *maze.Maze) {
	if m.Size() <= 1 {
		return
	}

	visited := make(map[*entities.Cell]bool, m.Size())
	frontier := bag.New[*entities.Cell](0)

	start := g.RandomCell(m)
	visited[start] = true

	for _, cell := range m.UnvisitedNeighbors(start, visited) {
		frontier.Add(cell.Cell)
	}

	for frontier.Len() > 0 {
		frontierCell := frontier.RandomItemAndDelete()

		visitedCells := m.VisitedNeighbors(frontierCell, visited)
		neighbor := visitedCells[g.IntN(len(visitedCells))]

		neighbor.Direction.RemoveWall(frontierCell, neighbor.Cell)

		visited[frontierCell] = true

		for _, cell := range m.UnvisitedNeighbors(frontierCell, visited) {
			frontier.Add(cell.Cell)
		}
	}
}
