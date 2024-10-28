package sorting

import (
	"sort-bench/internal/core"
)

type BaseSorter struct {
	name    string
	mode    core.SortMode
	metrics core.Metrics
}

func NewBaseSorter(name string, mode core.SortMode) BaseSorter {
	return BaseSorter{
		name: name,
		mode: mode,
	}
}

func (b *BaseSorter) Name() string             { return b.name }
func (b *BaseSorter) Mode() core.SortMode      { return b.mode }
func (b *BaseSorter) GetMetrics() core.Metrics { return b.metrics }

func (b *BaseSorter) compare(x, y int) bool {
	b.metrics.Comparisons++
	return x < y
}

func (b *BaseSorter) swap(arr []int, i, j int) {
	b.metrics.Swaps++
	arr[i], arr[j] = arr[j], arr[i]
}
