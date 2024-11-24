package sort

import (
	"cmp"
)

func InsertionSort[T cmp.Ordered](arr []T) []T {
	for i := 1; i < len(arr); i++ {
		index := i
		num := arr[i]
		for j := i - 1; j >= 0; j-- {
			if arr[j] < num {
				break
			}
			arr[j+1] = arr[j]
			index = j
		}
		if index != i {
			arr[index] = num
		}
	}
	return arr
}
