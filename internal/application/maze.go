package application

import (
	"fmt"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
)

type MazeService struct {
	generators map[string]Generator
	solvers    map[string]Solver
}

type Generator interface {
	Generate(m *maze.Maze)
}

type Solver interface {
	Solve(m *maze.Maze, start, end *entities.Cell) (*entities.Path, bool)
}

func NewMazeService() *MazeService {
	return &MazeService{
		generators: make(map[string]Generator),
		solvers:    make(map[string]Solver),
	}
}

func (s *MazeService) RegisterGenerator(name string, generator Generator) {
	s.generators[name] = generator
}

func (s *MazeService) RegisterSolver(name string, solver Solver) {
	s.solvers[name] = solver
}

func (s *MazeService) GenerateMaze(algorithm string, width, height int) (*maze.Maze, error) {
	gen, ok := s.generators[algorithm]
	if !ok {
		return nil, fmt.Errorf("unknown algorithm %q", algorithm)
	}

	return maze.New(width, height).Generate(gen), nil
}

func (s *MazeService) SolveMaze(algorithm string, m *maze.Maze, start, end *entities.Cell) (*entities.Path, error) {
	solver, ok := s.solvers[algorithm]
	if !ok {
		return nil, fmt.Errorf("unknown algorithm %q", algorithm)
	}

	path, ok := solver.Solve(m, start, end)
	if !ok {
		return nil, fmt.Errorf("no solution")
	}

	return path, nil
}
