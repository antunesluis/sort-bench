package analysis

import (
	"fmt"
	"github.com/goccy/go-graphviz"
)

type TreeVisualizer struct {
	graph *graphviz.Graph
}

func NewTreeVisualizer() *TreeVisualizer {
	g := graphviz.New()
	graph, _ := g.Graph()
	return &TreeVisualizer{graph: graph}
}

func (v *TreeVisualizer) GenerateRecursionTreeImage(tree *RecursionTree, outputPath string) error {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		return fmt.Errorf("failed to create graph: %w", err)
	}
	defer graph.Close()

	// Create nodes
	nodes := make(map[string]*graphviz.Node)
	for _, info := range tree.Nodes {
		label := fmt.Sprintf("%s\n%v", info.Operation, info.Value)
		node, err := graph.CreateNode(info.NodeID)
		if err != nil {
			return fmt.Errorf("failed to create node: %w", err)
		}
		node.SetLabel(label)

		// Style based on operation
		switch info.Operation {
		case "split":
			node.SetColor("blue")
		case "merge":
			node.SetColor("green")
		case "leaf":
			node.SetColor("gray")
		}

		nodes[info.NodeID] = node

		// Create edge to parent if exists
		if info.ParentID != "" {
			if parent, exists := nodes[info.ParentID]; exists {
				_, err := graph.CreateEdge("", parent, node)
				if err != nil {
					return fmt.Errorf("failed to create edge: %w", err)
				}
			}
		}
	}

	// Save to file
	if err := g.RenderFilename(graph, graphviz.PNG, outputPath); err != nil {
		return fmt.Errorf("failed to render graph: %w", err)
	}

	return nil
}
