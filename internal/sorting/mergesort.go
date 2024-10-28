package sorting

import (
	"sort-bench/internal/core"
	"time"
)

type MergeSort struct {
	BaseSorter
}

func NewMergeSort(mode core.SortMode) *MergeSort {
	return &MergeSort{
		BaseSorter: NewBaseSorter("MergeSort", mode),
	}
}

func (m *MergeSort) Sort(arr []int) []int {
	start := time.Now()
	m.metrics = core.Metrics{} // Inicializa as métricas

	result := arr
	switch m.Mode() {
	case core.ModeRecursive:
		m.sortRecursive(result, 0, len(result)-1)
	case core.ModeIterative:
		m.sortIterative(result)
	}

	m.metrics.Time = time.Since(start)
	return result
}

func (m *MergeSort) sortRecursive(arr []int, left, right int) {
	if left < right {
		mid := left + (right-left)/2
		m.sortRecursive(arr, left, mid)
		m.sortRecursive(arr, mid+1, right)
		m.merge(arr, left, mid, right)
	}
}

func (m *MergeSort) sortIterative(arr []int) {
	n := len(arr)
	for size := 1; size < n; size *= 2 {
		for left := 0; left < n-1; left += 2 * size {
			mid := min(left+size-1, n-1)
			right := min(left+2*size-1, n-1)
			m.merge(arr, left, mid, right)
		}
	}
}

func (m *MergeSort) merge(arr []int, left, mid, right int) {
	temp := make([]int, right-left+1)
	i, j, k := left, mid+1, 0

	// Comparações entre elementos de duas metades
	for i <= mid && j <= right {
		m.metrics.Comparisons++ // Incrementa comparações
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
		k++
		m.metrics.Swaps++ // Cada inserção em 'temp' é considerada uma "troca"
	}

	// Copia os elementos restantes da primeira metade
	for i <= mid {
		temp[k] = arr[i]
		i++
		k++
		m.metrics.Swaps++ // Cada inserção em 'temp' é uma "troca"
	}

	// Copia os elementos restantes da segunda metade
	for j <= right {
		temp[k] = arr[j]
		j++
		k++
		m.metrics.Swaps++ // Cada inserção em 'temp' é uma "troca"
	}

	// Copia de volta para o array original
	copy(arr[left:right+1], temp)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

