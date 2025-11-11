package maze_writer

var (
	horizontal        rune = '─'
	vertical          rune = '│'
	topRightCorner    rune = '┐'
	topLeftCorner     rune = '┌'
	bottomLeftCorner  rune = '└'
	bottomRightCorner rune = '┘'
	cross             rune = '┼'
	tUp               rune = '┴'
	tDown             rune = '┬'
	tLeft             rune = '┤'
	tRight            rune = '├'
	smallWallUp       rune = '╷'
	smallWallDown     rune = '╵'
	smallWallLeft     rune = '╴'
	smallWallRight    rune = '╶'
)

// renderGridUnicode converts the ASCII maze grid into a Unicode version with styled wall characters.
func renderGridUnicode(gridASCII [][]rune) [][]rune {
	rows, cols := len(gridASCII), len(gridASCII[0])

	gridUnicode := make([][]rune, rows)
	for i := range gridUnicode {
		gridUnicode[i] = make([]rune, cols)
	}

	for row := range rows {
		for col := range cols {
			if gridASCII[row][col] == ' ' {
				gridUnicode[row][col] = ' '
				continue
			}

			hasLeft := isWall(gridASCII, row, col-1)
			hasRight := isWall(gridASCII, row, col+1)
			hasTop := isWall(gridASCII, row-1, col)
			hasBottom := isWall(gridASCII, row+1, col)

			r := cellUnicode(hasLeft, hasRight, hasTop, hasBottom)
			gridUnicode[row][col] = r
		}
	}

	return gridUnicode
}

// cellUnicode selects the appropriate Unicode wall character based on neighboring walls.
func cellUnicode(hasLeft, hasRight, hasTop, hasBottom bool) rune {
	switch {
	// cross
	case hasLeft && hasRight && hasTop && hasBottom:
		return cross

	// t-shaped
	case !hasLeft && hasRight && hasTop && hasBottom:
		return tRight
	case hasLeft && !hasRight && hasTop && hasBottom:
		return tLeft
	case hasLeft && hasRight && !hasTop && hasBottom:
		return tDown
	case hasLeft && hasRight && hasTop && !hasBottom:
		return tUp

	// horizontal and vertical
	case hasLeft && hasRight && !hasTop && !hasBottom:
		return horizontal
	case !hasLeft && !hasRight && hasTop && hasBottom:
		return vertical

	// small shapes
	case hasLeft && !hasRight && !hasTop && !hasBottom:
		return smallWallLeft
	case !hasLeft && hasRight && !hasTop && !hasBottom:
		return smallWallRight
	case !hasLeft && !hasRight && hasTop && !hasBottom:
		return smallWallDown
	case !hasLeft && !hasRight && !hasTop && hasBottom:
		return smallWallUp

	// corners
	case !hasLeft && hasRight && !hasTop && hasBottom:
		return topLeftCorner
	case hasLeft && !hasRight && !hasTop && hasBottom:
		return topRightCorner
	case !hasLeft && hasRight && hasTop && !hasBottom:
		return bottomLeftCorner
	case hasLeft && !hasRight && hasTop && !hasBottom:
		return bottomRightCorner

	default:
		return 'U'
	}
}

// isWall checks if the given cell in the grid is a wall.
func isWall(gridASCII [][]rune, row, col int) bool {
	return isValid(gridASCII, row, col) && gridASCII[row][col] == '#'
}

// isValid checks if the given coordinates are within the grid bounds.
func isValid(gridASCII [][]rune, row, col int) bool {
	rows, cols := len(gridASCII), len(gridASCII[0])
	return 0 <= row && row < rows && 0 <= col && col < cols
}
