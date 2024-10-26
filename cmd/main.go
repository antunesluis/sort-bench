package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort-bench/internal/sorting"
	"sort-bench/internal/utils"
)

const (
	exitCodeSuccess = 0
	exitCodeError   = 1
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

	numbers, err := fileHandler.ReadNumbers(flags.inputFile)
	if err != nil {
		log.Fatal("Erro ao ler arquivo:", err)
	}

	// Execução da ordenação
	sorter, err := sorterFactory.GetSorter(flags.algorithm, flags.mode)
	if err != nil {
		log.Fatal("Erro ao criar sorter:", err)
	}

	var result []int
	result = sorter.Sort(numbers)

	// Escrita do resultado
	if flags.outputFile != "" {
		if err := fileHandler.WriteNumbers(flags.outputFile, result); err != nil {
			log.Fatal("Erro ao escrever resultado:", err)
		}
		fmt.Printf("Arquivo ordenado salvo em: %s\n", flags.outputFile)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Erro: %v\n", err)
		os.Exit(exitCodeError)
	}
	os.Exit(exitCodeSuccess)
}
