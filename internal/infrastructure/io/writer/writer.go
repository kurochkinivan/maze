package maze_writer

import (
	"fmt"
	"io"
	"strings"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
)

// WriteMaze renders maze to io.Writer.
func WriteMaze(w io.Writer, m *maze.Maze, unicode bool) error {
	gridASCII := renderGridASCII(m)

	var grid [][]rune
	if unicode {
		grid = renderGridUnicode(gridASCII)
	} else {
		grid = gridASCII
	}

	return writeGrid(w, grid)
}

// WriteMazeWithSolution writes maze with the solution to io.Writer.
func WriteMazeWithSolution(w io.Writer, m *maze.Maze, path *entities.Path, unicode bool) error {
	gridASCII := renderGridASCII(m)

	var grid [][]rune
	if unicode {
		grid = renderGridUnicode(gridASCII)
	} else {
		grid = gridASCII
	}

	addPath(grid, path)

	return writeGrid(w, grid)
}

// addPath overlays a path on top of a maze grid.
// The path is represented by '.' characters connecting cells from the start to the end.
// The start cell is marked with 'O', the end cell is marked with 'X'.
func addPath(grid [][]rune, path *entities.Path) {
	start := path.Cells[0]
	end := path.Cells[len(path.Cells)-1]

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
}

// gridCoord converts a maze cell index into a grid coordinate.
func gridCoord(n int) int {
	return n*2 + 1
}

// writeGrid writes the maze grid to the provided io.Writer.
func writeGrid(w io.Writer, grid [][]rune) error {
	var sb strings.Builder

	for _, row := range grid {
		for _, cell := range row {
			sb.WriteRune(rune(cell))
		}
		sb.WriteByte('\n')
	}

	_, err := fmt.Fprint(w, sb.String())
	return err
}
