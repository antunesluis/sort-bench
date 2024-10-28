package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort-bench/internal/analysis"
	"sort-bench/internal/sorting"
	"sort-bench/internal/utils"
	"strings"
)

const (
	exitCodeSuccess = 0
	exitCodeError   = 1
)

// Flags específicas para o comando sort
type sortFlags struct {
	inputFile  string
	outputFile string
	algorithm  string
	mode       string
	analyze    bool
}

// Flags específicas para o comando compare
type compareFlags struct {
	inputFile  string
	algorithms string
	mode       string
}

func parseSortFlags() (*sortFlags, error) {
	fs := flag.NewFlagSet("sort", flag.ExitOnError)
	f := &sortFlags{}

	fs.StringVar(&f.inputFile, "input", "", "Arquivo de entrada")
	fs.StringVar(&f.outputFile, "output", "", "Arquivo de saída")
	fs.StringVar(&f.algorithm, "algo", "mergesort", "Algoritmo de ordenação (mergesort, quicksort, bubblesort, heapsort)")
	fs.StringVar(&f.mode, "mode", "recursive", "Modo de execução (recursive, iterative, parallel)")
	fs.BoolVar(&f.analyze, "analyze", false, "Mostra análise detalhada do algoritmo")

	err := fs.Parse(os.Args[2:])
	return f, err
}

func parseCompareFlags() (*compareFlags, error) {
	fs := flag.NewFlagSet("compare", flag.ExitOnError)
	f := &compareFlags{}

	fs.StringVar(&f.inputFile, "input", "", "Arquivo de entrada")
	fs.StringVar(&f.algorithms, "algorithms", "", "Lista de algoritmos separados por vírgula")
	fs.StringVar(&f.mode, "mode", "recursive", "Modo de execução (recursive, iterative, parallel)")

	err := fs.Parse(os.Args[2:])
	return f, err
}

func validateSortFlags(f *sortFlags) error {
	if f.inputFile == "" {
		return fmt.Errorf("arquivo de entrada é obrigatório")
	}

	absPath, err := filepath.Abs(f.inputFile)
	if !filepath.IsAbs(f.inputFile) {
		if err != nil {
			return fmt.Errorf("erro ao resolver caminho do arquivo de entrada: %w", err)
		}
		f.inputFile = absPath
	}

	return nil
}

func validateCompareFlags(f *compareFlags) error {
	if f.inputFile == "" {
		return fmt.Errorf("arquivo de entrada é obrigatório")
	}

	if f.algorithms == "" {
		return fmt.Errorf("lista de algoritmos é obrigatória")
	}

	absPath, err := filepath.Abs(f.inputFile)
	if !filepath.IsAbs(f.inputFile) {
		if err != nil {
			return fmt.Errorf("erro ao resolver caminho do arquivo de entrada: %w", err)
		}
		f.inputFile = absPath
	}

	return nil
}

func printAnalysis(result analysis.BenchmarkResult) {
	fmt.Printf("\nResultados para %s (%s):\n", result.Algorithm, result.Mode)
	fmt.Printf("Tempo de execução: %v\n", result.Metrics.Time)
	fmt.Printf("Comparações: %d\n", result.Metrics.Comparisons)
	fmt.Printf("Trocas: %d\n", result.Metrics.Swaps)
	fmt.Printf("Memória utilizada: %d bytes\n", result.Metrics.Memory)
}

func executeSort(flags *sortFlags) error {
	fileHandler := utils.NewFileHandler()
	sorterFactory := sorting.NewSorterFactory()
	benchmarker := analysis.NewBenchmarker()

	numbers, err := fileHandler.ReadNumbers(flags.inputFile)
	if err != nil {
		return fmt.Errorf("erro ao ler arquivo: %w", err)
	}

	sorter, err := sorterFactory.GetSorter(flags.algorithm, flags.mode)
	if err != nil {
		return fmt.Errorf("erro ao criar sorter: %w", err)
	}

	var sortedNumbers []int
	if flags.analyze {
		result := benchmarker.RunBenchmark(sorter, numbers)
		printAnalysis(result)
		sortedNumbers = result.SortedData
	} else {
		sortedNumbers = sorter.Sort(numbers)
	}

	if flags.outputFile != "" {
		if err := fileHandler.WriteNumbers(flags.outputFile, sortedNumbers); err != nil {
			return fmt.Errorf("erro ao escrever resultado: %w", err)
		}
		fmt.Printf("Arquivo ordenado salvo em: %s\n", flags.outputFile)
	}

	return nil
}

func executeCompare(flags *compareFlags) error {
	fileHandler := utils.NewFileHandler()
	sorterFactory := sorting.NewSorterFactory()
	benchmarker := analysis.NewBenchmarker()

	numbers, err := fileHandler.ReadNumbers(flags.inputFile)
	if err != nil {
		return fmt.Errorf("erro ao ler arquivo: %w", err)
	}

	algorithms := strings.Split(flags.algorithms, ",")

	fmt.Printf("\nComparando algoritmos para %d elementos:\n", len(numbers))
	fmt.Println(strings.Repeat("-", 80))
	fmt.Printf("%-12s %-12s %-15s %-12s %-12s %-15s\n",
		"Algoritmo", "Modo", "Tempo", "Comparações", "Trocas", "Memória (bytes)")
	fmt.Println(strings.Repeat("-", 80))

	for _, algo := range algorithms {
		sorter, err := sorterFactory.GetSorter(strings.TrimSpace(algo), flags.mode)
		if err != nil {
			return fmt.Errorf("erro ao criar sorter para %s: %w", algo, err)
		}

		result := benchmarker.RunBenchmark(sorter, append([]int(nil), numbers...))
		fmt.Printf("%-12s %-12s %-15v %-12d %-12d %-15d\n",
			result.Algorithm,
			result.Mode,
			result.Metrics.Time,
			result.Metrics.Comparisons,
			result.Metrics.Swaps,
			result.Metrics.Memory)
	}

	fmt.Println(strings.Repeat("-", 80))
	return nil
}

func printUsage() {
	fmt.Println("Uso:")
	fmt.Println("  sort-bench sort -input <arquivo> [-output <arquivo>] [-algo <algoritmo>] [-mode <modo>] [-analyze]")
	fmt.Println("  sort-bench compare -input <arquivo> -algorithms <algo1,algo2,...> [-mode <modo>]")
	fmt.Println("\nModos disponíveis:")
	fmt.Println("  recursive, iterative, parallel")
	fmt.Println("\nAlgoritmos disponíveis:")
	fmt.Println("  mergesort, quicksort, heapsort, bubblesort")
}

func run() error {
	if len(os.Args) < 2 {
		printUsage()
		return fmt.Errorf("comando é obrigatório")
	}

	switch os.Args[1] {
	case "sort":
		flags, err := parseSortFlags()
		if err != nil {
			return err
		}
		if err := validateSortFlags(flags); err != nil {
			return err
		}
		return executeSort(flags)

	case "compare":
		flags, err := parseCompareFlags()
		if err != nil {
			return err
		}
		if err := validateCompareFlags(flags); err != nil {
			return err
		}
		return executeCompare(flags)

	default:
		printUsage()
		return fmt.Errorf("comando desconhecido: %s", os.Args[1])
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Erro: %v\n", err)
		os.Exit(exitCodeError)
	}
	os.Exit(exitCodeSuccess)
}

