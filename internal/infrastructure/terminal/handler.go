package terminal

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/urfave/cli/v3"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
	maze_reader "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/infrastructure/io/reader"
	maze_writer "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/infrastructure/io/writer"
)

// Handler is a CLI handler that uses a MazeService to generate and solve mazes.
type Handler struct {
	mazeService MazeService
	version     string
}

// MazeService defines the methods the Handler depends on to generate
// and solve mazes.
type MazeService interface {
	GenerateMaze(algorithm string, width, height int) (*maze.Maze, error)
	SolveMaze(algorithm string, m *maze.Maze, start, end *entities.Cell) (*entities.Path, error)
}

// New creates a new Handler with the provided MazeService.
func New(mazeService MazeService, version string) *Handler {
	return &Handler{
		mazeService: mazeService,
		version:     version,
	}
}

// handleGenerate handles the CLI command to generate a maze.
func (h *Handler) handleGenerate(_ context.Context, cmd *cli.Command) error {
	algorithm := cmd.String("algorithm")
	width := cmd.Int("width")
	height := cmd.Int("height")
	output := cmd.String("output")
	unicode := cmd.Bool("unicode")

	m, err := h.mazeService.GenerateMaze(algorithm, width, height)
	if err != nil {
		return fmt.Errorf("failed to generate maze: %w", err)
	}

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

	startPoint := entities.NewPoint(start[1], start[0])
	endPoint := entities.NewPoint(end[1], end[0])

	m, err := h.readMaze(file)
	if err != nil {
		return fmt.Errorf("failed to read maze: %w", err)
	}

	if !m.IsValid(startPoint) {
		return fmt.Errorf("starting point (%d, %d) is out of bounds", startPoint.Col(), startPoint.Row())
	}
	if !m.IsValid(endPoint) {
		return fmt.Errorf("ending point (%d, %d) is out of bounds", endPoint.Col(), endPoint.Row())
	}

	startCell := m.Cell(startPoint.Row(), startPoint.Col())
	endCell := m.Cell(endPoint.Row(), endPoint.Col())

	path, err := h.mazeService.SolveMaze(algorithm, m, startCell, endCell)
	if err != nil {
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
	path *entities.Path,
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
