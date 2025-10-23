package bag

import "math/rand/v2"

// Bag is a data structure optimized for O(1) insertion, random access,
// and deletion without preserving element order. Useful for randomized
// algorithms where order does not matter.
type Bag[T any] struct {
	items []T
}

// New creates an empty Bag with an optional initial capacity.
func New[T any](capacity int) *Bag[T] {
	return &Bag[T]{
		items: make([]T, 0, capacity),
	}
}

// Add inserts an item into the bag in O(1) time.
func (r *Bag[T]) Add(item T) {
	r.items = append(r.items, item)
}

// RemoveAt removes an element by index in O(1) time.
// The order of elements is not preserved.
func (r *Bag[T]) RemoveAt(idx int) {
	if idx < 0 || r.Len() <= idx {
		return
	}

	last := r.Len() - 1
	r.items[idx], r.items[last] = r.items[last], r.items[idx]
	r.items = r.items[:last]
}

// RandomItemAndDelete returns a random item and removes it from the bag in O(1).
// If the bag is empty, the zero value of T is returned.
func (r *Bag[T]) RandomItemAndDelete() T {
	if r.IsEmpty() {
		var zero T
		return zero
	}

	idx := rand.IntN(r.Len())
	item := r.items[idx]

	r.RemoveAt(idx)
	return item
}

// Len returns the number of elements currently stored in the bag.
func (r *Bag[T]) Len() int {
	return len(r.items)
}

// IsEmpty reports whether the bag contains no elements.
func (r *Bag[T]) IsEmpty() bool {
	return r.Len() == 0
}
