package maze

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"
)

const (
	numberOfWalls = 4
)

type Maze struct {
	cells [][]*entities.Cell
	rows  int
	cols  int
}

func New(width int, height int) *Maze {
	return newEmpty(width, height)
}

func newEmpty(width int, height int) *Maze {
	cells := make([][]*entities.Cell, height)
	for row := range height {
		cells[row] = make([]*entities.Cell, width)
		for col := range width {
			cells[row][col] = entities.NewCell(row, col)
		}
	}

	return &Maze{
		rows:  height,
		cols:  width,
		cells: cells,
	}
}

// ReachableNeighbors returns neighbors that has no walls with the cell.
func (m *Maze) ReachableNeighbors(cell *entities.Cell) []*entities.Cell {
	neighbors := make([]*entities.Cell, 0, numberOfWalls)

	for _, dir := range allDirections {
		if dir.HasWall(cell) {
			continue
		}

		neighbor := entities.NewPoint(cell.Row()+dir.DRow, cell.Col()+dir.DCol)

		if !m.IsValid(neighbor) {
			continue
		}

		neighbors = append(neighbors, m.Cell(neighbor.Row(), neighbor.Col()))
	}

	return neighbors
}

func (m *Maze) UnvisitedNeighbors(cell *entities.Cell, visited map[*entities.Cell]bool) []*Neighbor {
	return m.filteredNeighbors(cell, func(c *entities.Cell) bool {
		return !visited[c]
	})
}

func (m *Maze) VisitedNeighbors(cell *entities.Cell, visited map[*entities.Cell]bool) []*Neighbor {
	return m.filteredNeighbors(cell, func(c *entities.Cell) bool {
		return visited[c]
	})
}

func (m *Maze) filteredNeighbors(cell *entities.Cell, filter func(c *entities.Cell) bool) []*Neighbor {
	neighbors := m.neighbors(cell)
	filtered := make([]*Neighbor, 0, len(neighbors))

	for _, n := range neighbors {
		if filter(n.Cell) {
			filtered = append(filtered, n)
		}
	}

	return filtered
}

func (m *Maze) neighbors(cell *entities.Cell) []*Neighbor {
	neighbors := make([]*Neighbor, 0, numberOfWalls)

	for _, dir := range allDirections {
		neighbor := entities.NewPoint(cell.Row()+dir.DRow, cell.Col()+dir.DCol)
		if !m.IsValid(neighbor) {
			continue
		}

		neighbors = append(neighbors, &Neighbor{
			Cell:      m.Cell(neighbor.Row(), neighbor.Col()),
			Direction: dir,
		})
	}

	return neighbors
}

func (m *Maze) Cell(row, col int) *entities.Cell {
	return m.cells[row][col]
}

func (m *Maze) Size() int {
	return m.Rows() * m.Cols()
}

func (m *Maze) IsValid(p entities.Point) bool {
	return 0 <= p.Row() && p.Row() < m.Rows() && 0 <= p.Col() && p.Col() < m.Cols()
}

func (m *Maze) Rows() int {
	return m.rows
}

func (m *Maze) Cols() int {
	return m.cols
}
