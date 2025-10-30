package application

import (
	"errors"
	"fmt"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
)

// MazeService manages maze generation and solving using different algorithms.
type MazeService struct {
	generators map[string]Generator
	solvers    map[string]Solver
}

// Generator defines the interface for maze generation algorithms.
type Generator interface {
	Generate(m *maze.Maze)
}

// Solver defines the interface for maze solving algorithms.
type Solver interface {
	Solve(m *maze.Maze, start, end *entities.Cell) (*entities.Path, bool)
}

// NewMazeService creates a new MazeService with empty registries.
func NewMazeService() *MazeService {
	return &MazeService{
		generators: make(map[string]Generator),
		solvers:    make(map[string]Solver),
	}
}

// RegisterGenerator registers a maze generator with a name.
func (s *MazeService) RegisterGenerator(name string, generator Generator) {
	s.generators[name] = generator
}

// RegisterSolver registers a maze solver with a name.
func (s *MazeService) RegisterSolver(name string, solver Solver) {
	s.solvers[name] = solver
}

// GenerateMaze creates a new maze of the given size using the specified generator algorithm.
// Returns an error if the algorithm is unknown.
func (s *MazeService) GenerateMaze(algorithm string, width, height int) (*maze.Maze, error) {
	gen, ok := s.generators[algorithm]
	if !ok {
		return nil, fmt.Errorf("unknown algorithm %q", algorithm)
	}

	m := maze.New(width, height)
	gen.Generate(m)

	return m, nil
}

// SolveMaze finds a path in the maze from start to end using the specified solver algorithm.
// Returns an error if the algorithm is unknown or no solution exists.
func (s *MazeService) SolveMaze(algorithm string, m *maze.Maze, start, end *entities.Cell) (*entities.Path, error) {
	solver, ok := s.solvers[algorithm]
	if !ok {
		return nil, fmt.Errorf("unknown algorithm %q", algorithm)
	}

	path, ok := solver.Solve(m, start, end)
	if !ok {
		return nil, errors.New("no solution")
	}

	return path, nil
}
