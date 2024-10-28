package core

type Sorter interface {
	Sort([]int) []int
	Name() string
	Mode() SortMode
	GetMetrics() Metrics
}
