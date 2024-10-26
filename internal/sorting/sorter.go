package sorting

type Sorter interface {
	Sort([]int) []int
	Name() string
	Mode() string
	GetMetrics() SortMetrics
}

type SortMetrics interface {
	GetComparisons() int
	GetSwaps() int
	GetExecutionTime() time.Duration
	GetMemoryUsage() float64
}

type Visualizer interface {
	TrackOperation(data []int, depth int, operationType string)
	GenerateVisualization(filename string) error
	Clear()
}

type Analyzer interface {
	RunBenchmark(sorter Sorter, data []int, inputType string) AnalysisResult
	CompareAlgorithms(algorithms []Sorter, datasets map[string][]int) string
	GenerateReport(format string) (string, error)
}

