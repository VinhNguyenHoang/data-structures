package priorityqueue

import (
	"testing"
)

func Test_BHeap(t *testing.T) {
	t.Parallel()

	t.Run("init heap", func(t *testing.T) {
		arr := []int{5, 4, 7, 1, 3, 5}
		exp := []int{1, 3, 5, 4, 5, 7}
		bh := NewBHeap(arr)

		cmpArr(t, bh, exp)
	})

	t.Run("init heap with data already heapified", func(t *testing.T) {
		arr := []int{1, 3, 5, 4, 5, 7}
		exp := []int{1, 3, 5, 4, 5, 7}
		bh := NewBHeap(arr)

		cmpArr(t, bh, exp)
	})

	t.Run("add an element", func(t *testing.T) {
		arr := []int{5, 6, 8, 7, 19, 11, 12, 13, 14, 12}
		exp := []int{1, 5, 8, 7, 6, 11, 12, 13, 14, 19, 12}
		bh := NewBHeap(arr)
		bh.Add(1)

		cmpArr(t, bh, exp)
	})

	t.Run("poll the min element", func(t *testing.T) {
		arr := []int{5, 4, 7, 1, 3, 5}

		bh := NewBHeap(arr)

		expPolled := []int{1, 3, 4, 5, 5, 7}

		for i := 0; i < len(expPolled); i++ {
			polled := bh.Poll()

			if polled != expPolled[i] {
				t.Errorf("Expected polled value to be %d, got %d", expPolled[i], polled)
			}
		}
	})

	t.Run("remove a value", func(t *testing.T) {
		arr := []int{5, 6, 8, 7, 19, 11, 12, 13, 14, 12}

		bh := NewBHeap(arr)

		removeVal := 13
		bh.remove(removeVal)

		// check if element still exist in the heap
		for _, v := range bh.heap {
			if v == removeVal {
				t.Errorf("Expected value %d to be removed", removeVal)
			}
		}
	})
}

func cmpArr(t *testing.T, h *BHeap, exp []int) {
	if h.Size() != len(exp) {
		t.Errorf("Heap size is expected to be %v, got %v", len(exp), h.Size())
	}

	for i := 0; i < len(exp); i++ {
		val, _ := h.Get(i)
		if val != exp[i] {
			t.Errorf("Element at index %d is expected to be %v, got %v", i, exp[i], val)
		}
	}
}
