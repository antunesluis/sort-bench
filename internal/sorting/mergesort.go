package sorting

// MergeSort implementa diferentes variantes do algoritmo merge sort
type MergeSort struct {
	// buffer reutilizável para reduzir alocações de memória
	buffer []int
	// modo de operação (recursivo, iterativo)
	mode string
	// metricas de execução
	metrics SortMetrics
}

// NewMergeSort cria uma nova instância de MergeSort
// mode pode ser "recursive" ou "iterative"
func NewMergeSort(mode string) *MergeSort {
	return &MergeSort{
		mode: mode,
	}
}

// Name implementa a interface Sorter
func (m *MergeSort) Name() string {
	return "MergeSort-" + m.mode
}

// Sort implementa a interface Sorter
func (m *MergeSort) Sort(arr []int) []int {
	// Inicializa ou redimensiona o buffer se necessário
	if len(m.buffer) < len(arr) {
		m.buffer = make([]int, len(arr))
	}

	// Cria uma cópia do array para não modificar o original
	result := make([]int, len(arr))
	copy(result, arr)

	// Escolhe o método de ordenação baseado no modo
	switch m.mode {
	case "recursive":
		return m.sortRecursive(result, 0, len(result)-1)
	default: // "iterative" é o padrão
		return m.sortIterative(result)
	}
}

// sortRecursive implementa a versão recursiva
func (m *MergeSort) sortRecursive(arr []int, left, right int) []int {
	if left >= right {
		return arr
	}

	mid := left + (right-left)/2
	m.sortRecursive(arr, left, mid)
	m.sortRecursive(arr, mid+1, right)
	m.mergeInPlace(arr, left, mid, right)

	return arr
}

// sortIterative implementa a versão iterativa
func (m *MergeSort) sortIterative(arr []int) []int {
	n := len(arr)

	// Ordena em níveis crescentes
	for size := 1; size < n; size *= 2 {
		// Usa o buffer temporário para evitar alocações
		for left := 0; left < n-1; left += 2 * size {
			mid := min(left+size, n)
			right := min(left+2*size, n)
			m.mergeInPlace(arr, left, mid-1, right-1)
		}
	}

	return arr
}

// mergeInPlace realiza o merge in-place usando o buffer
func (m *MergeSort) mergeInPlace(arr []int, left, mid, right int) {
	// Copia para o buffer
	for i := left; i <= right; i++ {
		m.buffer[i] = arr[i]
	}

	// Índices para as duas metades
	i := left    // primeira metade
	j := mid + 1 // segunda metade
	k := left    // array resultado

	// Merge principal
	for i <= mid && j <= right {
		if m.buffer[i] <= m.buffer[j] {
			arr[k] = m.buffer[i]
			i++
		} else {
			arr[k] = m.buffer[j]
			j++
		}
		k++
	}

	// Copia elementos restantes da primeira metade
	for i <= mid {
		arr[k] = m.buffer[i]
		i++
		k++
	}
	// Nota: elementos da segunda metade já estão no lugar correto
}

// Função utilitária para encontrar o mínimo entre dois inteiros
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
