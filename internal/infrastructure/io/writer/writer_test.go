package maze_writer_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
	maze_writer "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/infrastructure/io/writer"
)

func TestWriteMaze(t *testing.T) {
	maze2x2, _ := create2x2Maze()
	maze3x3, _ := create3x3Maze()

	tests := []struct {
		name        string
		m           *maze.Maze
		wantASCII   string
		wantUnicode string
	}{
		{
			name: "valid 1x1 maze",
			m:    maze.New(1, 1),
			wantASCII: strings.TrimPrefix(`
###
# #
###
`, "\n"),
			wantUnicode: strings.TrimPrefix(`
┌─┐
│ │
└─┘
`, "\n"),
		},
		{
			name: "valid 2x2 maze",
			m:    maze2x2,
			wantASCII: strings.TrimPrefix(`
#####
# # #
# # #
#   #
#####
`, "\n"),
			wantUnicode: strings.TrimPrefix(`
┌─┬─┐
│ │ │
│ ╵ │
│   │
└───┘
`, "\n"),
		},
		{
			name: "valid 3x3 maze",
			m:    maze3x3,
			wantASCII: strings.TrimPrefix(`
#######
# #   #
# # # #
# # # #
# ### #
#     #
#######
`, "\n"),
			wantUnicode: strings.TrimPrefix(`
┌─┬───┐
│ │   │
│ │ ╷ │
│ │ │ │
│ └─┘ │
│     │
└─────┘
`, "\n"),
		},
		{
			name: "maze without connections",
			m:    maze.New(3, 3),
			wantASCII: strings.TrimPrefix(`
#######
# # # #
#######
# # # #
#######
# # # #
#######
`, "\n"),
			wantUnicode: strings.TrimPrefix(`
┌─┬─┬─┐
│ │ │ │
├─┼─┼─┤
│ │ │ │
├─┼─┼─┤
│ │ │ │
└─┴─┴─┘
`, "\n"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer

			err := maze_writer.WriteMaze(&buf, tt.m, false)
			require.NoError(t, err)

			gotASCII := buf.String()
			assert.Equal(t, gotASCII, tt.wantASCII)

			buf.Reset()

			err = maze_writer.WriteMaze(&buf, tt.m, true)
			require.NoError(t, err)

			gotUnicode := buf.String()
			assert.Equal(t, gotUnicode, tt.wantUnicode)
		})
	}
}

func TestWriteMazeWithSolutionASCII(t *testing.T) {
	maze2x2, path2x2 := create2x2Maze()
	maze3x3, path3x3 := create3x3Maze()

	tests := []struct {
		name        string
		m           *maze.Maze
		path        entities.Path
		wantASCII   string
		wantUnicode string
	}{
		{
			name: "2x2 maze with valid path",
			m:    maze2x2,
			path: path2x2,
			wantASCII: strings.TrimPrefix(`
#####
#O#X#
#.#.#
#...#
#####
`, "\n"),
			wantUnicode: strings.TrimPrefix(`
┌─┬─┐
│O│X│
│.╵.│
│...│
└───┘
`, "\n"),
		},
		{
			name: "3x3 maze with valid path",
			m:    maze3x3,
			path: path3x3,
			wantASCII: strings.TrimPrefix(`
#######
#O#...#
#.#.#.#
#.#X#.#
#.###.#
#.....#
#######
`, "\n"),
			wantUnicode: strings.TrimPrefix(`
┌─┬───┐
│O│...│
│.│.╷.│
│.│X│.│
│.└─┘.│
│.....│
└─────┘
`, "\n"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer

			err := maze_writer.WriteMazeWithSolution(&buf, tt.m, tt.path, false)
			require.NoError(t, err)

			gotASCII := buf.String()
			assert.Equal(t, gotASCII, tt.wantASCII)

			buf.Reset()

			err = maze_writer.WriteMazeWithSolution(&buf, tt.m, tt.path, true)
			require.NoError(t, err)

			gotUnicode := buf.String()
			assert.Equal(t, gotUnicode, tt.wantUnicode)
		})
	}
}

func create3x3Maze() (*maze.Maze, entities.Path) {
	/*
		[a]   [b] - [c]
		 |	   |	 |
		[d]   [e]   [f]
		 |			 |
		[g] - [h] - [i]
	*/

	m := maze.New(3, 3)
	a := m.Cell(0, 0)
	b := m.Cell(0, 1)
	c := m.Cell(0, 2)
	d := m.Cell(1, 0)
	e := m.Cell(1, 1)
	f := m.Cell(1, 2)
	g := m.Cell(2, 0)
	h := m.Cell(2, 1)
	i := m.Cell(2, 2)

	path := []*entities.Cell{a, d, g, h, i, f, c, b, e}

	entities.DirectionDown.RemoveWall(a, d)
	entities.DirectionDown.RemoveWall(d, g)
	entities.DirectionRight.RemoveWall(g, h)
	entities.DirectionRight.RemoveWall(h, i)
	entities.DirectionUp.RemoveWall(i, f)
	entities.DirectionUp.RemoveWall(f, c)
	entities.DirectionLeft.RemoveWall(c, b)
	entities.DirectionDown.RemoveWall(b, e)

	return m, entities.NewPath(path)
}

func create2x2Maze() (*maze.Maze, entities.Path) {
	/*
		[a]   [b]
		 |	   |
		[c] - [d]
	*/

	m := maze.New(2, 2)
	a := m.Cell(0, 0)
	b := m.Cell(0, 1)
	c := m.Cell(1, 0)
	d := m.Cell(1, 1)

	path := []*entities.Cell{a, c, d, b}

	entities.DirectionDown.RemoveWall(a, c)
	entities.DirectionRight.RemoveWall(c, d)
	entities.DirectionUp.RemoveWall(d, b)

	return m, entities.NewPath(path)
}
