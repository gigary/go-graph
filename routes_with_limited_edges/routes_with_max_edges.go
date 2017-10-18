package routes_with_limited_edges

import (
	"github.com/gigary/go-graph-algorithms/graph"
)

func CountRoutesBetweenTwoVerticesWithMaxEdge(g graph.Graph, from graph.Vertex, to graph.Vertex, maxEdge int) int {
	all := buildAllRouteCountWithMaxEdge(g, maxEdge)
	total := 0
	for i := 1; i <= maxEdge; i++ {
		total += all[from][to][i]
	}
	return total
}
