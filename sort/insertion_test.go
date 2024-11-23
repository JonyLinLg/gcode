package sort

import "testing"

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := InsertionSort(randomIntArray(arraySize))
		if !isSorted(res) {
			b.Fatal("not sorted")
		}
	}
}
