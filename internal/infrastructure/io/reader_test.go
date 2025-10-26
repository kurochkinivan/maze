package io

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadMaze_HappyPath(t *testing.T) {
	input := `
#######
#   # #
# ### #
#     #
#######
`

	r := strings.NewReader(strings.TrimSpace(input))

	m, err := ReadMaze(r)
	require.NoError(t, err)

	assert.Equal(t, 2, m.Rows())
	assert.Equal(t, 3, m.Cols())

	leftTop := m.Cell(0, 0)
	assert.True(t, leftTop.Top)
	assert.True(t, leftTop.Left)
	assert.False(t, leftTop.Right)
	assert.False(t, leftTop.Bottom)

	rightTop := m.Cell(0, 2)
	assert.True(t, rightTop.Top)
	assert.True(t, rightTop.Left)
	assert.True(t, rightTop.Right)
	assert.False(t, rightTop.Bottom)

	centalDown := m.Cell(1, 1)
	assert.True(t, centalDown.Top)
	assert.True(t, centalDown.Bottom)
	assert.False(t, centalDown.Left)
	assert.False(t, centalDown.Right)
}

func TestReadMaze_OneCellMaze(t *testing.T) {
	input := `
###
# #
###
`

	r := strings.NewReader(strings.TrimSpace(input))

	m, err := ReadMaze(r)
	require.NoError(t, err)

	assert.Equal(t, 1, m.Size())

	cell := m.Cell(0, 0)
	assert.True(t, cell.Top)
	assert.True(t, cell.Bottom)
	assert.True(t, cell.Left)
	assert.True(t, cell.Right)
}
func TestReadMaze_Errors(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "empty reader",
			input: ``,
		},
		{
			name: "even number of colums",
			input: `
####
#  #
####
`,
		},
		{
			name: "even number of rows",
			input: `
###
# #
# #
###
`,
		},
		{
			name: "not enough rows",
			input: `
###
###
`,
		},
		{
			name: "not enough columns",
			input: `
##
##
##
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(strings.TrimSpace(tt.input))
			m, err := ReadMaze(r)

			require.Error(t, err, "maze must be invalid")
			assert.Nil(t, m)
		})
	}
}
