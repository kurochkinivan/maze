package test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/generator/dfs"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/generator/prim"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/solver/astar"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/solver/dijkstra"
)

type Generator interface {
	Generate(m *maze.Maze)
}

type Solver interface {
	Solve(m *maze.Maze, start *entities.Cell, end *entities.Cell) (*entities.Path, bool)
}

type IntegrationTestSuite struct {
	suite.Suite
	Generator
	Solver
}

func TestIntegrationTestSuite_DFS_Astar(t *testing.T) {
	testSuite := new(IntegrationTestSuite)
	testSuite.Generator = dfs.New()
	testSuite.Solver = astar.New()
	suite.Run(t, testSuite)
}

func TestIntegrationTestSuite_DFS_Dijkstra(t *testing.T) {
	testSuite := new(IntegrationTestSuite)
	testSuite.Generator = dfs.New()
	testSuite.Solver = dijkstra.New()
	suite.Run(t, testSuite)
}

func TestIntegrationTestSuite_Prim_Astar(t *testing.T) {
	testSuite := new(IntegrationTestSuite)
	testSuite.Generator = prim.New()
	testSuite.Solver = astar.New()
	suite.Run(t, testSuite)
}

func TestIntegrationTestSuite_Prim_Dijkstra(t *testing.T) {
	testSuite := new(IntegrationTestSuite)
	testSuite.Generator = prim.New()
	testSuite.Solver = dijkstra.New()
	suite.Run(t, testSuite)
}

// TestNoIsolatedAreas checks that in generated maze there are no isolated areas.
// It checks if the solver can find a solution from any cell to another.
func (suite *IntegrationTestSuite) TestNoIsolatedAreas() {
	m := maze.New(10, 10)
	suite.Generator.Generate(m)

	start := m.Cell(0, 0)

	for row := range m.Rows() {
		for col := range m.Cols() {
			end := m.Cell(row, col)

			path, ok := suite.Solver.Solve(m, start, end)
			suite.Require().True(ok, "there must be no isolated areas")
			suite.NotNil(path, "there must be a path")
		}
	}
}

// TestSolutionIsWalkable checks if the given solution is valid.
func (suite *IntegrationTestSuite) TestSolutionIsWalkable() {
	width, height := 10, 10

	m := maze.New(width, height)
	suite.Generator.Generate(m)

	start := m.Cell(0, 0)
	end := m.Cell(height-1, width-1)

	path, ok := suite.Solver.Solve(m, start, end)
	suite.Require().True(ok, "there must be no isolated areas")
	suite.NotNil(path, "there must be a path")

	for i := 1; i < len(path.Cells); i++ {
		prev, cur := path.Cells[i-1], path.Cells[i]

		rowDiff := cur.Row() - prev.Row()
		colDiff := cur.Col() - prev.Col()

		switch {
		case rowDiff == 0 && colDiff == -1:
			suite.Require().False(prev.Walls.Left, "unexpected Left wall on the path")
			suite.Require().False(cur.Walls.Right, "unexpected Right wall on the path")

		case rowDiff == 0 && colDiff == 1:
			suite.Require().False(prev.Walls.Right, "unexpected Right wall on the path")
			suite.Require().False(cur.Walls.Left, "unexpected Left wall on the path")

		case rowDiff == 1 && colDiff == 0:
			suite.Require().False(prev.Walls.Bottom, "unexpected Bottom wall on the path")
			suite.Require().False(cur.Walls.Top, "unexpected Top wall on the path")

		case rowDiff == -1 && colDiff == 0:
			suite.Require().False(prev.Walls.Top, "unexpected Top wall on the path")
			suite.Require().False(cur.Walls.Bottom, "unexpected Bottom wall on the path")

		default:
			suite.T().Fatalf("invalid movement detected between %v and %v", prev, cur)
		}
	}
}
