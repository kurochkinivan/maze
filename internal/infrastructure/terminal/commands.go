package terminal

import "github.com/urfave/cli/v3"

func (h *Handler) GenerateCommand() *cli.Command {
	return &cli.Command{
		Name:  "generate",
		Usage: "Generate a maze using specified algorithm",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "algorithm",
				Value:    "dfs",
				Required: false,
				Usage:    "maze generation algorithm. Options: dfs, prim",
			},
			&cli.IntFlag{
				Name:     "width",
				Required: true,
				Usage:    "width of the generated maze",
			},
			&cli.IntFlag{
				Name:     "height",
				Required: true,
				Usage:    "height of the generated maze",
			},
			&cli.StringFlag{
				Name:     "output",
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
				Value:    "astar",
				Required: false,
				Usage:    "path finding algorithm. Options: astar, dijkstra",
			},
			&cli.StringFlag{
				Name:     "file",
				Required: false,
				Usage:    "input file with maze description. If not specified, program expects maze in the stdin",
			},
			&cli.StringFlag{
				Name:     "start",
				Required: true,
				Usage:    "starting point coordinates in format: x,y",
			},
			&cli.StringFlag{
				Name:     "end",
				Required: true,
				Usage:    "ending point coordinates in format: x,y",
			},
			&cli.StringFlag{
				Name:     "output",
				Required: false,
				Usage:    "output file path for the solution. If not specified, solution will be printed to console",
			},
		},
		Action: h.handlerSolve,
	}
}
