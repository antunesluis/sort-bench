# Estrutura

```shell
projeto-ordenacao/
├── cmd/
│   └── main.go
├── internal/
│   ├── sorting/
│   │   ├── mergesort.go
│   │   ├── quicksort.go
│   │   ├── bubblesort.go
│   │   ├── heapsort.go
│   │   └── parallel_mergesort.go
│   ├── analysis/
│   │   ├── complexity.go
│   │   └── comparison.go
│   ├── visualization/
│   │   └── recursion_tree.go
│   └── utils/
│       ├── file_handling.go
│       └── data_generation.go
├── pkg/
│   └── benchmark/
│       └── benchmark.go
├── tests/
│   └── sorting_test.go
├── data/
│   ├── input/
│   └── output/
├── go.mod
└── README.md
```

# Usos

- Para ordenar um arquivo de entrada:

```shell
go run main.go sort -algorithm mergesort -input ./data/input/numbers.txt -output ./data/output/sorted.txt
```

- Para comparar algoritmos:

```shell
go run main.go compare -algorithms mergesort,quicksort -input ./data/input/numbers.txt
```

- Para gerar a visualização gráfica da arvore de recursão:

```shell
go run main.go visualize -input ./data/input/numbers.txt -output ./data/output/visualization.png
```
