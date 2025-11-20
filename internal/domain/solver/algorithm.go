package solver

type Algorithm string

const (
	AlgoAStar    Algorithm = "astar"
	AlgoDijkstra Algorithm = "dijkstra"
)

func (a Algorithm) IsValid() bool {
	switch a {
	case AlgoAStar, AlgoDijkstra:
		return true
	default:
		return false
	}
}
