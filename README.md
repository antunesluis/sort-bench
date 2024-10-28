
# Comparação e Análise de Algoritmos de Ordenação

## Descrição

Este projeto implementa e compara diferentes algoritmos de ordenação, com foco em mergesort e sua aplicação em diferentes modos: recursivo, iterativo e paralelo. Além disso, inclui uma visualização em Python da árvore de recursão do mergesort.

Os algoritmos comparados incluem: Mergesort, Quicksort e Heapsort.

O projeto permite ordenar dados a partir de arquivos de entrada, analisar o desempenho e visualizar a divisão dos dados no mergesort.

## Funcionalidades

1. Mergesort Recursivo, Iterativo e Paralelo:

2. Quicksort e HeapSort iterativos e recursivos:
    - Implementação dos diferentes modos de execução do mergesort.
    - Implementação dos algoritmos de ordenação Quicksort e HeapSort.
    - Modo de execução iterativo e recursivo.

3. Comparação de Algoritmos:
    - Compara a eficiência de vários algoritmos de ordenação em termos de tempo, trocas e uso de memória.

4. Visualização da Árvore de Recursão:
    - Visualização gráfica da árvore de recursão do mergesort usando Python e Matplotlib.

## Estrutura do Projeto

```shell
sort-bench/
├── cmd/
    └── main.go            # Ponto de entrada principal do projeto
├── internal/
│   ├── analysis/          # Módulo de análise e benchmarking
│   ├── sorting/           # Implementações dos algoritmos de ordenação
│   ├── utils/             # Leitura/escrita de arquivos
│   └── core/              # Definição de interfaces e estrutura básica
├── python/
│   └── recursion_tree.py  # Script Python para visualização da recursão
├── README.md
└── go.mod                 # Arquivo de dependências Go
```

## Requisistos

- Go 1.20+
- Python 3.x (para visualização)
- Biblioteca Matplotlib e networkx (Instalação: pip install matplotlib)

## Como Executar

1. Ordenar Dados com um Algoritmo:

```bash
go run cmd/main.go sort -input <arquivo_entrada> -output <arquivo_saida> -algo mergesort -mode recursive -analyze
```

2. Comparar Algoritmos

```bash
go run cmd/main.go compare -input <arquivo_entrada> -algorithms "mergesort,quicksort"
```

3. Visualizar a Árvore de Recursão (Python)

```bash
python python/recursion_tree.py <arquivo_entrada>
```

## Arquivos de entrada

Os arquivos de entrada devem conter um número por linha. Exemplo:

```txt
10
5
3
8
1
```

## Análise e Benchmark

Durante a execução com a flag -analyze, o projeto exibe informações como:

- Tempo de execução
- Comparações realizadas
- Trocas efetuadas
- Memória utilizada
