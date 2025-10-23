package domain

type RemoveWall func(c1, c2 *Cell)

type Move struct {
	*Cell
	RemoveWall
}

type Direction struct {
	RemoveWall
	DX, DY int
}

var directions = []Direction{
	{DX: 1, DY: 0, RemoveWall: func(a, b *Cell) { a.Bottom, b.Top = false, false }},
	{DX: -1, DY: 0, RemoveWall: func(a, b *Cell) { a.Top, b.Bottom = false, false }},
	{DX: 0, DY: 1, RemoveWall: func(a, b *Cell) { a.Right, b.Left = false, false }},
	{DX: 0, DY: -1, RemoveWall: func(a, b *Cell) { a.Left, b.Right = false, false }},
}
