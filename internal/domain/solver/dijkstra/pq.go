package dijkstra

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"
)

// weightedCell wraps a maze cell with its cost for Dijkstra's algorithm.
type weightedCell struct {
	cell *entities.Cell
	cost int
}

// priorityQueue implements a simple priority queue for weightedCell.
// It satisfies heap.Interface.
type priorityQueue struct {
	items []*weightedCell
}

// Len returns the number of items in the queue.
func (pq *priorityQueue) Len() int {
	return len(pq.items)
}

// Less returns true if the item at index i has lower cost than item at index j.
func (pq *priorityQueue) Less(i, j int) bool {
	return pq.items[i].cost < pq.items[j].cost
}

// Swap swaps the items at indices i and j.
func (pq *priorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

// Push adds a new weightedCell to the end of the queue.
func (pq *priorityQueue) Push(x any) {
	pq.items = append(pq.items, x.(*weightedCell))
}

// Pop removes and returns the last weightedCell from the queue.
func (pq *priorityQueue) Pop() any {
	last := pq.Len() - 1
	item := pq.items[last]
	pq.items = pq.items[:last]
	return item
}
