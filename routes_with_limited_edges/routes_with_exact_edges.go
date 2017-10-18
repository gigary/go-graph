package routes_with_limited_edges

func CountRoutesBetweenTwoVerticesWithExactNumberOfEdges(g Graph, from Vertex, to Vertex, numberOfEdges int) int {
	all := buildAllRouteCountWithMaxEdge(g, numberOfEdges)
	return all[from][to][numberOfEdges]
}
