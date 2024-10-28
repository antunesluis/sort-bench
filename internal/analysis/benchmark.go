package analysis

import (
	"runtime"
	"sort-bench/internal/core"
	"time"
)

type BenchmarkResult struct {
	Algorithm  string
	Mode       core.SortMode
	InputSize  int
	Metrics    core.Metrics
	SortedData []int
}

type Benchmarker struct {
	results []BenchmarkResult
}

func NewBenchmarker() *Benchmarker {
	return &Benchmarker{
		results: make([]BenchmarkResult, 0),
	}
}

func (b *Benchmarker) RunBenchmark(sorter core.Sorter, input []int) BenchmarkResult {
	runtime.GC() // Force garbage collection before benchmark

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	startAlloc := m.TotalAlloc

	start := time.Now()
	sortedData := sorter.Sort(input)
	duration := time.Since(start)

	runtime.ReadMemStats(&m)
	memoryUsed := m.TotalAlloc - startAlloc

	metrics := sorter.GetMetrics()
	metrics.Time = duration
	metrics.Memory = int64(memoryUsed)

	benchResult := BenchmarkResult{
		Algorithm:  sorter.Name(),
		Mode:       sorter.Mode(),
		InputSize:  len(input),
		Metrics:    metrics,
		SortedData: sortedData,
	}

	b.results = append(b.results, benchResult)
	return benchResult
}

