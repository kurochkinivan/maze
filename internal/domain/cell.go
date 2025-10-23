package domain

type Cell struct {
	Point
	Walls
}

type Walls struct {
	Top, Right, Bottom, Left bool
}

type Point struct {
	X, Y int
}
