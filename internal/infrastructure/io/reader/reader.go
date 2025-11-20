package maze_reader

import (
	"bufio"
	"fmt"
	"io"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
)

// ReadMaze scans maze from io.Reader.
func ReadMaze(r io.Reader) (*maze.Maze, error) {
	sc := bufio.NewScanner(r)

	var lines []string
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	if err := sc.Err(); err != nil {
		return nil, fmt.Errorf("scan error: %w", err)
	}

	return parseMaze(lines)
}

// parseMaze parses maze from given lines according to the certain rules.
func parseMaze(lines []string) (*maze.Maze, error) {
	err := validateMaze(lines)
	if err != nil {
		return nil, fmt.Errorf("invalid maze: %w", err)
	}

	height := (len(lines) - 1) / 2
	width := (len(lines[0]) - 1) / 2
	m := maze.New(width, height)

	// Cells are located on the positions with odd indexes (e.g. (1,1), (1,3))
	// Iterate over odd positions and check if they walls surrounding them
	for i := 1; i < len(lines)-1; i += 2 {
		row := (i - 1) / 2
		for j := 1; j < len(lines[i])-1; j += 2 {
			col := (j - 1) / 2
			cell := m.Cell(row, col)

			cell.Left = lines[i][j-1] != ' '
			cell.Right = lines[i][j+1] != ' '
			cell.Top = lines[i-1][j] != ' '
			cell.Bottom = lines[i+1][j] != ' '
		}
	}

	return m, nil
}

// validateMaze checks that the maze has correct dimensions, consistent row lengths,
// odd number of rows and columns, and fully enclosed outer walls.
func validateMaze(lines []string) error {
	if len(lines) < 3 || len(lines[0]) < 3 {
		return fmt.Errorf("invalid maze: must contain at least 3 lines and cols")
	}

	// check all rows are of equal length
	expectedLen := len(lines[0])
	for i, line := range lines {
		if len(line) != expectedLen {
			return fmt.Errorf("inconsistent row length at line %d", i)
		}
	}

	if len(lines)%2 != 1 || expectedLen%2 != 1 {
		return fmt.Errorf("number of lines and cols must be odd")
	}

	// check that all cells on the left and right are walls
	for row := range lines {
		if lines[row][0] != '#' || lines[row][expectedLen-1] != '#' {
			return fmt.Errorf("missing outer wall on row=%d", row)
		}
	}

	// check that all cells on the top and bottom are walls
	for col := range expectedLen {
		if lines[0][col] != '#' || lines[len(lines)-1][col] != '#' {
			return fmt.Errorf("missing outer wall on col=%d", col)
		}
	}

	for row := range lines {
		for col := range lines[row] {
			// check that all wall cells at even coordinates are actually walls
			if row%2 == 0 && col%2 == 0 && lines[row][col] != '#' {
				return fmt.Errorf("passages must be on odd positions, row=%d, col=%d", row, col)
			}
		}
	}

	return nil
}
