package sorting

func qs(arr []int, low, high int) {
	if low >= high {
		return
	}

	pivotIdx := partition(arr, low, high)

	qs(arr, pivotIdx+1, high)
	qs(arr, low, pivotIdx-1)
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	idx := low - 1
	for i := low; i < high; i++ {
		if arr[i] <= pivot {
			idx++
			arr[i], arr[idx] = arr[idx], arr[i]
		}
	}

	idx++
	arr[high], arr[idx] = arr[idx], pivot
	return idx
}

func quick_sort(arr []int) []int {
	qs(arr, 0, len(arr)-1)
	return arr
}
