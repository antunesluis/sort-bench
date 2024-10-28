package core

type SortMode string

const (
	ModeRecursive SortMode = "recursive"
	ModeIterative SortMode = "iterative"
	ModeParallel  SortMode = "parallel"
)
