package sorting

import "fmt"

// SorterFactory cria instâncias de algoritmos de ordenação
type SorterFactory struct{}

func NewSorterFactory() *SorterFactory {
	return &SorterFactory{}
}

func (f *SorterFactory) GetSorter(algorithm string, mode string) (Sorter, error) {
	sortMode := SortMode(mode)

	switch algorithm {
	case "mergesort":
		return NewMergeSort(sortMode), nil
	// case "quicksort":
	// return NewQuickSort(sortMode), nil
	// case "bubblesort":
	// return NewBubbleSort(sortMode, f.tracker), nil
	// case "heapsort":
	// return NewHeapSort(sortMode, f.tracker), nil
	default:
		return nil, fmt.Errorf("algoritmo não suportado: %s", algorithm)
	}
}
