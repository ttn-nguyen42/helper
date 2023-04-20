package helper

// If a is larger than b, returns true.
//
// Otherwise, returns false.
type Comparator = func(a interface{}, b interface{}) bool

func QuickSort(arr []interface{}, larger Comparator) {
	if len(arr) < 2 {
		return
	}
	doQuickSort(arr, 0, len(arr)-1, larger)
}

func doQuickSort(arr []interface{}, low int, high int, larger Comparator) {
	if low < high {
		pivot := partition(arr, low, high, larger)
		doQuickSort(arr, low, pivot-1, larger)
		doQuickSort(arr, pivot+1, high, larger)
	}
}

func partition(arr []interface{}, low int, high int, larger Comparator) int {
	pivot := arr[high]
	i := low - 1
	// Move all smaller element to the left
	for j := low; j < high; j += 1 {
		// Pointer j finds the next smallest element
		if larger(pivot, arr[j]) {
			// Pointer i keeps the current smallest element
			i += 1
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	// Pointer i at this point is at the index of the last smallest element
	// Swap the next one with pivot to put pivot in the middle
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
