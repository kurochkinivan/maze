package dfs

import (
	"math/rand/v2"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
)

type DFSGenerator struct{}

func New() *DFSGenerator {
	return &DFSGenerator{}
}

// Generate generates maze using DFS algorithm.
func (g *DFSGenerator) Generate(m *maze.Maze) {
	visited := make(map[*entities.Cell]bool, m.Size())

	start := m.RandomCell()
	visited[start] = true

	stack := []*entities.Cell{start}

	for len(stack) > 0 {
		current := stack[len(stack)-1]

		unvisited := m.UnvisitedNeighbors(current, visited)
		
		if len(unvisited) == 0 {
			stack = stack[:len(stack)-1]
			continue
		}

		neighbor := unvisited[rand.IntN(len(unvisited))]
		next := neighbor.Cell
		neighbor.Direction.RemoveWall(current, next)

		visited[next] = true
		stack = append(stack, next)
	}
}
