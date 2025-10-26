package io

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
)

func TestRenderGrid_HappyPath(t *testing.T) {
	expect := [][]byte{
		{'#', '#', '#', '#', '#'},
		{'#', ' ', '#', ' ', '#'},
		{'#', ' ', '#', ' ', '#'},
		{'#', ' ', ' ', ' ', '#'},
		{'#', '#', '#', '#', '#'},
	}

	m, _, _, _, _ := create2x2Maze()
	grid := renderGrid(m)
	assert.Equal(t, expect, grid)
}

func TestRenderGrid_OneCell(t *testing.T) {
	expect := [][]byte{
		{'#', '#', '#'},
		{'#', ' ', '#'},
		{'#', '#', '#'},
	}

	m := maze.New(1, 1)
	grid := renderGrid(m)
	assert.Equal(t, expect, grid)
}

func TestAddPath(t *testing.T) {
	expect := [][]byte{
		{'#', '#', '#', '#', '#'},
		{'#', 'O', '#', 'X', '#'},
		{'#', '.', '#', '.', '#'},
		{'#', '.', '.', '.', '#'},
		{'#', '#', '#', '#', '#'},
	}

	m, a, b, c, d := create2x2Maze()
	grid := renderGrid(m)

	path := &entities.Path{Cells: []*entities.Cell{a, c, d, b}}
	gridWithPath := addPath(grid, a, b, path)

	assert.Equal(t, expect, gridWithPath)
}

// helper: 2x2 maze with certain structure.
func create2x2Maze() (*maze.Maze, *entities.Cell, *entities.Cell, *entities.Cell, *entities.Cell) {
	/*
		   Maze 2x2:
		   [A] | [B]
			|  +  |
		   [C] - [D]
	*/
	m := maze.New(2, 2)
	cellA := m.Cell(0, 0)
	cellB := m.Cell(0, 1)
	cellC := m.Cell(1, 0)
	cellD := m.Cell(1, 1)

	maze.DirectionDown.RemoveWall(cellA, cellC)
	maze.DirectionRight.RemoveWall(cellC, cellD)
	maze.DirectionUp.RemoveWall(cellD, cellB)

	return m, cellA, cellB, cellC, cellD
}
