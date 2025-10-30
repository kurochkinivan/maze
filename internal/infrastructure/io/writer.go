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

// renderGrid generates a visual 2D byte representation of a maze.
// Walls are represented by '#' and open paths by spaces (' ').
// The function scales the maze dimensions by a factor of 2 + 1 to
// accommodate walls and passages between cells.
func renderGrid(m *maze.Maze) [][]byte {
	height := m.Rows()*2 + 1
	width := m.Cols()*2 + 1

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

// addPath overlays a path on top of a maze grid.
// The path is represented by '.' characters connecting cells from the start to the end.
// The start cell is marked with 'O', the end cell is marked with 'X'.
func addPath(grid [][]byte, start, end *entities.Cell, path *entities.Path) [][]byte {
	for i := 1; i < len(path.Cells); i++ {
		prev := path.Cells[i-1]
		curr := path.Cells[i]

		prevRow, prevCol := gridCoord(prev.Row()), gridCoord(prev.Col())
		curRow, curCol := gridCoord(curr.Row()), gridCoord(curr.Col())
		midRow := (prevRow + curRow) / 2
		midCol := (prevCol + curCol) / 2

		grid[prevRow][prevCol] = '.'
		grid[midRow][midCol] = '.'
		grid[curRow][curCol] = '.'
	}

	startRow, startCol := gridCoord(start.Row()), gridCoord(start.Col())
	endRow, endCol := gridCoord(end.Row()), gridCoord(end.Col())

	grid[startRow][startCol] = 'O'
	grid[endRow][endCol] = 'X'

	return grid
}

// gridCoord converts a maze cell index into a grid coordinate.
func gridCoord(n int) int {
	return n*2 + 1
}

// writeGrid writes the maze grid to the provided io.Writer.
func writeGrid(w io.Writer, grid [][]byte) error {
	var sb strings.Builder
	for _, row := range grid {
		sb.Write(row)
		sb.WriteByte('\n')
	}

	_, err := fmt.Fprint(w, sb.String())
	return err
}
