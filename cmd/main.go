package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	// "path/filepath"
	// "sort-bench/internal/analysis"
	"sort-bench/internal/sorting"
	// "sort-bench/internal/utils"
	// "sort-bench/internal/visualization"
	// "sort-bench/pkg/benchmark"
)

func main() {
	// Definir subcomandos
	sortCmd := flag.NewFlagSet("sort", flag.ExitOnError)
	// compareCmd := flag.NewFlagSet("compare", flag.ExitOnError)
	// visualizeCmd := flag.NewFlagSet("visualize", flag.ExitOnError)
	// benchmarkCmd := flag.NewFlagSet("benchmark", flag.ExitOnError)

	// Flags para o subcomando 'sort'
	sortAlgorithm := sortCmd.String("algorithm", "mergesort", "Algoritmo de ordenação (mergesort, quicksort, bubblesort, heapsort)")
	sortInput := sortCmd.String("input", "", "Caminho do arquivo de entrada")
	sortOutput := sortCmd.String("output", "", "Caminho do arquivo de saída")

	// // Flags para o subcomando 'compare'
	// compareAlgorithms := compareCmd.String("algorithms", "mergesort,quicksort", "Algoritmos para comparar, separados por vírgula")
	// compareInput := compareCmd.String("input", "", "Caminho do arquivo de entrada para comparação")
	//
	// // Flags para o subcomando 'visualize'
	// visualizeInput := visualizeCmd.String("input", "", "Caminho do arquivo de entrada para visualização")
	// visualizeOutput := visualizeCmd.String("output", "", "Caminho do arquivo de saída para a visualização")
	//
	// // Flags para o subcomando 'benchmark'
	// benchmarkAlgorithm := benchmarkCmd.String("algorithm", "mergesort", "Algoritmo para benchmark")
	// benchmarkInput := benchmarkCmd.String("input", "", "Caminho do arquivo de entrada para benchmark")

	if len(os.Args) < 2 {
		fmt.Println("Uso esperado: [sort|compare|visualize|benchmark] [opções]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "sort":
		sortCmd.Parse(os.Args[2:])
		runSort(*sortAlgorithm, *sortInput, *sortOutput)
	// case "compare":
	// 	compareCmd.Parse(os.Args[2:])
	// 	runCompare(*compareAlgorithms, *compareInput)
	// case "visualize":
	// 	visualizeCmd.Parse(os.Args[2:])
	// 	runVisualize(*visualizeInput, *visualizeOutput)
	// case "benchmark":
	// 	benchmarkCmd.Parse(os.Args[2:])
	// 	runBenchmark(*benchmarkAlgorithm, *benchmarkInput)
	default:
		fmt.Println("Subcomando desconhecido")
		os.Exit(1)
	}
}

func runSort(algorithm, inputPath, outputPath string) {
	// Ler o arquivo de entrada
	numbers, err := utils.ReadNumbersFromFile(inputPath)
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo de entrada: %v", err)
	}

	// Ordenar os números
	var sortedNumbers []int
	switch algorithm {
	case "mergesort":
		sortedNumbers = sorting.MergeSort(numbers)
	// case "quicksort":
	// 	sortedNumbers = sorting.QuickSort(numbers)
	default:
		log.Fatalf("Algoritmo de ordenação desconhecido: %s", algorithm)
	}

	// Escrever o resultado no arquivo de saída
	err = utils.WriteNumbersToFile(outputPath, sortedNumbers)
	if err != nil {
		log.Fatalf("Erro ao escrever no arquivo de saída: %v", err)
	}

	fmt.Printf("Ordenação concluída. Resultado salvo em %s\n", outputPath)
}
