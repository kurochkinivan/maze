package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildPath(t *testing.T) {
	c1 := &Cell{}
	c2 := &Cell{}
	c3 := &Cell{}

	tests := []struct {
		name     string
		previous map[*Cell]*Cell
		end      *Cell
		want     []*Cell
	}{
		{
			name: "build path success",
			previous: map[*Cell]*Cell{
				c3: c2,
				c2: c1,
				c1: nil,
			},
			end:  c3,
			want: []*Cell{c1, c2, c3},
		},
		{
			name: "two elements path",
			previous: map[*Cell]*Cell{
				c2: c1,
				c1: nil,
			},
			end:  c2,
			want: []*Cell{c1, c2},
		},
		{
			name:     "single cell",
			previous: map[*Cell]*Cell{},
			end:      c1,
			want:     []*Cell{c1},
		},
		{
			name: "nil end",
			previous: map[*Cell]*Cell{
				c2: c1,
			},
			end:  nil,
			want: []*Cell{},
		},
	}

	for _, tt := range tests {
		path := BuildPath(tt.previous, tt.end)
		assert.Equal(t, path.Cells, tt.want)
	}
}

func TestReversePath(t *testing.T) {
	c1 := &Cell{}
	c2 := &Cell{}
	c3 := &Cell{}

	tests := []struct {
		name  string
		input []*Cell
		want  []*Cell
	}{
		{
			name:  "reverse odd number of elements",
			input: []*Cell{c1, c2, c3},
			want:  []*Cell{c3, c2, c1},
		},
		{
			name:  "reverse even number of elements",
			input: []*Cell{c1, c3},
			want:  []*Cell{c3, c1},
		},
		{
			name:  "empty input",
			input: []*Cell{},
			want:  []*Cell{},
		},
	}

	for _, tt := range tests {
		p := &Path{Cells: tt.input}
		p.ReversePath()

		assert.Equal(t, p.Cells, tt.want)
	}
}
