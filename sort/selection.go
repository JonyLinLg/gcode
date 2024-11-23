package sort

import "cmp"

func SelectionSort[T cmp.Ordered](arr []T) []T {
	for i := 0; i < len(arr)-1; i++ {
		index := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[index] {
				index = j
			}
		}
		if index != i {
			arr[i], arr[index] = arr[index], arr[i]
		}
	}
	return arr
}
