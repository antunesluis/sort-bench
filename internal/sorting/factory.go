package sorting

import (
	"fmt"
	"sort-bench/internal/core"
	"strings"
)

type SorterFactory struct{}

func NewSorterFactory() *SorterFactory {
	return &SorterFactory{}
}

func (f *SorterFactory) GetSorter(algorithm string, mode string) (core.Sorter, error) {
	// Normaliza o modo para lowercase
	mode = strings.ToLower(mode)

	// Valida o modo
	var sortMode core.SortMode
	switch mode {
	case string(core.ModeRecursive), string(core.ModeIterative), string(core.ModeParallel):
		sortMode = core.SortMode(mode)
	default:
		return nil, fmt.Errorf("modo não suportado: %s. Use 'recursive', 'iterative' ou 'parallel'", mode)
	}

	// Normaliza o algoritmo para lowercase
	algorithm = strings.ToLower(algorithm)

	switch algorithm {
	case string(core.AlgoMergeSort):
		return NewMergeSort(sortMode), nil
	case string(core.AlgoQuickSort):
		return NewQuickSort(sortMode), nil
	case string(core.AlgoHeapsort):
		return NewHeapSort(sortMode), nil
	default:
		return nil, fmt.Errorf("algoritmo não suportado: %s", algorithm)
	}
}

