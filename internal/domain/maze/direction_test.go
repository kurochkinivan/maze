package maze

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"
)

func TestRemoveWall(t *testing.T) {
	tests := []struct {
		name      string
		direction Direction
		from      *entities.Cell
		to        *entities.Cell
		wantFrom  entities.Walls
		wantTo    entities.Walls
	}{
		{
			name:      "remove up",
			direction: DirectionUp,
			from:      entities.NewCell(0, 0),
			to:        entities.NewCell(0, 0),
			wantFrom: entities.Walls{
				Top: false, Bottom: true, Left: true, Right: true,
			},
			wantTo: entities.Walls{
				Top: true, Bottom: false, Left: true, Right: true,
			},
		},
		{
			name:      "remove down",
			direction: DirectionDown,
			from:      entities.NewCell(0, 0),
			to:        entities.NewCell(0, 0),
			wantFrom: entities.Walls{
				Top: true, Bottom: false, Left: true, Right: true,
			},
			wantTo: entities.Walls{
				Top: false, Bottom: true, Left: true, Right: true,
			},
		},
		{
			name:      "remove left",
			direction: DirectionLeft,
			from:      entities.NewCell(0, 0),
			to:        entities.NewCell(0, 0),
			wantFrom: entities.Walls{
				Top: true, Bottom: true, Left: false, Right: true,
			},
			wantTo: entities.Walls{
				Top: true, Bottom: true, Left: true, Right: false,
			},
		},
		{
			name:      "remove right",
			direction: DirectionRight,
			from:      entities.NewCell(0, 0),
			to:        entities.NewCell(0, 0),
			wantFrom: entities.Walls{
				Top: true, Bottom: true, Left: true, Right: false,
			},
			wantTo: entities.Walls{
				Top: true, Bottom: true, Left: false, Right: true,
			},
		},
		{
			name:      "unknown direction",
			direction: Direction{},
			from:      entities.NewCell(0, 0),
			to:        entities.NewCell(0, 0),
			wantFrom: entities.Walls{
				Top: true, Bottom: true, Left: true, Right: true,
			},
			wantTo: entities.Walls{
				Top: true, Bottom: true, Left: true, Right: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.direction.RemoveWall(tt.from, tt.to)
			assert.Equal(t, tt.wantFrom, tt.from.Walls)
			assert.Equal(t, tt.wantTo, tt.to.Walls)
		})
	}
}

func TestHasWall(t *testing.T) {
	tests := []struct {
		name      string
		direction Direction
		cell      *entities.Cell
		want      bool
	}{
		{
			name:      "has top wall",
			direction: DirectionUp,
			cell:      &entities.Cell{Walls: entities.Walls{Top: true}},
			want:      true,
		},
		{
			name:      "doesn't have top wall",
			direction: DirectionUp,
			cell:      &entities.Cell{Walls: entities.Walls{Top: false}},
			want:      false,
		},
		{
			name:      "has bottom wall",
			direction: DirectionDown,
			cell:      &entities.Cell{Walls: entities.Walls{Bottom: true}},
			want:      true,
		},
		{
			name:      "doesn't have bottom wall",
			direction: DirectionDown,
			cell:      &entities.Cell{Walls: entities.Walls{Bottom: false}},
			want:      false,
		},
		{
			name:      "has left wall",
			direction: DirectionLeft,
			cell:      &entities.Cell{Walls: entities.Walls{Left: true}},
			want:      true,
		},
		{
			name:      "doesn't have left wall",
			direction: DirectionLeft,
			cell:      &entities.Cell{Walls: entities.Walls{Left: false}},
			want:      false,
		},
		{
			name:      "has right wall",
			direction: DirectionRight,
			cell:      &entities.Cell{Walls: entities.Walls{Right: true}},
			want:      true,
		},
		{
			name:      "doesn't have right wall",
			direction: DirectionRight,
			cell:      &entities.Cell{Walls: entities.Walls{Right: false}},
			want:      false,
		},
		{
			name:      "unknown direction",
			direction: Direction{},
			cell:      &entities.Cell{Walls: entities.Walls{Top: false, Right: false, Bottom: false, Left: false}},
			want:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.direction.HasWall(tt.cell)
			assert.Equal(t, tt.want, got)
		})
	}
}
