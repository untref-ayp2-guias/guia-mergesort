package ejercicios

func RecursiveHeapSort[T any](arr []T, compare func(T, T) int) {
	size := len(arr)
	heapify(arr, compare)

	end := size - 1
	for end > 0 {
		arr[end], arr[0] = arr[0], arr[end]
		//downHeap
		end--
	}
}

func heapify[T any](arr []T, compare func(T, T) int) {
	size := len(arr)
	start := (size - 2) / 2

	for start >= 0 {
		//downHeap
		start--
	}
}
