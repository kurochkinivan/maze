package domain

import (
	"math/rand/v2"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/lib/bag"
)

type PrimGenerator struct{}

func NewPrimGenerator() *PrimGenerator {
	return &PrimGenerator{}
}

// Generate generates maze using Prim algorithm.
func (g *PrimGenerator) Generate(m *Maze) {
	visited := make(map[*Cell]bool, m.Height*m.Width)
	frontier := bag.New[*Cell](4)

	start := m.Cells[rand.IntN(m.Height)][rand.IntN(m.Width)]
	visited[start] = true

	for _, cell := range m.unvisitedAdjacentCells(start, visited) {
		frontier.Add(cell.Cell)
	}

	for frontier.Len() > 0 {
		frontierCell := frontier.RandomItemAndDelete()

		inCells := m.visitedAdjacentCells(frontierCell, visited)
		move := inCells[rand.IntN(len(inCells))]

		move.RemoveWall(frontierCell, move.Cell)

		visited[frontierCell] = true

		for _, cell := range m.unvisitedAdjacentCells(start, visited) {
			frontier.Add(cell.Cell)
		}
	}
}
