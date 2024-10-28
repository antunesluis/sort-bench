package core

type SortMode string
type SortAlgo string

const (
	ModeRecursive SortMode = "recursive"
	ModeIterative SortMode = "iterative"
	ModeParallel  SortMode = "parallel"
)

const (
	AlgoMergeSort SortAlgo = "mergesort"
	AlgoQuickSort SortAlgo = "quicksort"
	AlgoHeapsort  SortAlgo = "heapsort"
)
