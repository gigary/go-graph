package routes_with_limited_edges

import (
	graph "github.com/gigary/go-graph-algorithms"
)

func CountRoutesBetweenTwoVerticesWithExactNumberOfEdges(g graph.Graph, from graph.Vertex, to graph.Vertex, numberOfEdges int) int {
	all := buildAllRouteCountWithMaxEdge(g, numberOfEdges)
	return all[from][to][numberOfEdges]
}
