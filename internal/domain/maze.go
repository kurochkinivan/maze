package domain

import (
	"fmt"
)

type Maze struct {
	Cells  [][]*Cell
	Width  int
	Height int
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
	cells := make([][]*Cell, height)
	for row := range height {
		cells[row] = make([]*Cell, width)
		for col := range width {
			cells[row][col] = &Cell{
				Point: Point{X: row, Y: col},
				Walls: Walls{true, true, true, true},
			}
		}
	}

	return &Maze{
		Height: height,
		Width:  width,
		Cells:  cells,
	}
}

func (m *Maze) Generate(generator Generator) {
	generator.Generate(m)
}

func (m *Maze) visitedAdjacentCells(start *Cell, visited map[*Cell]bool) []*Move {
	visitedCells := make([]*Move, 0, 4)

	for _, dir := range directions {
		neighbor := Point{X: start.X + dir.DX, Y: start.Y + dir.DY}
		if !m.IsValid(neighbor) {
			continue
		}

		cell := m.Cells[neighbor.X][neighbor.Y]
		if visited[cell] {
			visitedCells = append(visitedCells, &Move{
				Cell:       cell,
				RemoveWall: dir.RemoveWall,
			})
		}
	}

	return visitedCells
}

func (m *Maze) unvisitedAdjacentCells(start *Cell, visited map[*Cell]bool) []*Move {
	unvisited := make([]*Move, 0, 4)

	for _, dir := range directions {
		neighbor := Point{X: start.X + dir.DX, Y: start.Y + dir.DY}
		if !m.IsValid(neighbor) {
			continue
		}

		cell := m.Cells[neighbor.X][neighbor.Y]
		if !visited[cell] {
			unvisited = append(unvisited, &Move{
				Cell:       cell,
				RemoveWall: dir.RemoveWall,
			})
		}
	}

	return unvisited
}

func (m *Maze) IsValid(p Point) bool {
	return 0 <= p.X && p.X < m.Height && 0 <= p.Y && p.Y < m.Width
}

// Display выводит лабиринт в консоль в текстовом формате.
// Использует символы '#' для стен и пробелы для проходов между клетками.
func (m *Maze) Display() {
	// Создаем сетку: каждая клетка занимает 2x2 символа + границы
	height := m.Height*2 + 1
	width := m.Width*2 + 1
	grid := make([][]byte, height)

	for i := range grid {
		grid[i] = make([]byte, width)
		for j := range grid[i] {
			grid[i][j] = '#'
		}
	}

	// Заполняем внутренности клеток и проходы
	for row := 0; row < m.Height; row++ {
		for col := 0; col < m.Width; col++ {
			cell := m.Cells[row][col]

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
