package generator

import (
	"math/rand/v2"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
)

// BaseGenerator provides shared randomization logic for maze generators.
// It encapsulates the random number generator.
type BaseGenerator struct {
	rnd *rand.Rand
}

// NewBaseGenerator creates a new BaseGenerator with optional configuration options.
// If no options are provided, a random seed is used for unpredictable maze generation.
func NewBaseGenerator(opts ...Option) BaseGenerator {
	g := BaseGenerator{
		rnd: rand.New(rand.NewPCG(rand.Uint64(), rand.Uint64())),
	}

	for _, opt := range opts {
		opt(&g)
	}

	return g
}

// Option defines a configuration function for customizing the BaseGenerator.
type Option func(*BaseGenerator)

// WithSeed sets seeds for the random generator, ensuring reproducible mazes.
func WithSeed(seed1, seed2 uint64) Option {
	return func(g *BaseGenerator) {
		source := rand.NewPCG(seed1, seed2)
		g.rnd = rand.New(source)
	}
}

// IntN returns a random integer in the range [0, n) using the internal generator.
func (g *BaseGenerator) IntN(n int) int {
	return g.rnd.IntN(n)
}

// RandomCell selects and returns a random cell from the given maze.
func (g *BaseGenerator) RandomCell(m *maze.Maze) *entities.Cell {
	row := g.IntN(m.Rows())
	col := g.IntN(m.Cols())
	return m.Cell(row, col)
}
