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

type Handler struct {
	mazeService MazeService
}

type MazeService interface {
	GenerateMaze(algorithm string, width, height int) (*maze.Maze, error)
	SolveMaze(algorithm string, m *maze.Maze, start, end *entities.Cell) (*entities.Path, error)
}

func New(mazeService MazeService) *Handler {
	return &Handler{
		mazeService: mazeService,
	}
}

func (h *Handler) handleGenerate(c context.Context, cmd *cli.Command) error {
	algorithm := cmd.String("algorithm")
	width := cmd.Int("width")
	height := cmd.Int("height")
	output := cmd.String("output")

	m, err := h.mazeService.GenerateMaze(algorithm, width, height)
	if err != nil {
		return fmt.Errorf("failed to generate maze: %w", err)
	}

	if err := h.WriteMaze(output, m); err != nil {
		return fmt.Errorf("failed to write maze: %w", err)
	}

	return nil
}

func (h *Handler) handlerSolve(c context.Context, cmd *cli.Command) error {
	algorithm := cmd.String("algorithm")
	file := cmd.String("file")
	start := cmd.String("start")
	end := cmd.String("end")
	output := cmd.String("output")

	startPoint, err := parseToPoint(start)
	if err != nil {
		return fmt.Errorf("failed to parse start point: %w", err)
	}

	endPoint, err := parseToPoint(end)
	if err != nil {
		return fmt.Errorf("failed to parse end point: %w", err)
	}

	m, err := h.ReadMaze(file)
	if err != nil {
		return fmt.Errorf("failed to read maze: %w", err)
	}

	startCell := m.Cell(startPoint.Row, startPoint.Col)
	endCell := m.Cell(endPoint.Row, endPoint.Col)

	path, err := h.mazeService.SolveMaze(algorithm, m, startCell, endCell)
	if err != nil {
		return fmt.Errorf("failed to solve maze: %w", err)
	}

	if err := h.WriteMazeWithSolution(output, m, startCell, endCell, path); err != nil {
		return fmt.Errorf("failed to write solution: %w", err)
	}

	return nil
}

func (h *Handler) ReadMaze(filename string) (*maze.Maze, error) {
	var reader io.Reader = os.Stdin

	if filename != "" {
		if _, err := os.Stat(filename); err != nil {
			return nil, fmt.Errorf("failed to locate file %q: %w", filename, err)
		}

		f, err := os.Open(filename)
		if err != nil {
			return nil, fmt.Errorf("failed to open file: %w", err)
		}
		defer f.Close()

		reader = f
	} else {
		fmt.Fprintf(os.Stdout, "Enter maze (paste maze text, then press Ctrl+D on Unix or Ctrl+Z on Windows):\n")
	}

	return mazeio.ReadMaze(reader)
}

func (h *Handler) WriteMaze(filename string, m *maze.Maze) error {
	var writer io.Writer = os.Stdout

	if filename != "" {
		f, err := os.Create(filename)
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}
		defer f.Close()

		writer = f
	}

	return mazeio.WriteMaze(writer, m)
}

func (h *Handler) WriteMazeWithSolution(
	filename string,
	m *maze.Maze,
	start, end *entities.Cell,
	path *entities.Path,
) error {
	var writer io.Writer = os.Stdout

	if filename != "" {
		f, err := os.Create(filename)
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}
		defer f.Close()

		writer = f
	}

	return mazeio.WriteMazeWithSolution(writer, m, start, end, path)
}
