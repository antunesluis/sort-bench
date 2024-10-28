// package main
//
// import (
// 	"flag"
// 	"fmt"
// 	"os"
// 	"sort-bench/internal/analysis"
// 	"sort-bench/internal/sorting"
// 	"sort-bench/internal/utils"
// 	"strings"
// )
//
// type Command struct {
// 	Name        string
// 	Description string
// 	Execute     func([]string) error
// }
//
// func main2() {
// 	if len(os.Args) < 2 {
// 		printUsage()
// 		os.Exit(1)
// 	}
//
// 	commands := map[string]Command{
// 		"sort": {
// 			Name:        "sort",
// 			Description: "Ordena um arquivo usando um algoritmo específico",
// 			Execute:     executeSort,
// 		},
// 		"compare": {
// 			Name:        "compare",
// 			Description: "Compara diferentes algoritmos de ordenação",
// 			Execute:     executeCompare,
// 		},
// 	}
//
// 	command := os.Args[1]
// 	cmd, exists := commands[command]
// 	if !exists {
// 		fmt.Printf("Comando desconhecido: %s\n", command)
// 		printUsage()
// 		os.Exit(1)
// 	}
//
// 	if err := cmd.Execute(os.Args[2:]); err != nil {
// 		fmt.Fprintf(os.Stderr, "Erro: %v\n", err)
// 		os.Exit(1)
// 	}
// }
//
// func printUsage() {
// 	fmt.Println("Uso:")
// 	fmt.Println("  sort-bench sort -algorithm <algo> -mode <mode> -input <file> -output <file>")
// 	fmt.Println("  sort-bench compare -algorithms <algo1,algo2,...> -input <file>")
// 	fmt.Println("\nModos disponíveis:")
// 	fmt.Println("  recursive, iterative, parallel")
// 	fmt.Println("\nAlgoritmos disponíveis:")
// 	fmt.Println("  mergesort, quicksort, heapsort, bubblesort")
// }
//
// func executeSort(args []string) error {
// 	fs := flag.NewFlagSet("sort", flag.ExitOnError)
// 	input := fs.String("input", "", "Arquivo de entrada")
// 	output := fs.String("output", "", "Arquivo de saída")
// 	algorithm := fs.String("algorithm", "mergesort", "Algoritmo de ordenação")
// 	mode := fs.String("mode", "recursive", "Modo de execução")
//
// 	if err := fs.Parse(args); err != nil {
// 		return err
// 	}
//
// 	if *input == "" {
// 		return fmt.Errorf("arquivo de entrada é obrigatório")
// 	}
//
// 	fileHandler := utils.NewFileHandler()
// 	numbers, err := fileHandler.ReadNumbers(*input)
// 	if err != nil {
// 		return fmt.Errorf("erro ao ler arquivo: %w", err)
// 	}
//
// 	sorterFactory := sorting.NewSorterFactory()
// 	sorter, err := sorterFactory.GetSorter(*algorithm, *mode)
// 	if err != nil {
// 		return fmt.Errorf("erro ao criar sorter: %w", err)
// 	}
//
// 	benchmarker := analysis.NewBenchmarker()
// 	result := benchmarker.RunBenchmark(sorter, numbers)
//
// 	fmt.Printf("\nResultados para %s (%s):\n", result.Algorithm, result.Mode)
// 	fmt.Printf("Tempo de execução: %v\n", result.Metrics.Time)
// 	fmt.Printf("Comparações: %d\n", result.Metrics.Comparisons)
// 	fmt.Printf("Trocas: %d\n", result.Metrics.Swaps)
// 	fmt.Printf("Memória utilizada: %d bytes\n", result.Metrics.Memory)
//
// 	if *output != "" {
// 		if err := fileHandler.WriteNumbers(*output, result.SortedData); err != nil {
// 			return fmt.Errorf("erro ao escrever resultado: %w", err)
// 		}
// 		fmt.Printf("\nArquivo ordenado salvo em: %s\n", *output)
// 	}
//
// 	return nil
// }
//
// func executeCompare(args []string) error {
// 	fs := flag.NewFlagSet("compare", flag.ExitOnError)
// 	input := fs.String("input", "", "Arquivo de entrada")
// 	algorithmsStr := fs.String("algorithms", "", "Lista de algoritmos separados por vírgula")
// 	mode := fs.String("mode", "recursive", "Modo de execução")
//
// 	if err := fs.Parse(args); err != nil {
// 		return err
// 	}
//
// 	if *input == "" || *algorithmsStr == "" {
// 		return fmt.Errorf("arquivo de entrada e lista de algoritmos são obrigatórios")
// 	}
//
// 	algorithms := strings.Split(*algorithmsStr, ",")
// 	fileHandler := utils.NewFileHandler()
// 	numbers, err := fileHandler.ReadNumbers(*input)
// 	if err != nil {
// 		return fmt.Errorf("erro ao ler arquivo: %w", err)
// 	}
//
// 	benchmarker := analysis.NewBenchmarker()
// 	sorterFactory := sorting.NewSorterFactory()
//
// 	fmt.Printf("\nComparando algoritmos para %d elementos:\n", len(numbers))
// 	fmt.Println(strings.Repeat("-", 80))
// 	fmt.Printf("%-12s %-12s %-15s %-12s %-12s %-15s\n",
// 		"Algoritmo", "Modo", "Tempo", "Comparações", "Trocas", "Memória (bytes)")
// 	fmt.Println(strings.Repeat("-", 80))
//
// 	for _, algo := range algorithms {
// 		sorter, err := sorterFactory.GetSorter(strings.TrimSpace(algo), *mode)
// 		if err != nil {
// 			return fmt.Errorf("erro ao criar sorter para %s: %w", algo, err)
// 		}
//
// 		result := benchmarker.RunBenchmark(sorter, append([]int(nil), numbers...))
// 		fmt.Printf("%-12s %-12s %-15v %-12d %-12d %-15d\n",
// 			result.Algorithm,
// 			result.Mode,
// 			result.Metrics.Time,
// 			result.Metrics.Comparisons,
// 			result.Metrics.Swaps,
// 			result.Metrics.Memory)
// 	}
//
// 	fmt.Println(strings.Repeat("-", 80))
// 	return nil
// }
