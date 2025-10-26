package application

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/application/mocks"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
)

func TestGenerateMaze_HappyPath(t *testing.T) {
	mazeUseCase := NewMazeService()

	const algorithm = "prim"

	mockGenerator := mocks.NewMockGenerator(t)
	mockGenerator.EXPECT().Generate(mock.AnythingOfType("*maze.Maze")).Return().Once()

	mazeUseCase.RegisterGenerator(algorithm, mockGenerator)

	m, err := mazeUseCase.GenerateMaze(algorithm, 5, 5)
	assert.NoError(t, err)
	assert.NotNil(t, m)

	mockGenerator.AssertExpectations(t)
}

func TestGenerateMaze_UnknownAlgorithm(t *testing.T) {
	mazeUseCase := NewMazeService()

	m, err := mazeUseCase.GenerateMaze("some_algorithm", 10, 10)
	assert.Error(t, err)
	assert.Nil(t, m)
}

func TestSolveMaze_HappyPath(t *testing.T) {
	mazeUseCase := NewMazeService()

	const algorithm = "astar"

	m := maze.New(2, 2)
	start, end := m.Cell(0, 0), m.Cell(0, 1)
	maze.DirectionRight.RemoveWall(start, end)
	expectPath := &entities.Path{Cells: []*entities.Cell{start, end}}

	mockSolver := mocks.NewMockSolver(t)
	mockSolver.EXPECT().Solve(m, start, end).Return(expectPath, true).Once()

	mazeUseCase.RegisterSolver(algorithm, mockSolver)

	path, err := mazeUseCase.SolveMaze(algorithm, m, start, end)
	assert.NoError(t, err)
	assert.Equal(t, expectPath, path)

	mockSolver.AssertExpectations(t)
}

func TestSolveMaze_NoSolution(t *testing.T) {
	mazeUseCase := NewMazeService()

	const algorithm = "astar"

	m := maze.New(2, 2)
	start, end := m.Cell(0, 0), m.Cell(0, 1)

	mockSolver := mocks.NewMockSolver(t)
	mockSolver.EXPECT().Solve(m, start, end).Return(nil, false).Once()

	mazeUseCase.RegisterSolver(algorithm, mockSolver)

	path, err := mazeUseCase.SolveMaze(algorithm, m, start, end)
	assert.Error(t, err)
	assert.Nil(t, path)
	assert.Contains(t, err.Error(), "no solution")

	mockSolver.AssertExpectations(t)
}

func TestSolveMaze_UnknownAlgorithm(t *testing.T) {
	mazeUseCase := NewMazeService()

	m := maze.New(2, 2)
	start, end := m.Cell(0, 0), m.Cell(0, 1)
	maze.DirectionRight.RemoveWall(start, end)

	path, err := mazeUseCase.SolveMaze("some_algorithm", m, start, end)
	assert.Error(t, err)
	assert.Nil(t, path)
}
