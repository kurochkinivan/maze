package maze

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"
)

func TestReachableNeighbors_HappyPath(t *testing.T) {
	maze := New(3, 3)

	centralCell := maze.Cell(1, 1)
	rightCell := maze.Cell(1, 2)
	leftCell := maze.Cell(1, 0)

	DirectionRight.RemoveWall(centralCell, rightCell)
	DirectionLeft.RemoveWall(centralCell, leftCell)

	reachable := maze.ReachableNeighbors(centralCell)

	assert.Len(t, reachable, 2)
	assert.Contains(t, reachable, rightCell)
	assert.Contains(t, reachable, leftCell)
}

func TestReachableNeighbors_NoNeigbors(t *testing.T) {
	maze := New(3, 3)

	centralCell := maze.Cell(1, 1)
	reachable := maze.ReachableNeighbors(centralCell)

	assert.Empty(t, reachable)
}

func TestNeighbours(t *testing.T) {
	maze := New(3, 3)

	tests := []struct {
		name string
		cell *entities.Cell
		want map[*entities.Cell]Direction
	}{
		{
			name: "central cell",
			cell: maze.Cell(1, 1),
			want: map[*entities.Cell]Direction{
				maze.Cell(0, 1): DirectionUp,
				maze.Cell(1, 0): DirectionLeft,
				maze.Cell(1, 2): DirectionRight,
				maze.Cell(2, 1): DirectionDown,
			},
		},
		{
			name: "border cell",
			cell: maze.Cell(0, 0),
			want: map[*entities.Cell]Direction{
				maze.Cell(0, 1): DirectionRight,
				maze.Cell(1, 0): DirectionDown,
			},
		},
		{
			name: "central top cell",
			cell: maze.Cell(0, 1),
			want: map[*entities.Cell]Direction{
				maze.Cell(0, 2): DirectionRight,
				maze.Cell(0, 0): DirectionLeft,
				maze.Cell(1, 1): DirectionDown,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			neighbors := maze.neighbors(tt.cell)
			require.Len(t, neighbors, len(tt.want))

			for _, neighbor := range neighbors {
				assert.Contains(t, tt.want, neighbor.Cell)
				assert.Equal(t, tt.want[neighbor.Cell], neighbor.Direction)
			}
		})
	}
}

func TestNeighbors_NoNeighbors(t *testing.T) {
	maze := New(1, 1)

	cell := maze.Cell(0, 0)
	neighbors := maze.neighbors(cell)

	assert.Empty(t, neighbors)
}

func TestIsValid(t *testing.T) {
	width, height := 5, 10
	maze := New(width, height)

	tests := []struct {
		name  string
		point entities.Point
		want  bool
	}{
		{
			name:  "regular point",
			point: entities.NewPoint(height/2, width/2),
			want:  true,
		},
		{
			name:  "starting point",
			point: entities.NewPoint(0, 0),
			want:  true,
		},
		{
			name:  "ending point",
			point: entities.NewPoint(height-1, width-1),
			want:  true,
		},
		{
			name:  "out of bounds +1",
			point: entities.NewPoint(height, width),
			want:  false,
		},
		{
			name:  "out of bounds -1",
			point: entities.NewPoint(-1, -1),
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maze.IsValid(tt.point)
			assert.Equal(t, tt.want, got)
		})
	}
}
