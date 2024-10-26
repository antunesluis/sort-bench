package sorting

import (
	"errors"
	"fmt"
)

// SortFactory é responsável por fornecer a implementação correta de Sorter
type SortFactory struct{}

// NewSorterFactory cria uma nova instância de SortFactory
func NewSorterFactory() *SortFactory {
	return &SortFactory{}
}

// GetSorter retorna uma implementação de Sorter com base no algoritmo e modo especificado
func (f *SortFactory) GetSorter(algorithm string, mode string) (Sorter, error) {
	switch algorithm {
	case "mergesort":
		return NewMergeSort(mode), nil
	// Aqui você pode adicionar outros algoritmos como quicksort, bubblesort, etc.
	default:
		return nil, errors.New(fmt.Sprintf("Algoritmo '%s' não é suportado", algorithm))
	}
}
