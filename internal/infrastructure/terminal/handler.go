package terminal

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/urfave/cli/v3"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
	mazeio "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/infrastructure/io"
)

// Handler is a CLI handler that uses a MazeService to generate and solve mazes.
type Handler struct {
	mazeService MazeService
}

// MazeService defines the methods the Handler depends on to generate
// and solve mazes.
type MazeService interface {
	GenerateMaze(algorithm string, width, height int) (*maze.Maze, error)
	SolveMaze(algorithm string, m *maze.Maze, start, end *entities.Cell) (*entities.Path, error)
}

// New creates a new Handler with the provided MazeService.
func New(mazeService MazeService) *Handler {
	return &Handler{
		mazeService: mazeService,
	}
}

// handleGenerate handles the CLI command to generate a maze.
func (h *Handler) handleGenerate(_ context.Context, cmd *cli.Command) error {
	algorithm := cmd.String("algorithm")
	width := cmd.Int("width")
	height := cmd.Int("height")
	output := cmd.String("output")

	m, err := h.mazeService.GenerateMaze(algorithm, width, height)
	if err != nil {
		return fmt.Errorf("failed to generate maze: %w", err)
	}

	if err = h.WriteMaze(output, m); err != nil {
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

	startPoint := entities.NewPoint(start[0], start[1])
	endPoint := entities.NewPoint(end[0], end[1])

	m, err := h.ReadMaze(file)
	if err != nil {
		return fmt.Errorf("failed to read maze: %w", err)
	}

	startCell := m.Cell(startPoint.Row(), startPoint.Col())
	endCell := m.Cell(endPoint.Row(), endPoint.Col())

	path, err := h.mazeService.SolveMaze(algorithm, m, startCell, endCell)
	if err != nil {
		return fmt.Errorf("failed to solve maze: %w", err)
	}

	if err = h.WriteMazeWithSolution(output, m, startCell, endCell, path); err != nil {
		return fmt.Errorf("failed to write solution: %w", err)
	}

	return nil
}

// ReadMaze reads a maze from the provided filename or from stdin if filename is empty.
// It checks the file existence, opens it and delegates parsing to the mazeio package.
func (h *Handler) ReadMaze(filename string) (*maze.Maze, error) {
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

	return mazeio.ReadMaze(reader)
}

// WriteMaze writes the provided maze to the given filename or to stdout if filename is empty.
// It creates/truncates the file when a filename is provided and delegates formatting to mazeio.
func (h *Handler) WriteMaze(filename string, m *maze.Maze) error {
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

	return mazeio.WriteMaze(writer, m)
}

// WriteMazeWithSolution writes the maze with the solution path to the given filename
// or stdout if filename is empty. It delegates the actual rendering to mazeio.
func (h *Handler) WriteMazeWithSolution(
	filename string,
	m *maze.Maze,
	start, end *entities.Cell,
	path *entities.Path,
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

	return mazeio.WriteMazeWithSolution(writer, m, start, end, path)
}
