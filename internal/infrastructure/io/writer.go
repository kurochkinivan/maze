package io

import (
	"fmt"
	"io"
	"strings"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
)

// WriteMaze renders maze to io.Writer.
func WriteMaze(w io.Writer, m *maze.Maze) error {
	grid := renderGrid(m)
	return writeGrid(w, grid)
}

// WriteMazeWithSolution writes maze with the solution to io.Writer.
func WriteMazeWithSolution(w io.Writer, m *maze.Maze, start, end *entities.Cell, path *entities.Path) error {
	grid := renderGrid(m)
	grid = addPath(grid, start, end, path)
	return writeGrid(w, grid)
}

func renderGrid(m *maze.Maze) [][]byte {
	height := m.Rows*2 + 1
	width := m.Cols*2 + 1

	grid := make([][]byte, height)
	for row := range grid {
		grid[row] = make([]byte, width)
		for col := range grid[row] {
			grid[row][col] = '#'
		}
	}

	for i := 1; i < height-1; i += 2 {
		row := (i - 1) / 2
		for j := 1; j < width-1; j += 2 {
			col := (j - 1) / 2
			cell := m.Cell(row, col)

			grid[i][j] = ' '
			if !cell.Top {
				grid[i-1][j] = ' '
			}
			if !cell.Bottom {
				grid[i+1][j] = ' '
			}
			if !cell.Left {
				grid[i][j-1] = ' '
			}
			if !cell.Right {
				grid[i][j+1] = ' '
			}
		}
	}

	return grid
}

func addPath(grid [][]byte, start, end *entities.Cell, path *entities.Path) [][]byte {
	for i := 1; i < len(path.Cells); i++ {
		prev := path.Cells[i-1]
		curr := path.Cells[i]

		prevRow, prevCol := prev.Row*2+1, prev.Col*2+1
		curRow, curCol := curr.Row*2+1, curr.Col*2+1
		midRow := (prevRow + curRow) / 2
		midCol := (prevCol + curCol) / 2

		grid[prevRow][prevCol] = '.'
		grid[midRow][midCol] = '.'
		grid[curRow][curCol] = '.'
	}

	startRow, startCol := start.Row*2+1, start.Col*2+1
	endRow, endCol := end.Row*2+1, end.Col*2+1

	grid[startRow][startCol] = 'O'
	grid[endRow][endCol] = 'X'

	return grid
}

func writeGrid(w io.Writer, grid [][]byte) error {
	var sb strings.Builder
	for _, row := range grid {
		sb.Write(row)
		sb.WriteByte('\n')
	}

	_, err := fmt.Fprint(w, sb.String())
	return err
}
