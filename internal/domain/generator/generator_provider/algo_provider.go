package generator_provider

import (
	"fmt"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/generator"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/generator/dfs"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/generator/prim"
)

type GeneratorProvider struct{}

func New() *GeneratorProvider {
	return &GeneratorProvider{}
}

func (g *GeneratorProvider) Algorithm(algoName generator.Algorithm) (generator.Generator, error) {
	switch algoName {
	case generator.AlgoDFS:
		return dfs.New(), nil
	case generator.AlgoPrim:
		return prim.New(), nil
	default:
		return nil, fmt.Errorf("unknown algorithm %q", algoName)
	}
}
