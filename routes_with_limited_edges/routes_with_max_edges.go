package routes_with_limited_edges

import (
	graph "github.com/gigary/go-graph-algorithms"
)

func CountRoutesBetweenTwoVerticesWithMaxEdge(g graph.Graph, from graph.Vertex, to graph.Vertex, maxEdge int) int {
	all := buildAllRouteCountWithMaxEdge(g, maxEdge)
	total := 0
	for i := 1; i <= maxEdge; i++ {
		total += all[from][to][i]
	}
	return total
}
