package helper

type MaxHeap []int

// Returns the length of the heap
func (h MaxHeap) Len() int {
	return len(h)
}

// Compares 2 heap elements
func (h MaxHeap) Less(i int, j int) bool {
	return h[i] > h[j]
}

// Swap value of 2 heap elements
func (h MaxHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

// Pushes an element to the heap
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Remove the last element from the heap
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
