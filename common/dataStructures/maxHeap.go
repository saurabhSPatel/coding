package common

import (
	"container/heap"
	"fmt"
)

// Define a custom type for the max heap
type MaxHeap []int

// Implement the heap.Interface for MaxHeap

// Len returns the number of elements in the heap
func (h MaxHeap) Len() int {
	return len(h)
}

// Less defines the ordering of elements (for max heap, it returns true if i > j)
func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j] // Change to < for min heap
}

// Swap swaps two elements in the heap
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push adds an element to the heap
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int)) // Append the new element
}

// Pop removes and returns the top element from the heap
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]     // Get the last element
	*h = old[0 : n-1] // Remove the last element
	return x
}

func main() {
	// Create a max heap
	h := &MaxHeap{}

	// Initialize the heap
	heap.Init(h)

	// Push elements into the heap
	heap.Push(h, 3)
	heap.Push(h, 2)
	heap.Push(h, 5)
	heap.Push(h, 1)
	heap.Push(h, 4)

	fmt.Println("Max Heap:", *h) // Output: [5 4 2 1 3]

	// Pop elements from the heap (in descending order)
	fmt.Println("Popping elements from the max heap:")
	for h.Len() > 0 {
		fmt.Println(heap.Pop(h))
	}
}
