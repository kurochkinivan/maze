package terminal

import (
	"fmt"

	"github.com/urfave/cli/v3"
)

func (h *Handler) GenerateCommand() *cli.Command {
	return &cli.Command{
		Name:  "generate",
		Usage: "Generate a maze using specified algorithm",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "algorithm",
				Aliases:  []string{"a"},
				Value:    "dfs",
				Required: false,
				Usage:    "maze generation algorithm. Options: dfs, prim",
				Validator: func(algorithm string) error {
					if algorithm != "dfs" && algorithm != "prim" {
						return fmt.Errorf("unknown algorithm %s", algorithm)
					}
					return nil
				},
			},
			&cli.IntFlag{
				Name:      "width",
				Aliases:   []string{"w"},
				Required:  true,
				Usage:     "width of the generated maze",
				Validator: validateGT0,
			},
			&cli.IntFlag{
				Name:      "height",
				Aliases:   []string{"h"},
				Required:  true,
				Usage:     "height of the generated maze",
				Validator: validateGT0,
			},
			&cli.StringFlag{
				Name:     "output",
				Aliases:  []string{"o"},
				Required: false,
				Usage:    "output file path. If not specified, maze will be printed to console",
			},
		},
		Action: h.handleGenerate,
	}
}

func (h *Handler) SolveCommand() *cli.Command {
	return &cli.Command{
		Name:  "solve",
		Usage: "Solve a maze using specified algorithm",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "algorithm",
				Aliases:  []string{"a"},
				Value:    "astar",
				Required: false,
				Usage:    "path finding algorithm. Options: astar, dijkstra",
				Validator: func(algorithm string) error {
					if algorithm != "astar" && algorithm != "dijkstra" {
						return fmt.Errorf("unknown algorithm %s", algorithm)
					}
					return nil
				},
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
		},
		Action: h.handlerSolve,
	}
}

func validateCoordinates(coords []int) error {
	if len(coords) != 2 {
		return fmt.Errorf("invalid amount of coordinates, expected 2, got %d", len(coords))
	}
	if coords[0] < 0 || coords[1] < 0 {
		return fmt.Errorf("coordinates should be positive")
	}

	return nil
}

func validateGT0(number int) error {
	if number <= 0 {
		return fmt.Errorf("should be greater than 0")
	}
	return nil
}
