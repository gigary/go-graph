package routes_with_limited_edges

import (
	"github.com/gigary/go-graph-algorithms/graph"
)

func CountRoutesBetweenTwoVerticesWithExactNumberOfEdges(g graph.Graph, from graph.Vertex, to graph.Vertex, numberOfEdges int) int {
	all := buildAllRouteCountWithMaxEdge(g, numberOfEdges)
	return all[from][to][numberOfEdges]
}
