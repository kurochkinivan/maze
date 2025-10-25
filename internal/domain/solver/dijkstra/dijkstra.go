package dijkstra

import (
	"container/heap"
	"math"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

func (s *Solver) Solve(m *maze.Maze, start, end *entities.Cell) (*entities.Path, bool) {
	visited := make(map[*entities.Cell]bool)
	previous := make(map[*entities.Cell]*entities.Cell)
	costs := make(map[*entities.Cell]int, m.Size())

	for row := range m.Rows {
		for col := range m.Cols {
			cell := m.Cell(row, col)
			costs[cell] = math.MaxInt32
		}
	}
	costs[start] = 0

	pq := &priorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &weightedCell{cell: start, cost: 0})

	for pq.Len() != 0 {
		current := heap.Pop(pq).(*weightedCell)

		if current.cell == end {
			return entities.BuildPath(previous, end), true
		}

		if visited[current.cell] {
			continue
		}

		visited[current.cell] = true

		for _, reachable := range m.ReachableNeighbors(current.cell) {
			newCost := current.cost + 1
			oldCost := costs[reachable]

			if newCost < oldCost {
				costs[reachable] = newCost
				previous[reachable] = current.cell
				heap.Push(pq, &weightedCell{
					cell: reachable,
					cost: newCost,
				})
			}

		}
	}

	return nil, false
}
