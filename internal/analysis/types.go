package analysis

import "time"

type RecursionTree struct {
	Nodes []NodeInfo
}

type NodeInfo struct {
	Value     []int
	Depth     int
	NodeID    string
	ParentID  string
	Operation string
}

type Benchmark struct {
	results []BenchmarkResult
}

type BenchmarkResult struct {
	InputSize int
	Metrics   map[string]AlgorithmMetrics
}

type AlgorithmMetrics struct {
	Name        string
	Comparisons int
	Swaps       int
	Time        time.Duration
	Memory      int64
	Tree        *RecursionTree
}
