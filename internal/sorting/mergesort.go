package sorting

import (
	"sort-bench/internal/core"
	"sync"
	"sync/atomic"
	"time"
)

type MergeSort struct {
	BaseSorter
	mu sync.Mutex
}

type SafeMetrics struct {
	Comparisons atomic.Int64
	Swaps       atomic.Int64
}

func NewMergeSort(mode core.SortMode) *MergeSort {
	// Validação do modo de operação
	switch mode {
	case core.ModeRecursive, core.ModeIterative, core.ModeParallel:
		// Modos válidos
	default:
		// Se o modo for inválido, usa o modo recursivo como fallback
		mode = core.ModeRecursive
	}

	return &MergeSort{
		BaseSorter: NewBaseSorter("MergeSort", mode),
	}
}

func (m *MergeSort) Sort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	start := time.Now()
	m.metrics = core.Metrics{}

	// Cria uma cópia do array original
	result := make([]int, len(arr))
	copy(result, arr)

	safeMetrics := &SafeMetrics{}

	switch m.Mode() {
	case core.ModeRecursive:
		m.sortRecursive(result, 0, len(result)-1)
	case core.ModeIterative:
		m.sortIterative(result)
	case core.ModeParallel:
		m.sortParallel(result, 0, len(result)-1, safeMetrics)
		m.metrics.Comparisons = int(safeMetrics.Comparisons.Load())
		m.metrics.Swaps = int(safeMetrics.Swaps.Load())
	default:
		// Caso ocorra algum erro no modo, usa recursivo como fallback
		m.sortRecursive(result, 0, len(result)-1)
	}

	m.metrics.Time = time.Since(start)
	return result
}

func (m *MergeSort) sortParallel(arr []int, left, right int, metrics *SafeMetrics) {
	if left >= right {
		return
	}

	// Threshold para mudar para algoritmo sequencial
	threshold := 1000

	if right-left < threshold {
		m.sortRecursiveWithMetrics(arr, left, right, metrics)
		return
	}
	//fmt.Println("oiaaaaaaaa")

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

func (m *MergeSort) sortRecursiveWithMetrics(arr []int, left, right int, metrics *SafeMetrics) {
	if left >= right {
		return
	}

	mid := left + (right-left)/2
	m.sortRecursiveWithMetrics(arr, left, mid, metrics)
	m.sortRecursiveWithMetrics(arr, mid+1, right, metrics)
	m.mergeWithMetrics(arr, left, mid, right, metrics)
}

func (m *MergeSort) mergeWithMetrics(arr []int, left, mid, right int, metrics *SafeMetrics) {
	if left >= right {
		return
	}

	temp := make([]int, right-left+1)
	i, j, k := left, mid+1, 0

	for i <= mid && j <= right {
		metrics.Comparisons.Add(1)
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
		k++
		metrics.Swaps.Add(1)
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

	// Copia o resultado de volta para o array original
	for i := 0; i < len(temp); i++ {
		arr[left+i] = temp[i]
	}
}

func (m *MergeSort) sortRecursive(arr []int, left, right int) {
	if left >= right {
		return
	}

	mid := left + (right-left)/2
	m.sortRecursive(arr, left, mid)
	m.sortRecursive(arr, mid+1, right)
	m.merge(arr, left, mid, right)
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
	if left >= right {
		return
	}

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

	for i := 0; i < len(temp); i++ {
		arr[left+i] = temp[i]
	}
}
