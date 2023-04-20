package helper

// A priority queue stores the element in a way that higher priority elements are at the top.
// It uses a heap underhood, therefore, to use PQ:
//
//	heap.Init(&pq)
//	heap.Push(&pq, 5)
type PQ struct {
	list []interface{}

	// To compare between the elements.
	// To make a max heap:
	//
	//	pq.Compare = func(list []interface{}, i int, j int) bool {
	//		return list[i].(int) > list[j].(int)
	//	}
	Compare func(list []interface{}, i int, j int) bool
}

// Get the length of the queue
func (q PQ) Len() int {
	return len(q.list)
}

func (q PQ) Less(i int, j int) bool {
	return q.Compare(q.list, i, j)
}

func (q PQ) Swap(i int, j int) {
	q.list[i], q.list[j] = q.list[j], q.list[i]
}

// Push adds an element to the priority queue.
//
// Item of higher priority goes to the front.
//
//	heap.Push(&pq, 2)
//
// Complexity: O(log(n))
func (h *PQ) Push(x interface{}) {
	h.list = append(h.list, x)
}

// Remove an element to the queue.
// This should be done through the container/heap package like so:
//
//	removed := heap.Pop(&h)
//
// Complexity: O(log n)
func (h *PQ) Pop() interface{} {
	old := h.list
	n := len(old)
	x := old[n-1]
	h.list = old[:n-1]
	return x
}
