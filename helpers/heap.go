package helpers

import (
	"cmp"
)

type Heap[T cmp.Ordered] []T

func (h Heap[T]) Len() int           { return len(h) }
func (h Heap[T]) Less(i, j int) bool { return h[i] < h[j] }
func (h Heap[T]) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap[T]) Push(x any) {
	*h = append(*h, x.(T))
}

func (h *Heap[T]) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
