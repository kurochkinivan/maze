package main

import (
	"context"
	"fmt"
	"os"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/generator/generator_provider"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/solver/solver_provider"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/infrastructure/terminal"
)

var (
	version string = "dev"
)

func main() {
	genProvider := generator_provider.New()
	solverProvider := solver_provider.New()

	handler := terminal.New(genProvider, solverProvider, version)

	if err := handler.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
