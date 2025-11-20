package maze_writer

import (
	"fmt"
	"io"
	"strings"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
)

// WriteMaze writes the maze to an io.Writer in either ASCII or Unicode format.
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

// WriteMazeWithSolution writes the maze with the solution path to an io.Writer.
func WriteMazeWithSolution(w io.Writer, m *maze.Maze, path entities.Path, unicode bool) error {
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

// addPath overlays the solution path onto the maze grid using '.'; marks start 'O' and end 'X'.
func addPath(grid [][]rune, path entities.Path) {
	start := path[0]
	end := path[len(path)-1]

	for i := 1; i < len(path); i++ {
		prev := path[i-1]
		curr := path[i]

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
