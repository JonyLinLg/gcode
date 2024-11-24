package sort

import "testing"

func TestInputCompare(t *testing.T) {
	RandomInputCompare([]SortType{Selection, Insertion}, 100000)
	t.Log("random input over")
	SortedInputCompare([]SortType{Selection, Insertion}, 100000)
	t.Log("sorted input over")
	Sorted2InputCompare([]SortType{Selection, Insertion}, 100000)
	t.Log("sorted2 input over")
}
