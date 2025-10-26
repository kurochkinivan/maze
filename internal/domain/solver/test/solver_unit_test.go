package test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/solver/astar"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/solver/dijkstra"
)

const (
	Astar    = "astar"
	Dijkstra = "dijkstra"
)

type SolverTestSuite struct {
	suite.Suite
	Solvers map[string]Solver
}

type Solver interface {
	Solve(m *maze.Maze, start, end *entities.Cell) (*entities.Path, bool)
}

func TestSolverTestSuite(t *testing.T) {
	suite.Run(t, new(SolverTestSuite))
}

func (suite *SolverTestSuite) SetupTest() {
	suite.Solvers = map[string]Solver{
		Astar:    astar.New(),
		Dijkstra: dijkstra.New(),
	}
}

// TestSolve_SameCell checks the path to the same cell.
func (suite *SolverTestSuite) TestSolve_SameCell() {
	for _, solver := range suite.Solvers {
		m := maze.New(3, 3)
		cell := m.Cell(1, 1)

		path, ok := solver.Solve(m, cell, cell)

		suite.True(ok, "should find path to same cell")
		suite.NotNil(path)
		suite.Len(path.Cells, 1, "path to same cell should have length 1")
		suite.Equal(cell, path.Cells[0])
	}
}

// TestSolve_DirectPath checks direct path (no walls).
func (suite *SolverTestSuite) TestSolve_DirectPath() {
	/*
	   maze 3x1: [A] - [B] - [C]
	   No walls between the cells
	*/
	for _, solver := range suite.Solvers {
		m := maze.New(3, 1)

		cellA := m.Cell(0, 0)
		cellB := m.Cell(0, 1)
		cellC := m.Cell(0, 2)

		// Remove walls between cells
		maze.DirectionRight.RemoveWall(cellA, cellB)
		maze.DirectionRight.RemoveWall(cellB, cellC)

		path, ok := solver.Solve(m, cellA, cellC)

		suite.True(ok, "should find direct path")
		suite.NotNil(path)
		suite.Len(path.Cells, 3, "path should be A->B->C")
		suite.Equal(cellA, path.Cells[0])
		suite.Equal(cellB, path.Cells[1])
		suite.Equal(cellC, path.Cells[2])
	}
}

// TestSolve_NoPath checks case where no path can be found.
func (suite *SolverTestSuite) TestSolve_NoPath() {
	/*
	   Maze 2x2:
	   [A] | [B]
	   ----+----
	   [C] | [D]

	   All cells are isolated
	*/
	for _, solver := range suite.Solvers {
		m := maze.New(2, 2)

		start := m.Cell(0, 0)
		end := m.Cell(1, 1)

		path, ok := solver.Solve(m, start, end)

		suite.False(ok, "should not find path in isolated maze")
		suite.Nil(path)
	}
}

// TestSolve_OptimalPath checks that algorithm finds the oprimal way.
func (suite *SolverTestSuite) TestSolve_OptimalPath() {
	/*
	   Maze 3x3:
	   [A] - [B] - [C]
	    |           |
	   [D] - [E] - [F]

	   Optimal: A->D->E (length 3)
	   Alternative: A->B->C->F->E (length 5)
	*/
	for _, solver := range suite.Solvers {
		m := maze.New(3, 2)

		cellA := m.Cell(0, 0)
		cellB := m.Cell(0, 1)
		cellC := m.Cell(0, 2)
		cellD := m.Cell(1, 0)
		cellE := m.Cell(1, 1)
		cellF := m.Cell(1, 2)

		maze.DirectionRight.RemoveWall(cellA, cellB)
		maze.DirectionRight.RemoveWall(cellB, cellC)
		maze.DirectionRight.RemoveWall(cellD, cellE)
		maze.DirectionRight.RemoveWall(cellE, cellF)

		maze.DirectionDown.RemoveWall(cellA, cellD)
		maze.DirectionDown.RemoveWall(cellC, cellF)

		path, ok := solver.Solve(m, cellA, cellE)

		suite.True(ok, "should find path")
		suite.NotNil(path)
		suite.Equal(3, len(path.Cells), "optimal path should have length 3")
		suite.Equal(cellA, path.Cells[0], "path should start at A")
		suite.Equal(cellE, path.Cells[len(path.Cells)-1], "path should end at E")
	}
}
