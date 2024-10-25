package main

import (
	"sort-bench/internal/sorting"
	"sort-bench/internal/visualization"
)

// TrackedMergeSort estende o MergeSort para registrar as operações de visualização
type TrackedMergeSort struct {
	*sorting.MergeSort
	visualizer *visualization.TreeVisualizer
}

func NewTrackedMergeSort(mode string) *TrackedMergeSort {
	return &TrackedMergeSort{
		MergeSort:  sorting.NewMergeSort(mode),
		visualizer: visualization.NewTreeVisualizer(),
	}
}

// Sobrescreve os métodos do MergeSort para adicionar tracking
func (t *TrackedMergeSort) sortRecursive(arr []int, left, right int) []int {
	if left >= right {
		return arr
	}

	// Registra o estado antes da divisão
	t.visualizer.TrackRecursion(arr[left:right+1], left, right, false)

	mid := left + (right-left)/2
	t.sortRecursive(arr, left, mid)
	t.sortRecursive(arr, mid+1, right)
	t.mergeInPlace(arr, left, mid, right)

	// Registra o estado após o merge
	t.visualizer.TrackRecursion(arr[left:right+1], left, right, true)

	return arr
}

func (t *TrackedMergeSort) Sort(arr []int) []int {
	result := t.MergeSort.Sort(arr)

	// Gera a visualização após a ordenação
	err := t.visualizer.GenerateVisualization("merge_sort_tree.png")
	if err != nil {
		panic(err)
	}

	return result
}

