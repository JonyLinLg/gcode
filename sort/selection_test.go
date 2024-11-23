package sort

import "testing"

func BenchmarkSelectionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := SelectionSort(randomIntArray(arraySize))
		if !isSorted(res) {
			b.Fatal("not sorted")
		}
	}
}
