package domain

import "math/rand/v2"

type DFSGenerator struct{}

func NewDFSGenerator() *DFSGenerator {
	return &DFSGenerator{}
}

// Generate generates maze using DFS algorithm.
func (g *DFSGenerator) Generate(m *Maze) {
	visited := make(map[*Cell]bool, m.Height*m.Width)

	start := m.Cells[rand.IntN(m.Height)][rand.IntN(m.Width)]
	visited[start] = true

	stack := []*Cell{start}

	for len(stack) > 0 {
		current := stack[len(stack)-1]

		unvisited := m.unvisitedAdjacentCells(current, visited)
		if len(unvisited) == 0 {
			stack = stack[:len(stack)-1]
			continue
		}

		next := unvisited[rand.IntN(len(unvisited))]
		next.RemoveWall(current, next.Cell)

		visited[next.Cell] = true
		stack = append(stack, next.Cell)
	}
}
