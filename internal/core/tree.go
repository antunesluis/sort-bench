package core

type TreeNode struct {
	Value     []int
	Depth     int
	NodeID    string
	ParentID  string
	Operation string
}

type RecursionTree struct {
	Nodes []TreeNode
}
