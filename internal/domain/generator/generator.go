package generator

import (
	"math/rand/v2"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
)

type BaseGenerator struct {
	rnd *rand.Rand
}

func NewBaseGenerator(opts ...GeneratorOption) BaseGenerator {
	g := BaseGenerator{
		rnd: rand.New(rand.NewPCG(rand.Uint64(), rand.Uint64())),
	}

	for _, opt := range opts {
		opt(&g)
	}

	return g
}

type GeneratorOption func(*BaseGenerator)

func WithSeed(seed1, seed2 uint64) GeneratorOption {
	return func(g *BaseGenerator) {
		source := rand.NewPCG(seed1, seed2)
		g.rnd = rand.New(source)
	}
}

func (g *BaseGenerator) Rand() *rand.Rand {
	return g.rnd
}

func (g *BaseGenerator) RandomCell(m *maze.Maze) *entities.Cell {
	row := g.Rand().IntN(m.Rows())
	col := g.Rand().IntN(m.Cols())
	return m.Cell(row, col)
}
