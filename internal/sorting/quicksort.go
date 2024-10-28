package sorting

import (
	"sort-bench/internal/core"
	"time"
)

type QuickSort struct {
	BaseSorter
}

func NewQuickSort(mode core.SortMode) *QuickSort {
	return &QuickSort{
		BaseSorter: NewBaseSorter("QuickSort", mode),
	}
}

func (q *QuickSort) Sort(arr []int) []int {
	start := time.Now()
	q.metrics = core.Metrics{} // Inicializa as métricas

	result := arr
	switch q.Mode() {
	case core.ModeRecursive:
		q.quickSort(result, 0, len(arr)-1)
	case core.ModeIterative:
		q.quickSortIterative(result, 0, len(arr)-1)
	}

	q.metrics.Time = time.Since(start)
	return arr
}

func (q *QuickSort) quickSortIterative(arr []int, low, high int) {
	// Simula a recursão com uma pilha
	stack := make([][2]int, 0, high-low+1) // Cada entrada é [low, high]

	// Empilha o intervalo inicial
	stack = append(stack, [2]int{low, high})

	for len(stack) > 0 {
		// Remove o topo da pilha
		n := len(stack) - 1
		low, high = stack[n][0], stack[n][1]
		stack = stack[:n]

		if low < high {
			p := q.partition(arr, low, high)

			// Empilha os sub-arranjos direito e esquerdo
			stack = append(stack, [2]int{low, p - 1})
			stack = append(stack, [2]int{p + 1, high})
		}
	}
}

func (q *QuickSort) quickSort(arr []int, low, high int) {
	if low < high {
		p := q.partition(arr, low, high)
		q.quickSort(arr, low, p-1)
		q.quickSort(arr, p+1, high)
	}
}

func (q *QuickSort) partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		q.metrics.Comparisons++ // Incrementa comparações
		if arr[j] < pivot {
			i++
			q.swap(arr, i, j)
		}
	}
	q.swap(arr, i+1, high)
	return i + 1
}

func (q *QuickSort) swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
	q.metrics.Swaps++ // Incrementa trocas
}
