package io

import (
	"bufio"
	"fmt"
	"io"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
)

// ReadMaze scans maze from io.Reader.
func ReadMaze(r io.Reader) (*maze.Maze, error) {
	sc := bufio.NewScanner(r)
	lines := make([]string, 0, 4)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	if err := sc.Err(); err != nil {
		return nil, fmt.Errorf("scan error: %w", err)
	}

	return parseMaze(lines)
}

func parseMaze(lines []string) (*maze.Maze, error) {
	if len(lines) < 3 || len(lines[0]) < 3 {
		return nil, fmt.Errorf("invalid maze: must contain at least 3 lines and cols")
	}

	if len(lines)%2 != 1 || len(lines[0])%2 != 1 {
		return nil, fmt.Errorf("invalid maze: number of lines and cols must be odd")
	}

	height := (len(lines) - 1) / 2
	width := (len(lines[0]) - 1) / 2
	m := maze.New(width, height)

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
