package main

import (
	"flag"
	"fmt"
	"log"
	"sort-bench/internal/analysis"
	"sort-bench/internal/config"
	"sort-bench/internal/sorting"
	"sort-bench/internal/utils"
)

func main() {
	// Flags globais usando o package flag padrão
	inputFile := flag.String("input", "", "Arquivo de entrada")
	outputFile := flag.String("output", "", "Arquivo de saída")
	algorithm := flag.String("algo", "mergesort", "Algoritmo de ordenação (mergesort, quicksort, bubblesort, heapsort)")
	mode := flag.String("mode", "recursive", "Modo de execução (recursive, iterative, parallel)")
	analyze := flag.Bool("analyze", false, "Realizar análise de performance")
	visualize := flag.Bool("visualize", false, "Gerar visualização da árvore de recursão")

	// Parsing dos argumentos
	flag.Parse()

	if *inputFile == "" {
		log.Fatal("Arquivo de entrada é obrigatório")
		// Validação de argumentos obrigatórios
	}

	cfg := &config.Config{
		InputFile:  *inputFile,
		OutputFile: *outputFile,
		Algorithm:  *algorithm,
		Mode:       *mode,
		Analyze:    *analyze,
	}

	fmt.Println("Configurações:")
	fmt.Printf("Input: %s\n", cfg.InputFile)
	fmt.Printf("Output: %s\n", cfg.OutputFile)
	fmt.Printf("Algorithm: %s\n", cfg.Algorithm)
	fmt.Printf("Mode: %s\n", cfg.Mode)
	fmt.Printf("Analyze: %s\n", cfg.Analyze)

	fileHandler := utils.NewFileHandler()
	sorterFactory := sorting.NewSorterFactory()
	analyzer := analysis.NewPerformanceAnalyzer()
	// Leitura dos dados
	numbers, err := fileHandler.ReadNumbers(cfg.InputFile)
	if err != nil {
		log.Fatal("Erro ao ler arquivo:", err)
	}

	// Execução da ordenação
	sorter, err := sorterFactory.GetSorter(cfg.Algorithm, cfg.Mode)
	if err != nil {
		log.Fatal("Erro ao criar sorter:", err)
	}

	// Ordenação e análise
	var result []int
	var metrics analysis.PerformanceMetrics

	if *analyze {
		metrics = analyzer.Analyze(numbers, sorter)
		result = metrics.SortedData
	} else {
		result = sorter.Sort(numbers)
	}

	// Escrita do resultado
	if cfg.OutputFile != "" {
		if err := fileHandler.WriteNumbers(cfg.OutputFile, result); err != nil {
			log.Fatal("Erro ao escrever resultado:", err)
		}
		fmt.Printf("Arquivo ordenado salvo em: %s\n", cfg.OutputFile)
	}

	// Exibição de métricas se solicitado
	if *analyze {
		fmt.Printf("Performance:\n")
		fmt.Printf("Tempo de execução: %v\n", metrics.ExecutionTime)
		fmt.Printf("Comparações: %d\n", metrics.Comparisons)
		fmt.Printf("Trocas: %d\n", metrics.Swaps)
	}

	// Visualização se solicitada
	if *visualize {
		if err := visualize.GenerateTree(numbers, "tree.png"); err != nil {
			log.Printf("Erro ao gerar visualização: %v", err)
		}
	}
}
