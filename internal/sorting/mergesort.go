package sorting

import (
	"fmt"
	"time"
)

type MergeSort struct {
	BaseSorter
	tree *analysis.RecursionTree
}

func NewMergeSort(mode SortMode) *MergeSort {
	return &MergeSort{
		BaseSorter: NewBaseSorter("MergeSort", mode),
		tree: &analysis.RecursionTree{
			Nodes: make([]analysis.NodeInfo, 0),
		},
	}
}

func (m *MergeSort) Sort(arr []int) []int {
	start := time.Now()
	result := make([]int, len(arr))
	copy(result, arr)

	switch m.mode {
	case ModeRecursive:
		m.sortRecursive(result, 0, len(result)-1, 0, "root")
	case ModeIterative:
		m.sortIterative(result)
	}

	m.metrics.Time = time.Since(start)
	return result
}

func (m *MergeSort) sortRecursive(arr []int, left, right, depth int, parentID string) {
	if left < right {
		mid := left + (right-left)/2

		// Record split
		nodeID := fmt.Sprintf("node_%d_%d", left, right)
		subArray := make([]int, right-left+1)
		copy(subArray, arr[left:right+1])

		m.tree.Nodes = append(m.tree.Nodes, analysis.NodeInfo{
			Value:     subArray,
			Depth:     depth,
			NodeID:    nodeID,
			ParentID:  parentID,
			Operation: "split",
		})

		m.sortRecursive(arr, left, mid, depth+1, nodeID)
		m.sortRecursive(arr, mid+1, right, depth+1, nodeID)

		// Record merge
		mergeNodeID := fmt.Sprintf("merge_%d_%d", left, right)
		m.merge(arr, left, mid, right)

		mergedArray := make([]int, right-left+1)
		copy(mergedArray, arr[left:right+1])

		m.tree.Nodes = append(m.tree.Nodes, analysis.NodeInfo{
			Value:     mergedArray,
			Depth:     depth,
			NodeID:    mergeNodeID,
			ParentID:  nodeID,
			Operation: "merge",
		})
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

func (m *MergeSort) GetRecursionTree() *analysis.RecursionTree {
	return m.tree
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

