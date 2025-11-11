package maze_writer

import "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"

// renderGridASCII creates an ASCII representation of the maze using '#' for walls and ' ' for paths.
func renderGridASCII(m *maze.Maze) [][]rune {
	height := m.Rows()*2 + 1
	width := m.Cols()*2 + 1

	grid := make([][]rune, height)
	for row := range grid {
		grid[row] = make([]rune, width)
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
