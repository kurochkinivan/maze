package generator

type Algorithm string

const (
	AlgoDFS  Algorithm = "dfs"
	AlgoPrim Algorithm = "prim"
)

func (a Algorithm) IsValid() bool {
	switch a {
	case AlgoDFS, AlgoPrim:
		return true
	default:
		return false
	}
}
