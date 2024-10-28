package analysis

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"sort-bench/internal/core"
)

// Estrutura para o visualizador da árvore
type TreeVisualizer struct {
	tree *core.RecursionTree
}

type D3Node struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Children  []D3Node `json:"children,omitempty"`
	Depth     int      `json:"depth"`
	Data      []int    `json:"data"`
	Operation string   `json:"operation"`
}

func NewTreeVisualizer(tree *core.RecursionTree) *TreeVisualizer {
	return &TreeVisualizer{tree: tree}
}

func (tv *TreeVisualizer) buildD3Tree() D3Node {
	nodeMap := make(map[string]*D3Node)
	var root *D3Node

	// Criar nós da árvore
	for _, node := range tv.tree.Nodes {
		d3node := &D3Node{
			ID:        node.NodeID,
			Name:      fmt.Sprintf("Node %s", node.NodeID),
			Depth:     node.Depth,
			Data:      node.Value,
			Operation: node.Operation,
			Children:  make([]D3Node, 0),
		}
		nodeMap[node.NodeID] = d3node

		if node.ParentID == "" {
			root = d3node
		}
	}

	// Estabelecer relações pai-filho
	for _, node := range tv.tree.Nodes {
		if node.ParentID != "" {
			if parent, exists := nodeMap[node.ParentID]; exists {
				parent.Children = append(parent.Children, *nodeMap[node.NodeID])
			}
		}
	}

	if root == nil {
		return D3Node{Name: "Empty Tree"}
	}
	return *root
}

func (tv *TreeVisualizer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/tree/data":
		tv.serveTreeData(w, r)
	default:
		tv.serveVisualization(w, r)
	}
}

func (tv *TreeVisualizer) serveTreeData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tv.buildD3Tree())
}

func (tv *TreeVisualizer) serveVisualization(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("visualization").Parse(visualizationTemplate))
	tmpl.Execute(w, nil)
}

var visualizationTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>MergeSort Recursion Tree</title>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/d3/7.8.5/d3.min.js"></script>
    <style>
        .node circle {
            fill: #fff;
            stroke: steelblue;
            stroke-width: 3px;
        }
        .node text {
            font: 12px sans-serif;
        }
        .link {
            fill: none;
            stroke: #ccc;
            stroke-width: 2px;
        }
        .tooltip {
            position: absolute;
            padding: 8px;
            background: rgba(0, 0, 0, 0.8);
            color: white;
            border-radius: 4px;
            font-size: 12px;
            pointer-events: none;
        }
    </style>
</head>
<body>
    <div id="tree"></div>
    <script>
        var width = window.innerWidth;
        var height = window.innerHeight;
        var margin = { top: 40, right: 90, bottom: 50, left: 90 };

        var svg = d3.select("#tree")
            .append("svg")
            .attr("width", width)
            .attr("height", height)
            .append("g")
            .attr("transform", "translate(" + margin.left + "," + margin.top + ")");

        var tooltip = d3.select("body")
            .append("div")
            .attr("class", "tooltip")
            .style("opacity", 0);

        d3.json("/tree/data").then(function(data) {
            var tree = d3.tree().size([height - margin.top - margin.bottom, width - margin.left - margin.right]);

            var root = d3.hierarchy(data);
            tree(root);

            var link = svg.selectAll(".link")
                .data(root.links())
                .enter().append("path")
                .attr("class", "link")
                .attr("d", d3.linkHorizontal()
                    .x(function(d) { return d.y; })
                    .y(function(d) { return d.x; }));

            var node = svg.selectAll(".node")
                .data(root.descendants())
                .enter().append("g")
                .attr("class", "node")
                .attr("transform", function(d) {
                    return "translate(" + d.y + "," + d.x + ")";
                });

            node.append("circle")
                .attr("r", 10)
                .on("mouseover", function(event, d) {
                    tooltip.transition()
                        .duration(200)
                        .style("opacity", 0.9);
                    tooltip.html(
                        "Operation: " + d.data.operation + "<br/>" +
                        "Data: [" + d.data.data.join(", ") + "]<br/>" +
                        "Depth: " + d.data.depth
                    )
                    .style("left", (event.pageX + 10) + "px")
                    .style("top", (event.pageY - 28) + "px");
                })
                .on("mouseout", function() {
                    tooltip.transition()
                        .duration(500)
                        .style("opacity", 0);
                });

            node.append("text")
                .attr("dy", ".35em")
                .attr("x", function(d) { return d.children ? -13 : 13; })
                .style("text-anchor", function(d) { return d.children ? "end" : "start"; })
                .text(function(d) { return d.data.operation; });
        });
    </script>
</body>
</html>
`

