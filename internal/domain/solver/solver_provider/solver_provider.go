package solver_provider

import (
	"fmt"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/solver"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/solver/astar"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/solver/dijkstra"
)

type SolverProvider struct{}

func New() *SolverProvider {
	return &SolverProvider{}
}

func (p *SolverProvider) Algorithm(algo solver.Algorithm) (solver.Solver, error) {
	switch algo {
	case solver.AlgoAStar:
		return astar.New(), nil
	case solver.AlgoDijkstra:
		return dijkstra.New(), nil
	default:
		return nil, fmt.Errorf("unknown algorithm %q", algo)
	}
}
