package terminal

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/urfave/cli/v3"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/generator"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/solver"
	maze_reader "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/infrastructure/io/reader"
	maze_writer "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/infrastructure/io/writer"
)

type Handler struct {
	generatorProvider GeneratorProvider
	solverProvider    SolverProvider
	version           string
}

type GeneratorProvider interface {
	Algorithm(algorithm generator.Algorithm) (generator.Generator, error)
}

type SolverProvider interface {
	Algorithm(algorithm solver.Algorithm) (solver.Solver, error)
}

func New(
	generatorProvider GeneratorProvider,
	solverProvider SolverProvider,
	version string,
) *Handler {
	return &Handler{
		generatorProvider: generatorProvider,
		solverProvider:    solverProvider,
		version:           version,
	}
}

// handleGenerate handles the CLI command to generate a maze.
func (h *Handler) handleGenerate(_ context.Context, cmd *cli.Command) error {
	algorithm := cmd.String("algorithm")
	width := cmd.Int("width")
	height := cmd.Int("height")
	output := cmd.String("output")
	unicode := cmd.Bool("unicode")

	algo := generator.Algorithm(algorithm)
	generator, err := h.generatorProvider.Algorithm(algo)
	if err != nil {
		return fmt.Errorf("failed to get algorithm: %w", err)
	}

	m := maze.New(width, height)
	generator.Generate(m)

	if err = h.writeMaze(output, m, unicode); err != nil {
		return fmt.Errorf("failed to write maze: %w", err)
	}

	return nil
}

// handlerSolve handles the CLI command to solve a maze.
func (h *Handler) handlerSolve(_ context.Context, cmd *cli.Command) error {
	algorithm := cmd.String("algorithm")
	file := cmd.String("file")
	start := cmd.IntSlice("start")
	end := cmd.IntSlice("end")
	output := cmd.String("output")
	unicode := cmd.Bool("unicode")

	m, err := h.readMaze(file)
	if err != nil {
		return fmt.Errorf("failed to read maze: %w", err)
	}

	startCell, endCell, err := processPoints(m, start, end)
	if err != nil {
		return fmt.Errorf("failed to process start and end points: %w", err)
	}

	algo := solver.Algorithm(algorithm)
	solver, err := h.solverProvider.Algorithm(algo)
	if err != nil {
		return fmt.Errorf("failed to get algorithm: %w", err)
	}

	path, ok := solver.Solve(m, startCell, endCell)
	if !ok {
		return fmt.Errorf("failed to solve maze: %w", err)
	}

	if err = h.writeMazeWithSolution(output, m, path, unicode); err != nil {
		return fmt.Errorf("failed to write solution: %w", err)
	}

	return nil
}

// ReadMaze reads a maze from the provided filename or from stdin if filename is empty.
// It checks the file existence, opens it and delegates parsing to the mazeio package.
func (h *Handler) readMaze(filename string) (*maze.Maze, error) {
	var reader io.Reader

	switch filename {
	case "":
		reader = os.Stdin
	default:
		if _, err := os.Stat(filename); err != nil {
			return nil, fmt.Errorf("failed to locate file %q: %w", filename, err)
		}

		f, err := os.Open(filename)
		if err != nil {
			return nil, fmt.Errorf("failed to open file: %w", err)
		}
		defer f.Close()

		reader = f
	}

	return maze_reader.ReadMaze(reader)
}

// WriteMaze writes the provided maze to the given filename or to stdout if filename is empty.
// It creates/truncates the file when a filename is provided and delegates formatting to mazeio.
func (h *Handler) writeMaze(filename string, m *maze.Maze, unicode bool) error {
	var writer io.Writer

	switch filename {
	case "":
		writer = os.Stdout
	default:
		f, err := os.Create(filename)
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}
		defer f.Close()

		writer = f
	}

	return maze_writer.WriteMaze(writer, m, unicode)
}

// WriteMazeWithSolution writes the maze with the solution path to the given filename
// or stdout if filename is empty. It delegates the actual rendering to mazeio.
func (h *Handler) writeMazeWithSolution(
	filename string,
	m *maze.Maze,
	path entities.Path,
	unicode bool,
) error {
	var writer io.Writer

	switch filename {
	case "":
		writer = os.Stdout
	default:
		f, err := os.Create(filename)
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}
		defer f.Close()

		writer = f
	}

	return maze_writer.WriteMazeWithSolution(writer, m, path, unicode)
}

func processPoints(m *maze.Maze, start, end []int) (*entities.Cell, *entities.Cell, error) {
	startPoint := entities.NewPoint(start[1], start[0])
	endPoint := entities.NewPoint(end[1], end[0])

	if !m.IsValid(startPoint) {
		return nil, nil, fmt.Errorf("starting point (%d, %d) is out of bounds", startPoint.Col(), startPoint.Row())
	}
	if !m.IsValid(endPoint) {
		return nil, nil, fmt.Errorf("ending point (%d, %d) is out of bounds", endPoint.Col(), endPoint.Row())
	}

	startCell := m.Cell(startPoint.Row(), startPoint.Col())
	endCell := m.Cell(endPoint.Row(), endPoint.Col())

	return startCell, endCell, nil
}
