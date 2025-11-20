package solver

import "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"

func BuildPath(previous map[*entities.Cell]*entities.Cell, end *entities.Cell) entities.Path {
	cells := make([]*entities.Cell, 0, len(previous))
	current := end

	for current != nil {
		cells = append(cells, current)
		current = previous[current]
	}

	reversePath(cells)

	return entities.NewPath(cells)
}

func reversePath(cells []*entities.Cell) {
	for i, j := 0, len(cells)-1; i < j; i, j = i+1, j-1 {
		cells[i], cells[j] = cells[j], cells[i]
	}
}
