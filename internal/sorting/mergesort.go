package sorting

import (
	"time"
)

type MergeSort struct {
	BaseSorter
}

func NewMergeSort(mode SortMode) *MergeSort {
	return &MergeSort{
		BaseSorter: NewBaseSorter("MergeSort", mode),
	}
}

func (m *MergeSort) Sort(arr []int) []int {
	start := time.Now()
	result := make([]int, len(arr))
	copy(result, arr)

	switch m.mode {
	case ModeRecursive:
		m.sortRecursive(result, 0, len(result)-1, 0)
	case ModeIterative:
		m.sortIterative(result)
		// case ModeParallel:
		// result = m.sortParallel(result)
	}

	m.metrics.Time = time.Since(start)
	return result
}

func (m *MergeSort) sortRecursive(arr []int, left, right, depth int) {
	if left < right {
		mid := left + (right-left)/2

		m.sortRecursive(arr, left, mid, depth+1)
		m.sortRecursive(arr, mid+1, right, depth+1)
		m.merge(arr, left, mid, right, depth)
	}
}

func (m *MergeSort) sortIterative(arr []int) {
	n := len(arr)
	for size := 1; size < n; size *= 2 {
		for left := 0; left < n-1; left += 2 * size {
			mid := min(left+size-1, n-1)
			right := min(left+2*size-1, n-1)
			m.merge(arr, left, mid, right, 0)
		}
	}
}

func (m *MergeSort) merge(arr []int, left, mid, right, depth int) {
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
		m.metrics.Swaps++
		k++
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
