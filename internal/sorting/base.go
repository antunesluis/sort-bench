package sorting

import (
	"sort-bench/internal/core"
)

type BaseSorter struct {
	name    string
	mode    core.SortMode
	metrics core.Metrics
	tree    *core.RecursionTree
}

func NewBaseSorter(name string, mode core.SortMode) BaseSorter {
	return BaseSorter{
		name: name,
		mode: mode,
		tree: &core.RecursionTree{
			Nodes: make([]core.TreeNode, 0),
		},
	}
}

func (b *BaseSorter) Name() string                          { return b.name }
func (b *BaseSorter) Mode() core.SortMode                   { return b.mode }
func (b *BaseSorter) GetMetrics() core.Metrics              { return b.metrics }
func (b *BaseSorter) GetRecursionTree() *core.RecursionTree { return b.tree }

func (b *BaseSorter) compare(x, y int) bool {
	b.metrics.Comparisons++
	return x < y
}

func (b *BaseSorter) swap(arr []int, i, j int) {
	b.metrics.Swaps++
	arr[i], arr[j] = arr[j], arr[i]
}

func (b *BaseSorter) recordNode(value []int, depth int, nodeID, parentID, operation string) {
	b.tree.Nodes = append(b.tree.Nodes, core.TreeNode{
		Value:     value,
		Depth:     depth,
		NodeID:    nodeID,
		ParentID:  parentID,
		Operation: operation,
	})
}

