package test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/generator"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/generator/dfs"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/generator/prim"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/solver"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/solver/astar"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/solver/dijkstra"
)

type IntegrationTestSuite struct {
	suite.Suite
	generator.Generator
	solver.Solver
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
