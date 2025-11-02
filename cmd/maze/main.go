package main

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/application"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/generator/dfs"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/generator/prim"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/solver/astar"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/solver/dijkstra"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/infrastructure/terminal"
)

const (
	version = "1.0.0"
)

func main() {
	mazeService := application.NewMazeService()

	mazeService.RegisterGenerator("dfs", dfs.New())
	mazeService.RegisterGenerator("prim", prim.New())

	mazeService.RegisterSolver("dijkstra", dijkstra.New())
	mazeService.RegisterSolver("astar", astar.New())

	handler := terminal.New(mazeService)

	app := &cli.Command{
		Name:  "maze",
		Usage: "Generate and solve mazes",
		Commands: []*cli.Command{
			handler.GenerateCommand(),
			handler.SolveCommand(),
		},
		Version: version,
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
