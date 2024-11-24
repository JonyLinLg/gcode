package sort

import "math"

const (
	arraySize   = 10000
	randomRange = math.MaxInt32
)

type SortType int

const (
	Selection SortType = iota // 选择排序
	Insertion                 // 插入排序
)
