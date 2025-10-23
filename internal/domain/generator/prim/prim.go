package prim

import (
	"math/rand/v2"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/lib/bag"
)

type PrimGenerator struct{}

func New() *PrimGenerator {
	return &PrimGenerator{}
}

// Generate generates maze using Prim algorithm.
func (g *PrimGenerator) Generate(m *maze.Maze) {
	visited := make(map[*entities.Cell]bool, m.Size())
	frontier := bag.New[*entities.Cell](4)

	start := m.RandomCell()
	visited[start] = true

	for _, cell := range m.UnvisitedNeighbors(start, visited) {
		frontier.Add(cell.Cell)
	}

	for frontier.Len() > 0 {
		frontierCell := frontier.RandomItemAndDelete()

		visitedCells := m.VisitedNeighbors(frontierCell, visited)
		neighbor := visitedCells[rand.IntN(len(visitedCells))]

		neighbor.Direction.RemoveWall(frontierCell, neighbor.Cell)

		visited[frontierCell] = true

		for _, cell := range m.UnvisitedNeighbors(frontierCell, visited) {
			frontier.Add(cell.Cell)
		}
	}
}
