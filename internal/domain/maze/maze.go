package maze

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"
)

const numberOfWalls = 4

// Maze represents a grid of cells forming the maze structure.
type Maze struct {
	cells [][]*entities.Cell
	rows  int
	cols  int
}

// New creates a maze filled with cells, each initially surrounded by walls.
func New(width int, height int) *Maze {
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

// ReachableNeighbors returns all neighboring cells that can be reached from the given cell (i.e. no wall separates them).
func (m *Maze) ReachableNeighbors(cell *entities.Cell) []*entities.Cell {
	neighbors := make([]*entities.Cell, 0, numberOfWalls)

	for _, dir := range entities.AllDirections {
		if dir.HasWall(cell) {
			continue
		}

		newRow, newCol := cell.Row()+dir.DRow, cell.Col()+dir.DCol
		neighbor := entities.NewPoint(newRow, newCol)
		if !m.IsValid(neighbor) {
			continue
		}

		neighbors = append(neighbors, m.Cell(neighbor.Row(), neighbor.Col()))
	}

	return neighbors
}

// UnvisitedNeighbors returns neighbors of a cell that have not been visited.
func (m *Maze) UnvisitedNeighbors(cell *entities.Cell, visited map[*entities.Cell]bool) []*entities.Neighbor {
	return m.filteredNeighbors(cell, func(c *entities.Cell) bool {
		return !visited[c]
	})
}

// VisitedNeighbors returns neighbors of a cell that have already been visited.
func (m *Maze) VisitedNeighbors(cell *entities.Cell, visited map[*entities.Cell]bool) []*entities.Neighbor {
	return m.filteredNeighbors(cell, func(c *entities.Cell) bool {
		return visited[c]
	})
}

// filteredNeighbors returns all neighboring cells that match a given filter condition.
func (m *Maze) filteredNeighbors(cell *entities.Cell, filter func(c *entities.Cell) bool) []*entities.Neighbor {
	neighbors := m.neighbors(cell)
	filtered := make([]*entities.Neighbor, 0, len(neighbors))

	for _, n := range neighbors {
		if filter(n.Cell) {
			filtered = append(filtered, n)
		}
	}
	return filtered
}

// neighbors returns all valid neighboring cells (i.e. cells that are in bounds of the maze)
// of the given cell, regardless of walls.
func (m *Maze) neighbors(cell *entities.Cell) []*entities.Neighbor {
	neighbors := make([]*entities.Neighbor, 0, numberOfWalls)

	for _, dir := range entities.AllDirections {
		newRow, newCol := cell.Row()+dir.DRow, cell.Col()+dir.DCol
		neighbor := entities.NewPoint(newRow, newCol)
		if !m.IsValid(neighbor) {
			continue
		}

		neighbors = append(neighbors, &entities.Neighbor{
			Cell:      m.Cell(neighbor.Row(), neighbor.Col()),
			Direction: dir,
		})
	}

	return neighbors
}

// Cell returns the cell located at the given row and column.
func (m *Maze) Cell(row, col int) *entities.Cell {
	return m.cells[row][col]
}

// Size returns the total number of cells in the maze.
func (m *Maze) Size() int {
	return m.Rows() * m.Cols()
}

// IsValid checks if the given point lies within the maze boundaries.
func (m *Maze) IsValid(p entities.Point) bool {
	return 0 <= p.Row() && p.Row() < m.Rows() && 0 <= p.Col() && p.Col() < m.Cols()
}

// Rows returns the number of rows in the maze.
func (m *Maze) Rows() int {
	return m.rows
}

// Cols returns the number of columns in the maze.
func (m *Maze) Cols() int {
	return m.cols
}
