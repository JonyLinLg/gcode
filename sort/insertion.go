package sort

import "cmp"

func InsertionSort[T cmp.Ordered](arr []T) []T {
	for i := 1; i < len(arr); i++ {
		index := i
		for j := i; j > 0; j-- {
			if arr[j] >= arr[j-1] {
				break
			}
			arr[j] = arr[j-1]
			index = j - 1
		}
		if index != i {
			arr[index] = arr[i]
		}
	}
	return arr
}
