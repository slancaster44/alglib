package alglib

import "alglib/generics"

type Heap[T any] struct {
	data          []T
	comparator    func(T, T) bool
	eq_comparator func(T, T) bool
}

func NewMaxHeap[T generics.Sortable]() *Heap[T] {
	return NewHeapWithComparator(generics.ComparatorGreaterThan[T], generics.ComparatorEquals[T])
}

func NewMinHeap[T generics.Sortable]() *Heap[T] {
	return NewHeapWithComparator(generics.ComparatorLessThan[T], generics.ComparatorEquals[T])
}

func NewHeapWithComparator[T any](heapsort_comparator, eq_comparator func(T, T) bool) *Heap[T] {
	return &Heap[T]{comparator: heapsort_comparator, eq_comparator: eq_comparator}
}

func (h *Heap[T]) GetValue(idx int) T {
	return h.data[idx]
}

func (h *Heap[T]) Parent(idx int) int {
	return idx / 2
}

func (h *Heap[T]) LeftChild(idx int) int {
	return idx * 2
}

func (h *Heap[T]) RightChild(idx int) int {
	return idx*2 + 1
}

func (h *Heap[T]) HeapCompareIndices(i1, i2 int) bool {
	return h.comparator(h.GetValue(i1), h.GetValue(i2))
}

func (h *Heap[T]) CompareIndicies(i1, i2 int) bool {
	return h.eq_comparator(h.GetValue(1), h.GetValue(i2))
}

func (h *Heap[T]) IsOutOfBounds(idx int) bool {
	return idx >= len(h.data)
}

func (h *Heap[T]) Insert(val T) {
	h.data = append([]T{val}, h.data...) /* TODO: Seems inefficient. If memory copy, the insert becomes O(n) */
	h.heapify(0)
}

func (h *Heap[T]) swap(i1, i2 int) {
	tmp := h.data[i1]
	h.data[i1] = h.data[i2]
	h.data[i2] = tmp
}

func (h *Heap[T]) heapify(idx int) {
	l_idx := h.LeftChild(idx)
	r_idx := h.RightChild(idx)

	lightest := idx

	if !h.IsOutOfBounds(l_idx) && h.HeapCompareIndices(l_idx, idx) {
		lightest = l_idx
	} else if !h.IsOutOfBounds(r_idx) && h.HeapCompareIndices(r_idx, idx) {
		lightest = r_idx
	}

	if lightest != idx {
		h.swap(idx, lightest) /* Let lighter item float up the heap */
		h.heapify(lightest)   /* Ensure that everything below here still maintains the heap property */
	}
}

func (h *Heap[T]) isHeap() bool {
	for i := 1; i < len(h.data); i++ {
		if !h.HeapCompareIndices(h.Parent(i), i) && h.CompareIndicies(h.Parent(i), i) {
			return false
		}
	}

	return true
}
