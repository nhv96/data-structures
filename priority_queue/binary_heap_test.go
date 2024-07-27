package priorityqueue

import (
	"testing"
)

func Test_BHeap(t *testing.T) {
	t.Parallel()

	// WIP to be updated
	t.Run("init heap", func(t *testing.T) {
		arr := []int{5, 4, 7, 1, 3, 5}

		bheap := NewBHeap(arr)

		expectedArr := []int{1, 3, 5, 4, 5, 7}

		if bheap.Size() != len(expectedArr) {
			t.Errorf("Heap size is expected to be %v, got %v", len(expectedArr), bheap.Size())
		}

		for i := 0; i < len(expectedArr); i++ {
			val, _ := bheap.Get(i)
			if val != expectedArr[i] {
				t.Errorf("Element at index %d is expected to be %v, got %v", i, expectedArr[i], val)
			}
		}
	})
}
