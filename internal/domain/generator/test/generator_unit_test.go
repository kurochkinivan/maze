package test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/generator/dfs"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/generator/prim"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
)

type GeneratorTestSuite struct {
	suite.Suite
	Generator
}

func TestGeneratorTestSuiteDFS(t *testing.T) {
	testSuite := new(GeneratorTestSuite)
	testSuite.Generator = dfs.New()
	suite.Run(t, testSuite)
}

func TestGeneratorTestSuitePrim(t *testing.T) {
	testSuite := new(GeneratorTestSuite)
	testSuite.Generator = prim.New()
	suite.Run(t, testSuite)
}

type Generator interface {
	Generate(m *maze.Maze)
}

func (suite *GeneratorTestSuite) TestGenerate_EmptyMaze() {
	m := maze.New(0, 0)

	suite.NotPanics(func() {
		suite.Generate(m)
	})
}

// Generation should change nothing. All walls should stay.
func (suite *GeneratorTestSuite) TestGenerate_SingleCell() {
	m := maze.New(1, 1)

	suite.Generator.Generate(m)

	cell := m.Cell(0, 0)

	suite.True(cell.Walls.Top, "top wall should exist")
	suite.True(cell.Walls.Right, "right wall should exist")
	suite.True(cell.Walls.Bottom, "bottom wall should exist")
	suite.True(cell.Walls.Left, "left wall should exist")
}

// TestGenerate_ModifiesMaze checks that Generator actually changes the maze.
func (suite *GeneratorTestSuite) TestGenerate_ModifiesMaze() {
	m := maze.New(5, 5)

	wallsBefore := suite.countWalls(m)

	suite.Generator.Generate(m)

	wallsAfter := suite.countWalls(m)

	suite.Less(wallsAfter, wallsBefore, "generator should remove walls")
}

// TestGenerate_LinearMaze checks generation for linear maze.
func (suite *GeneratorTestSuite) TestGenerate_LinearMaze() {
	tests := []struct {
		name   string
		width  int
		height int
	}{
		{"horizontal", 5, 1},
		{"vertical", 1, 5},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			m := maze.New(tt.width, tt.height)

			suite.NotPanics(func() {
				suite.Generator.Generate(m)
			})

			wallsCount := suite.countWalls(m)
			totalPossibleWalls := suite.initialWalls(m)

			wantDifference := (tt.height*tt.width - 1) * 2
			gotDifference := totalPossibleWalls - wallsCount

			// all walls between cells should be removed
			suite.Equal(wantDifference, gotDifference)
		})
	}
}

// countWalls counts the number of walls in the maze.
func (suite *GeneratorTestSuite) countWalls(m *maze.Maze) int {
	var cnt int

	for row := 0; row < m.Rows; row++ {
		for col := 0; col < m.Cols; col++ {
			cell := m.Cell(row, col)

			if cell.Walls.Top {
				cnt++
			}
			if cell.Walls.Right {
				cnt++
			}
			if cell.Walls.Bottom {
				cnt++
			}
			if cell.Walls.Left {
				cnt++
			}
		}
	}

	return cnt
}

// initialWalls returns the initial amount of walls in the maze
func (suite *GeneratorTestSuite) initialWalls(m *maze.Maze) int {
	return m.Cols * m.Rows * 4
}
