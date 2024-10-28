package sorting

import (
	"sort-bench/internal/core"
	"time"
)

type HeapSort struct {
	BaseSorter
}

func NewHeapSort(mode core.SortMode) *HeapSort {
	return &HeapSort{
		BaseSorter: NewBaseSorter("HeapSort", mode),
	}
}

func (h *HeapSort) Sort(arr []int) []int {
	start := time.Now()
	h.metrics = core.Metrics{} // Inicializa as métricas

	switch h.Mode() {
	case core.ModeRecursive:
		h.sortRecursive(arr)
	case core.ModeIterative:
		h.sortIterative(arr)
	}

	h.metrics.Time = time.Since(start)
	return arr
}

// Modo recursivo: usa a abordagem tradicional de heapify com recursão
func (h *HeapSort) sortRecursive(arr []int) {
	n := len(arr)

	// Constrói o heap
	for i := n/2 - 1; i >= 0; i-- {
		h.heapifyRecursive(arr, n, i)
	}

	// Extrai elementos do heap
	for i := n - 1; i > 0; i-- {
		h.swap(arr, 0, i) // Move a raiz para o fim
		h.heapifyRecursive(arr, i, 0)
	}
}

func (h *HeapSort) heapifyRecursive(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	// Verifica se o filho esquerdo é maior que a raiz
	if left < n && arr[left] > arr[largest] {
		largest = left
		h.metrics.Comparisons++
	}

	// Verifica se o filho direito é maior que o maior até agora
	if right < n && arr[right] > arr[largest] {
		largest = right
		h.metrics.Comparisons++
	}

	// Se o maior não é a raiz, faz a troca e continua heapificando
	if largest != i {
		h.swap(arr, i, largest)
		h.heapifyRecursive(arr, n, largest)
	}
}

// Modo iterativo: elimina a recursão, controlando a heapificação com um loop
func (h *HeapSort) sortIterative(arr []int) {
	n := len(arr)

	// Constrói o heap
	for i := n/2 - 1; i >= 0; i-- {
		h.heapifyIterative(arr, n, i)
	}

	// Extrai elementos do heap
	for i := n - 1; i > 0; i-- {
		h.swap(arr, 0, i) // Move a raiz para o fim
		h.heapifyIterative(arr, i, 0)
	}
}

// Versão iterativa do heapify
func (h *HeapSort) heapifyIterative(arr []int, n, i int) {
	for {
		largest := i
		left := 2*i + 1
		right := 2*i + 2

		// Verifica se o filho esquerdo é maior que a raiz
		if left < n && arr[left] > arr[largest] {
			largest = left
			h.metrics.Comparisons++
		}

		// Verifica se o filho direito é maior que o maior até agora
		if right < n && arr[right] > arr[largest] {
			largest = right
			h.metrics.Comparisons++
		}

		// Se o maior ainda é a raiz, termina a iteração
		if largest == i {
			break
		}

		// Caso contrário, troca e continua ajustando o sub-heap
		h.swap(arr, i, largest)
		i = largest
	}
}

// Realiza a troca de dois elementos no array
func (h *HeapSort) swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
	h.metrics.Swaps++ // Incrementa o número de trocas
}

