package sort

import (
	"cmp"
	"fmt"
	"slices"
	"time"
)

func RandomInputCompare(st []SortType, al int) {
	arr := randomIntArray(al)
	for i := 0; i < len(st); i++ {
		runSortFun(st[i], slices.Clone(arr))
	}
}

func SortedInputCompare(st []SortType, al int) {
	arr := randomIntArray(al)
	InsertionSort(arr)
	for i := 0; i < len(st); i++ {
		runSortFun(st[i], slices.Clone(arr))
	}
}

func Sorted2InputCompare(st []SortType, al int) {
	arr := randomIntArray(al)
	InsertionSort(arr)
	for i := 0; i < len(arr)-i; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	for i := 0; i < len(st); i++ {
		runSortFun(st[i], slices.Clone(arr))
	}
}

func runSortFun[T cmp.Ordered](st SortType, arr []T) {
	now := time.Now()
	switch st {
	case Selection:
		SelectionSort(arr)
	case Insertion:
		InsertionSort(arr)
	}
	fmt.Println(st, time.Since(now))
}
