package entities_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"
)

func TestNewPath_Empty(t *testing.T) {
	path := entities.NewPath([]*entities.Cell{})
	require.Len(t, path, 0)
}

func TestNewPath_SingleCell(t *testing.T) {
	cell := entities.NewCell(2, 3)
	path := entities.NewPath([]*entities.Cell{cell})

	require.Len(t, path, 1)
	assert.Equal(t, 2, path[0].Row())
	assert.Equal(t, 3, path[0].Col())
}

func TestNewPath_MultipleCells(t *testing.T) {
	c1 := entities.NewCell(0, 0)
	c2 := entities.NewCell(1, 1)
	c3 := entities.NewCell(2, 2)

	path := entities.NewPath([]*entities.Cell{c1, c2, c3})
	require.Len(t, path, 3)

	assert.Equal(t, 0, path[0].Row())
	assert.Equal(t, 0, path[0].Col())

	assert.Equal(t, 1, path[1].Row())
	assert.Equal(t, 1, path[1].Col())

	assert.Equal(t, 2, path[2].Row())
	assert.Equal(t, 2, path[2].Col())
}

func TestNewPath_NoMutationOnCellChange(t *testing.T) {
	cell := entities.NewCell(5, 6)
	path := entities.NewPath([]*entities.Cell{cell})

	// mutate original cell
	cell.Point = entities.NewPoint(9, 9)

	// path must remain unchanged
	assert.Equal(t, 5, path[0].Row())
	assert.Equal(t, 6, path[0].Col())
}
