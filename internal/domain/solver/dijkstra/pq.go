package dijkstra

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"
)

type weightedCell struct {
	cell *entities.Cell
	cost int
}

type priorityQueue struct {
	items []*weightedCell
}

func (pq *priorityQueue) Len() int {
	return len(pq.items)
}

func (pq *priorityQueue) Less(i, j int) bool {
	return pq.items[i].cost < pq.items[j].cost
}

func (pq *priorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

func (pq *priorityQueue) Push(x any) {
	pq.items = append(pq.items, x.(*weightedCell))
}

func (pq *priorityQueue) Pop() any {
	last := pq.Len() - 1

	item := pq.items[last]
	pq.items = pq.items[:last]

	return item
}
