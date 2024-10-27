package sorting

import "time"

// SortMode define os modos de ordenação
type SortMode string

const (
	ModeRecursive SortMode = "recursive"
	ModeIterative SortMode = "iterative"
)

// Metrics armazena métricas básicas de performance
type Metrics struct {
	Comparisons int
	Swaps       int
	Time        time.Duration
}

// Sorter interface básica para algoritmos de ordenação
type Sorter interface {
	Sort([]int) []int
	Name() string
	Mode() SortMode
	GetMetrics() Metrics
}

// BaseSorter implementação base simplificada
type BaseSorter struct {
	name    string
	mode    SortMode
	metrics Metrics
}

func NewBaseSorter(name string, mode SortMode) BaseSorter {
	return BaseSorter{
		name:    name,
		mode:    mode,
		metrics: Metrics{},
	}
}

func (b *BaseSorter) Name() string        { return b.name }
func (b *BaseSorter) Mode() SortMode      { return b.mode }
func (b *BaseSorter) GetMetrics() Metrics { return b.metrics }

// Métodos utilitários simplificados
func (b *BaseSorter) compare(x, y int) bool {
	b.metrics.Comparisons++
	return x < y
}

func (b *BaseSorter) swap(arr []int, i, j int) {
	b.metrics.Swaps++
	arr[i], arr[j] = arr[j], arr[i]
}

