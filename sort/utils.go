package sort

import (
	"cmp"
	"math/rand"
)

// randomIntArray generates a random array of integers
func randomIntArray(n int) (arr []int) {
	arr = make([]int, n)
	for i := range arr {
		arr[i] = rand.Intn(randomRange)
	}
	return arr
}

// isSorted checks if an array is sorted or not
func isSorted[T cmp.Ordered](arr []T) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}
