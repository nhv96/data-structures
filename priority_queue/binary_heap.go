package priorityqueue

import (
	"fmt"
	"math"
)

// BHeap or Binary Heap is the implementation of min priority queue
type BHeap struct {
	heapSize int
	heap     []int
}

// NewBHeap returns a heap with all the elements added at the correct position
func NewBHeap(elements []int) *BHeap {
	h := &BHeap{}
	for _, e := range elements {
		h.heap = append(h.heap, e)
	}
	h.heapSize = len(elements)

	for i := math.Max(0, (float64(h.heapSize)/2)-1); i >= 0; i-- {
		h.sink(int(i))
	}
	return h
}

// Size returns heap's size
func (h *BHeap) Size() int {
	return h.heapSize
}

// Get returns the value of element at k index
func (h *BHeap) Get(k int) (int, error) {
	if k >= h.Size() {
		return 0, fmt.Errorf("Invalid index %d", k)
	}
	return h.heap[k], nil
}

func (h *BHeap) sink(k int) {
	for {
		left := 2*k + 1  // position of left node
		right := 2*k + 2 // position of right node

		smallest := left // assume the left node is the smallest

		if right < h.heapSize && h.heap[right] < h.heap[left] {
			smallest = right
		}

		if left >= h.heapSize || h.heap[k] < h.heap[smallest] {
			break
		}

		h.swap(k, smallest)
		k = smallest
	}
}

// swap the value of two position
func (h *BHeap) swap(i, j int) {
	tmp := h.heap[j]
	h.heap[j] = h.heap[i]
	h.heap[i] = tmp
}

func (h *BHeap) swim(k int) {
	parent := (k - 1) / 2
	for k > 0 && h.heap[parent] > h.heap[k] {
		h.swap(parent, k)
		k = parent
		parent = (k - 1) / 2
	}
}

// Add adds an element to the heap
func (h *BHeap) Add(v int) {
	h.heap = append(h.heap, v)
	h.heapSize = len(h.heap)
	h.swim(h.heapSize - 1)
}
