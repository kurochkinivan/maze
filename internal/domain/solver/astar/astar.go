package astar

import (
	"math"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

func (s *Solver) Solve(m *maze.Maze, start, end *entities.Cell) (*entities.Path, bool) {
	reachable := make(map[*entities.Cell]struct{})
	visited := make(map[*entities.Cell]bool)

	costs := make(map[*entities.Cell]int)
	previous := make(map[*entities.Cell]*entities.Cell)

	costs[start] = 0
	previous[start] = nil
	reachable[start] = struct{}{}

	for len(reachable) > 0 {
		current := s.chooseNode(end, costs, reachable)

		if current == end {
			return entities.BuildPath(previous, end).ReversePath(), true
		}

		delete(reachable, current)
		visited[current] = true

		newReachable := m.ReachableNeighbors(current)
		for _, reach := range newReachable {
			if visited[reach] {
				continue
			}

			reachable[reach] = struct{}{}

			newCost := costs[current] + 1
			if oldCost, exists := costs[reach]; !exists || newCost < oldCost {
				costs[reach] = newCost
				previous[reach] = current
			}
		}
	}

	return nil, false
}

func (s *Solver) chooseNode(
	end *entities.Cell,
	costs map[*entities.Cell]int,
	reachable map[*entities.Cell]struct{},
) *entities.Cell {
	var best *entities.Cell
	minCost := math.MaxInt32

	for current := range reachable {
		costToDistance := s.manhattanDistance(current, end)
		totalCost := costToDistance + costs[current]

		if totalCost < minCost {
			minCost = totalCost
			best = current
		}
	}

	return best
}

func (s *Solver) manhattanDistance(c1, c2 *entities.Cell) int {
	return abs(c1.Row-c2.Row) + abs(c1.Col-c2.Col)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
