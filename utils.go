package helper

// Get the maximum element in the array
func MaxInArray(arr []interface{}, larger Comparator) interface{} {
	if len(arr) == 0 {
		return nil
	}
	max := arr[0]
	for _, e := range arr {
		if larger(e, max) {
			max = e
		}
	}
	return max
}

// Get the minimum element in the array
func MinInArray(arr []interface{}, larger Comparator) interface{} {
	if len(arr) == 0 {
		return nil
	}
	min := arr[0]
	for _, e := range arr {
		if larger(min, e) {
			min = e
		}
	}
	return min
}
