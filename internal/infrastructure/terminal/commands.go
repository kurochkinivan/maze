package terminal

import (
	"context"
	"errors"
	"fmt"

	"github.com/urfave/cli/v3"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/generator"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/solver"
)

func (h *Handler) Run(ctx context.Context, osArgs []string) error {
	return h.app().Run(ctx, osArgs)
}

func (h *Handler) app() *cli.Command {
	return &cli.Command{
		Name:  "maze",
		Usage: "Generate and solve mazes",
		Commands: []*cli.Command{
			h.generateCommand(),
			h.solveCommand(),
		},
		Version: h.version,
	}
}

// generateCommand returns a configured *cli.Command for the "generate" subcommand.
func (h *Handler) generateCommand() *cli.Command {
	return &cli.Command{
		Name:  "generate",
		Usage: "Generate a maze using specified algorithm",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:      "algorithm",
				Aliases:   []string{"a"},
				Value:     "dfs",
				Required:  false,
				Usage:     "maze generation algorithm. Options: dfs, prim",
				Validator: validateGeneratorAlgoritm,
			},
			&cli.IntFlag{
				Name:      "width",
				Aliases:   []string{"W"},
				Required:  true,
				Usage:     "width of the generated maze",
				Validator: validateGreaterThan0,
			},
			&cli.IntFlag{
				Name:      "height",
				Aliases:   []string{"H"},
				Required:  true,
				Usage:     "height of the generated maze",
				Validator: validateGreaterThan0,
			},
			&cli.StringFlag{
				Name:     "output",
				Aliases:  []string{"o"},
				Required: false,
				Usage:    "output file path. If not specified, maze will be printed to console",
			},
			&cli.BoolFlag{
				Name:     "unicode",
				Aliases:  []string{"u"},
				Value:    false,
				Required: false,
				Usage:    "Render the maze using Unicode symbols instead of ASCII characters",
			},
		},
		Action: h.handleGenerate,
	}
}

// solveCommand returns a configured *cli.Command for the "solve" subcommand.
func (h *Handler) solveCommand() *cli.Command {
	return &cli.Command{
		Name:  "solve",
		Usage: "Solve a maze using specified algorithm",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:      "algorithm",
				Aliases:   []string{"a"},
				Value:     "astar",
				Required:  false,
				Usage:     "path finding algorithm. Options: astar, dijkstra",
				Validator: validateSolverAlgorithm,
			},
			&cli.StringFlag{
				Name:     "file",
				Aliases:  []string{"f"},
				Required: false,
				Usage:    "input file with maze description. If not specified, program expects maze in the stdin",
			},
			&cli.IntSliceFlag{
				Name:      "start",
				Aliases:   []string{"s"},
				Required:  true,
				Usage:     "starting point coordinates in format: x,y",
				Validator: validateCoordinates,
			},
			&cli.IntSliceFlag{
				Name:      "end",
				Aliases:   []string{"e"},
				Required:  true,
				Usage:     "ending point coordinates in format: x,y",
				Validator: validateCoordinates,
			},
			&cli.StringFlag{
				Name:     "output",
				Aliases:  []string{"o"},
				Required: false,
				Usage:    "output file path for the solution. If not specified, solution will be printed to console",
			},
			&cli.BoolFlag{
				Name:     "unicode",
				Aliases:  []string{"u"},
				Value:    false,
				Required: false,
				Usage:    "Render the maze using Unicode symbols instead of ASCII characters",
			},
		},
		Action: h.handlerSolve,
	}
}

func validateGeneratorAlgoritm(algorithm string) error {
	if !generator.Algorithm(algorithm).IsValid() {
		return fmt.Errorf("unknown algorithm %q", algorithm)
	}
	return nil
}

func validateSolverAlgorithm(algorithm string) error {
	if !solver.Algorithm(algorithm).IsValid() {
		return fmt.Errorf("unknown algorithm %q", algorithm)
	}
	return nil
}

func validateCoordinates(coords []int) error {
	if len(coords) != 2 {
		return fmt.Errorf("invalid amount of coordinates, expected 2, got %d", len(coords))
	}
	if coords[0] < 0 || coords[1] < 0 {
		return errors.New("coordinates should be positive")
	}

	return nil
}

func validateGreaterThan0(number int) error {
	if number <= 0 {
		return errors.New("should be greater than 0")
	}
	return nil
}
