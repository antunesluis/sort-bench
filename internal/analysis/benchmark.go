package analysis

import (
	"fmt"
	"runtime"
	"sort-bench/internal/sorting"
	"strings"
	"time"
)

func (b *Benchmark) RunComparison(algorithms []sorting.Sorter, inputs [][]int) error {
	for _, input := range inputs {
		result := BenchmarkResult{
			InputSize: len(input),
			Metrics:   make(map[string]AlgorithmMetrics),
		}

		for _, algo := range algorithms {
			// Measure memory before
			var memStatsBefore runtime.MemStats
			runtime.ReadMemStats(&memStatsBefore)

			// Run sort
			start := time.Now()
			sorted := algo.Sort(input)
			duration := time.Since(start)

			// Measure memory after
			var memStatsAfter runtime.MemStats
			runtime.ReadMemStats(&memStatsAfter)

			// Get metrics
			metrics := algo.GetMetrics()

			result.Metrics[algo.Name()] = sorting.AlgorithmMetrics{
				Name:        algo.Name(),
				Comparisons: metrics.Comparisons,
				Swaps:       metrics.Swaps,
				Time:        duration,
				Memory:      int64(memStatsAfter.TotalAlloc - memStatsBefore.TotalAlloc),
				Tree:        metrics.Tree,
			}

			// Validate sort
			if !IsSorted(sorted) {
				return fmt.Errorf("algorithm %s failed to sort correctly", algo.Name())
			}
		}

		b.results = append(b.results, result)
	}

	return nil
}

func (b *Benchmark) GenerateReport() string {
	var report strings.Builder

	report.WriteString("\nSorting Algorithm Comparison Report\n")
	report.WriteString("================================\n\n")

	for _, result := range b.results {
		report.WriteString(fmt.Sprintf("Input Size: %d\n", result.InputSize))
		report.WriteString("--------------------\n")

		for algo, metrics := range result.Metrics {
			report.WriteString(fmt.Sprintf("\n%s:\n", algo))
			report.WriteString(fmt.Sprintf("  Time: %v\n", metrics.Time))
			report.WriteString(fmt.Sprintf("  Comparisons: %d\n", metrics.Comparisons))
			report.WriteString(fmt.Sprintf("  Swaps: %d\n", metrics.Swaps))
			report.WriteString(fmt.Sprintf("  Memory: %d bytes\n", metrics.Memory))
		}
		report.WriteString("\n")
	}

	return report.String()
}

func IsSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}
