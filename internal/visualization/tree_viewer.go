package visualization

import (
	"fmt"
	"os"

	"github.com/goccy/go-graphviz"
)

type TreeVisualizer struct {
	nodes    []Node
	edges    []Edge
	maxDepth int
}

type Node struct {
	ID       int
	Value    []int
	Depth    int
	IsMerged bool
}

type Edge struct {
	From int
	To   int
}

func NewTreeVisualizer() *TreeVisualizer {
	return &TreeVisualizer{
		nodes:    make([]Node, 0),
		edges:    make([]Edge, 0),
		maxDepth: 0,
	}
}

// TrackRecursion registra uma operação de divisão ou merge
func (t *TreeVisualizer) TrackRecursion(arr []int, depth int, isMerged bool) int {
	nodeID := len(t.nodes)
	t.nodes = append(t.nodes, Node{
		ID:       nodeID,
		Value:    arr,
		Depth:    depth,
		IsMerged: isMerged,
	})

	if depth > t.maxDepth {
		t.maxDepth = depth
	}

	return nodeID
}

// AddEdge adiciona uma conexão entre dois nós
func (t *TreeVisualizer) AddEdge(from, to int) {
	t.edges = append(t.edges, Edge{From: from, To: to})
}

// GenerateVisualization cria uma visualização da árvore de recursão
func (t *TreeVisualizer) GenerateVisualization(filename string) error {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		return err
	}
	defer graph.Close()

	// Criar nós
	graphNodes := make(map[int]*graphviz.Node)
	for _, node := range t.nodes {
		label := fmt.Sprintf("%v", node.Value)
		if len(label) > 20 {
			label = label[:17] + "..."
		}

		n, err := graph.CreateNode(fmt.Sprintf("%d", node.ID))
		if err != nil {
			return err
		}

		n.SetLabel(label)
		if node.IsMerged {
			n.SetColor("green")
		}
		graphNodes[node.ID] = n
	}

	// Criar arestas
	for _, edge := range t.edges {
		_, err := graph.CreateEdge("", graphNodes[edge.From], graphNodes[edge.To])
		if err != nil {
			return err
		}
	}

	// Salvar o gráfico
	return g.RenderFilename(graph, graphviz.PNG, filename)
}

