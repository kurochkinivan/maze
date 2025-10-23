package maze

import (
	"fmt"
	"math/rand/v2"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"
)

type Maze struct {
	cells [][]*entities.Cell
	Rows  int
	Cols  int
}

type Generator interface {
	Generate(m *Maze)
}

func New(width int, height int, generator Generator) *Maze {
	m := NewEmpty(width, height)
	m.Generate(generator)
	return m
}

func NewEmpty(width int, height int) *Maze {
	cells := make([][]*entities.Cell, height)
	for row := range height {
		cells[row] = make([]*entities.Cell, width)
		for col := range width {
			cells[row][col] = &entities.Cell{
				Point: entities.Point{Row: row, Col: col},
				Walls: entities.Walls{Top: true, Right: true, Bottom: true, Left: true},
			}
		}
	}

	return &Maze{
		Rows:  height,
		Cols:  width,
		cells: cells,
	}
}

func (m *Maze) Generate(generator Generator) {
	generator.Generate(m)
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

func (m *Maze) Cell(row, col int) *entities.Cell {
	return m.cells[row][col]
}

func (m *Maze) RandomCell() *entities.Cell {
	return m.cells[rand.IntN(m.Rows)][rand.IntN(m.Cols)]
}

func (m *Maze) Size() int {
	return m.Rows * m.Cols
}

func (m *Maze) IsValid(p entities.Point) bool {
	return 0 <= p.Row && p.Row < m.Rows && 0 <= p.Col && p.Col < m.Cols
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
	neighbors := make([]*Neighbor, 0, 4)

	for _, dir := range allDirections {
		neighbor := entities.Point{Row: cell.Row + dir.DRow, Col: cell.Col + dir.DCol}
		if !m.IsValid(neighbor) {
			continue
		}

		neighbors = append(neighbors, &Neighbor{
			Cell:      m.Cell(neighbor.Row, neighbor.Col),
			Direction: dir,
		})
	}

	return neighbors
}

// Display выводит лабиринт в консоль в текстовом формате.
// Использует символы '#' для стен и пробелы для проходов между клетками.
func (m *Maze) Display() {
	// Создаем сетку: каждая клетка занимает 2x2 символа + границы
	height := m.Rows*2 + 1
	width := m.Cols*2 + 1
	grid := make([][]byte, height)

	for i := range grid {
		grid[i] = make([]byte, width)
		for j := range grid[i] {
			grid[i][j] = '#'
		}
	}

	// Заполняем внутренности клеток и проходы
	for row := 0; row < m.Rows; row++ {
		for col := 0; col < m.Cols; col++ {
			cell := m.Cell(row, col)

			// Центр клетки всегда пустой
			gridRow := row*2 + 1
			gridCol := col*2 + 1
			grid[gridRow][gridCol] = ' '

			// Убираем стены если их нет
			if !cell.Top {
				grid[gridRow-1][gridCol] = ' '
			}
			if !cell.Bottom {
				grid[gridRow+1][gridCol] = ' '
			}
			if !cell.Left {
				grid[gridRow][gridCol-1] = ' '
			}
			if !cell.Right {
				grid[gridRow][gridCol+1] = ' '
			}
		}
	}

	// Выводим результат
	for _, row := range grid {
		fmt.Println(string(row))
	}
}
