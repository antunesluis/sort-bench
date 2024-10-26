package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"sort-bench/internal/analysis"
	"sort-bench/internal/config"
	"sort-bench/internal/sorting"
	"sort-bench/internal/utils"
)

type flags struct {
	inputFile  string
	outputFile string
	algorithm  string
	mode       string
	analyze    bool
}

func parseFlags() *flags {
	f := &flags{}

	flag.StringVar(&f.inputFile, "input", "", "Arquivo de entrada")
	flag.StringVar(&f.outputFile, "output", "", "Arquivo de saída")
	flag.StringVar(&f.algorithm, "algo", "mergesort", "Algoritmo de ordenação (mergesort, quicksort, bubblesort, heapsort)")
	flag.StringVar(&f.mode, "mode", "recursive", "Modo de execução (recursive, iterative, parallel)")
	flag.BoolVar(&f.analyze, "analyze", false, "Realizar análise de performance")

	flag.Parse()
	return f
}

func validateFlags(f *flags) error {
	if f.inputFile == "" {
		return fmt.Errorf("arquivo de entrada é obrigatório")
	}

	if !filepath.IsAbs(f.inputFile) {
		absPath, err := filepath.Abs(f.inputFile)
		if err != nil {
			return fmt.Errorf("erro ao resolver caminho do arquivo de entrada: %w", err)
		}
		f.inputFile = absPath
	}

	return nil
}

func run() error {
	flags := parseFlags()
	if err := validateFlags(flags); err != nil {
		return fmt.Errorf("erro de validação dos argumentos: %w", err)
	}

	// Criar componentes
	fileHandler := utils.NewFileHandler()
	sorterFactory := sorting.NewSorterFactory()
	analyzer := analysis.NewPerformanceAnalyzer()

	// Salvar resultados
	// if flags.outputFile != "" {
	// 	if err := fileHandler.WriteNumbers()(
	// 		flags.outputFile,
	// 		result.SortedData,
	// 		flags.outputFormat,
	// 	); err != nil {
	// 		return fmt.Errorf("erro ao salvar resultado: %w", err)
	// 	}
	// }

	return nil
}

func main() {
	// Flags globais usando o package flag padrão
	inputFile := flag.String("input", "", "Arquivo de entrada")
	outputFile := flag.String("output", "", "Arquivo de saída")
	algorithm := flag.String("algo", "mergesort", "Algoritmo de ordenação (mergesort, quicksort, bubblesort, heapsort)")
	mode := flag.String("mode", "recursive", "Modo de execução (recursive, iterative, parallel)")
	analyze := flag.Bool("analyze", false, "Realizar análise de performance")

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
}
