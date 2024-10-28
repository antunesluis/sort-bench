package sorting

import (
	"sort-bench/internal/core"
	"sync"
	"sync/atomic"
	"time"
)

type MergeSort struct {
	BaseSorter
	mu sync.Mutex // Mutex para sincronização das métricas
}

// SafeMetrics é uma estrutura thread-safe para métricas
type SafeMetrics struct {
	Comparisons atomic.Int64
	Swaps       atomic.Int64
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

	// Cria métricas seguras para threads
	safeMetrics := &SafeMetrics{}

	switch m.Mode() {
	case core.ModeRecursive:
		m.sortRecursive(result, 0, len(result)-1)
	case core.ModeIterative:
		m.sortIterative(result)
	case core.ModeParallel:
		m.sortParallel(result, 0, len(result)-1, safeMetrics)
		// Atualiza as métricas finais
		m.metrics.Comparisons = int(safeMetrics.Comparisons.Load())
		m.metrics.Swaps = int(safeMetrics.Swaps.Load())
	}

	m.metrics.Time = time.Since(start)
	return result
}

func (m *MergeSort) sortParallel(arr []int, left, right int, metrics *SafeMetrics) {
	threshold := 1000

	if left < right {
		if right-left < threshold {
			// Use a versão sequencial para arrays pequenos
			m.sortRecursiveWithMetrics(arr, left, right, metrics)
			return
		}

		mid := left + (right-left)/2

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			m.sortParallel(arr, left, mid, metrics)
		}()

		go func() {
			defer wg.Done()
			m.sortParallel(arr, mid+1, right, metrics)
		}()

		wg.Wait()

		m.mergeWithMetrics(arr, left, mid, right, metrics)
	}
}

func (m *MergeSort) sortRecursiveWithMetrics(arr []int, left, right int, metrics *SafeMetrics) {
	if left < right {
		mid := left + (right-left)/2
		m.sortRecursiveWithMetrics(arr, left, mid, metrics)
		m.sortRecursiveWithMetrics(arr, mid+1, right, metrics)
		m.mergeWithMetrics(arr, left, mid, right, metrics)
	}
}

func (m *MergeSort) mergeWithMetrics(arr []int, left, mid, right int, metrics *SafeMetrics) {
	temp := make([]int, right-left+1)
	i, j, k := left, mid+1, 0

	for i <= mid && j <= right {
		metrics.Comparisons.Add(1) // Incrementa comparações de forma atômica
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
		k++
		metrics.Swaps.Add(1) // Incrementa swaps de forma atômica
	}

	for i <= mid {
		temp[k] = arr[i]
		i++
		k++
		metrics.Swaps.Add(1)
	}

	for j <= right {
		temp[k] = arr[j]
		j++
		k++
		metrics.Swaps.Add(1)
	}

	copy(arr[left:right+1], temp)
}

// Mantendo os métodos originais para os modos não paralelos
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

	for i <= mid && j <= right {
		m.metrics.Comparisons++
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
		k++
		m.metrics.Swaps++
	}

	for i <= mid {
		temp[k] = arr[i]
		i++
		k++
		m.metrics.Swaps++
	}

	for j <= right {
		temp[k] = arr[j]
		j++
		k++
		m.metrics.Swaps++
	}

	copy(arr[left:right+1], temp)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

