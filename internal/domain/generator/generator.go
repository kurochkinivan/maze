package generator

import (
	"math/rand/v2"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
)

// RandomGenerator provides randomization utilities used by maze generators.
type RandomGenerator struct {
	rnd *rand.Rand
}

// NewRandomGenerator creates a RandomGenerator with a random seed is used.
func NewRandomGenerator() RandomGenerator {
	return RandomGenerator{
		rnd: rand.New(rand.NewPCG(rand.Uint64(), rand.Uint64())),
	}
}

// IntN returns a random integer in the range [0, n).
func (g *RandomGenerator) IntN(n int) int {
	return g.rnd.IntN(n)
}

// RandomCell selects and returns a random cell from the maze.
func (g *RandomGenerator) RandomCell(m *maze.Maze) *entities.Cell {
	row := g.IntN(m.Rows())
	col := g.IntN(m.Cols())
	return m.Cell(row, col)
}
