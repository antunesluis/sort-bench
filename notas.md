# Estrutura

```shell
sort-bench/
├── cmd/
│   └── sort-bench/
│       └── main.go
├── internal/
│   ├── domain/
│   │   ├── algorithm/
│   │   │   ├── sorter.go        # interfaces e tipos base
│   │   │   ├── metrics.go       # tipos de métricas
│   │   │   └── types.go         # tipos compartilhados do domínio
│   │   └── benchmark/
│   │       └── types.go         # tipos específicos de benchmark
│   ├── algorithm/
│   │   ├── factory.go
│   │   └── implementations/
│   │       ├── mergesort/
│   │       │   ├── mergesort.go
│   │       │   └── tree.go
│   │       └── quicksort/
│   ├── benchmark/
│   │   ├── service.go           # lógica de benchmark
│   │   └── report/
│   │       ├── generator.go     # geração de relatórios
│   │       └── visualization.go  # visualizações
│   └── platform/
│       └── file/
```

# Arquitetura

Principais características desta arquitetura:

1. Camada de Domínio:

- Define interfaces e tipos principais
- Independente de implementações
- Contém regras de negócio centrais

2. Camada de Algoritmos:

- Implementações concretas dos algoritmos
- Cada algoritmo é um módulo independente
- Métricas coletadas automaticamente

3. Camada de Benchmark:

- Serviço central que coordena execuções
- Geração de relatórios
- Visualizações dos resultados

4. Camada de Plataforma:

- Manipulação de arquivos
- Outros serviços de infraestrutura

### Para implementar um novo algoritmo de ordenação

- Criar novo pacote em `algorithm/implementations/`
- Implementar interface `algorithm.Sorter`
- Adicionar ao slice de algoritmos no `main.go`

### Para adicionar novas funcionalidades

1. Novo tipo de relatório:

- Implementar interface `benchmark.Reporter`
- Injetar no serviço de benchmark

2. Nova visualização:

- Implementar interface `benchmark.Visualizer`
- Adicionar ao serviço de benchmark

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
