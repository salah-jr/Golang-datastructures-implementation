package main

import "fmt"

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// return the largest key, and removes it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	lastIndex := len(h.array) - 1

	if len(h.array) == 0 {
		fmt.Println("Can't extract because array length is 0")
		return -1
	}
	h.array[0] = h.array[lastIndex]
	h.array = h.array[:lastIndex]
	h.maxHeapifyDown(0)
	return extracted
}

//heapify from bottom to top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

//heapify from top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}
	}

	if h.array[index] < h.array[childToCompare] {
		h.swap(index, childToCompare)
		index = childToCompare
		l, r = left(index), right(index)
	} else {
		return
	}
}

//get the parent index
func parent(i int) int {
	return (i - 1) / 2
}

//get the left child
func left(i int) int {
	return 2*i + 1
}

//get the right chikd
func right(i int) int {
	return 2*i + 2
}

//swap keys in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}
func main() {
	m := &MaxHeap{}
	list := []int{
		10, 20, 30, 40, 50, 5, 17, 67, 77,
	}
	for _, v := range list {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
