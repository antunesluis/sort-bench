package analysis

import (
	"sort-bench/internal/sorting"
	"time"
)

// PerformanceMetrics armazena métricas de performance do algoritmo
type PerformanceMetrics struct {
	Algorithm     string        // Nome do algoritmo
	Mode          string        // Modo de operação (recursivo, iterativo, paralelo)
	InputSize     int           // Tamanho da entrada
	ExecutionTime time.Duration // Tempo de execução
	Comparisons   int           // Número de comparações
	Swaps         int           // Número de trocas
	MemoryUsage   int64         // Uso de memória em bytes
	SortedData    []int         // Dados ordenados (opcional)
}

// PerformanceAnalyzer analisa o desempenho dos algoritmos de ordenação
type PerformanceAnalyzer struct {
	metrics []PerformanceMetrics
}

func NewPerformanceAnalyzer() *PerformanceAnalyzer {
	return &PerformanceAnalyzer{}
}

// Analyze executa a análise de performance de um algoritmo
func (p *PerformanceAnalyzer) Analyze(data []int, sorter sorting.Sorter) PerformanceMetrics {
	metrics := PerformanceMetrics{
		Algorithm: sorter.Name(),
		InputSize: len(data),
	}

	// Copia os dados para não modificar o original
	input := make([]int, len(data))
	copy(input, data)

	// Medição do tempo de execução
	start := time.Now()
	metrics.SortedData = sorter.Sort(input)
	metrics.ExecutionTime = time.Since(start)

	return metrics
}

// CompareAlgorithms compara diferentes algoritmos com o mesmo conjunto de dados
func (p *PerformanceAnalyzer) CompareAlgorithms(data []int, sorters []sorting.Sorter) []PerformanceMetrics {
	results := make([]PerformanceMetrics, 0)

	for _, sorter := range sorters {
		metrics := p.Analyze(data, sorter)
		results = append(results, metrics)
	}

	return results
}

// GenerateReport cria um relatório detalhado das análises
func (p *PerformanceAnalyzer) GenerateReport(metrics []PerformanceMetrics) string {
	report := "Performance Analysis Report\n"
	report += "========================\n\n"

	for _, m := range metrics {
		report += "Algorithm: " + m.Algorithm + "\n"
		report += "Input Size: " + string(m.InputSize) + "\n"
		report += "Execution Time: " + m.ExecutionTime.String() + "\n"
		if m.Comparisons > 0 {
			report += "Comparisons: " + string(m.Comparisons) + "\n"
		}
		if m.Swaps > 0 {
			report += "Swaps: " + string(m.Swaps) + "\n"
		}
		report += "------------------------\n"
	}

	return report
}

