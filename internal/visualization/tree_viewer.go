package visualization

import (
	"fmt"
	"github.com/goccy/go-graphviz"
)

// MergeNode representa um nó no processo de merge
type MergeNode struct {
	Left     []int
	Right    []int
	Result   []int
	Children []*MergeNode
}

// MergeVisualizer gera visualizações do processo de merge
type MergeVisualizer struct {
	root *MergeNode
}

func NewMergeVisualizer() *MergeVisualizer {
	return &MergeVisualizer{}
}

// TrackMerge registra uma operação de merge
func (m *MergeVisualizer) TrackMerge(left, right, result []int) {
	node := &MergeNode{
		Left:   make([]int, len(left)),
		Right:  make([]int, len(right)),
		Result: make([]int, len(result)),
	}
	copy(node.Left, left)
	copy(node.Right, right)
	copy(node.Result, result)

	if m.root == nil {
		m.root = node
	} else {
		// Encontra o nó pai apropriado e adiciona como filho
		parent := m.findParentNode(m.root, left, right)
		if parent != nil {
			parent.Children = append(parent.Children, node)
		}
	}
}

// findParentNode encontra o nó pai apropriado baseado nos arrays de entrada
func (m *MergeVisualizer) findParentNode(current *MergeNode, left, right []int) *MergeNode {
	// Se o resultado atual contém todos os elementos dos arrays de entrada
	// este é o nó pai apropriado
	resultMap := make(map[int]bool)
	for _, v := range current.Result {
		resultMap[v] = true
	}

	allFound := true
	for _, v := range left {
		if !resultMap[v] {
			allFound = false
			break
		}
	}
	for _, v := range right {
		if !resultMap[v] {
			allFound = false
			break
		}
	}

	if allFound {
		return current
	}

	// Procura recursivamente nos filhos
	for _, child := range current.Children {
		if found := m.findParentNode(child, left, right); found != nil {
			return found
		}
	}

	return nil
}

// GenerateVisualization cria uma visualização do processo de merge
func (m *MergeVisualizer) GenerateVisualization(filename string) error {
	g := graphviz.New()
	defer g.Close()

	graph, err := g.Graph()
	if err != nil {
		return err
	}
	defer graph.Close()

	// Configuração do grafo
	graph.SetRankDir("TB") // Top to Bottom layout

	// Adiciona nós e arestas recursivamente
	if m.root != nil {
		m.addNodeToGraph(graph, m.root, "")
	}

	// Salva o resultado
	if err := g.RenderFilename(graph, graphviz.PNG, filename); err != nil {
		return err
	}

	return nil
}

// addNodeToGraph adiciona um nó e suas conexões ao grafo
func (m *MergeVisualizer) addNodeToGraph(graph *graphviz.Graph, node *MergeNode, parentID string) string {
	// Cria identificador único para o nó
	nodeID := fmt.Sprintf("node_%p", node)

	// Cria o nó no grafo
	n, _ := graph.CreateNode(nodeID)

	// Configura aparência do nó
	n.SetShape("record")
	label := fmt.Sprintf("{Left: %v|Right: %v|Result: %v}", node.Left, node.Right, node.Result)
	n.SetLabel(label)
	n.SetStyle("filled")
	n.SetFillColor("lightblue")

	// Conecta ao pai se existir
	if parentID != "" {
		parent, _ := graph.Node(parentID)
		graph.CreateEdge("", parent, n)
	}

	// Processa filhos recursivamente
	for _, child := range node.Children {
		m.addNodeToGraph(graph, child, nodeID)
	}

	return nodeID
}
