package sorting

import (
	"fmt"
	"sort-bench/internal/core"
)

type SorterFactory struct{}

func NewSorterFactory() *SorterFactory {
	return &SorterFactory{}
}

func (f *SorterFactory) GetSorter(algorithm string, mode string) (core.Sorter, error) {
	sortMode := core.SortMode(mode)

	switch algorithm {
	case "mergesort":
		return NewMergeSort(sortMode), nil
	case "quicksort":
		return NewQuickSort(sortMode), nil
	case "heapsort":
		return NewHeapSort(sortMode), nil
	// case "quicksort":
	// 	return NewQuickSort(sortMode), nil
	// case "heapsort":
	// 	return NewHeapSort(sortMode), nil
	// case "bubblesort":
	// 	return NewBubbleSort(sortMode), nil
	default:
		return nil, fmt.Errorf("algoritmo não suportado: %s", algorithm)
	}
}
