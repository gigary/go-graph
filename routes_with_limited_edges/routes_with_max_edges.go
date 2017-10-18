package routes_with_limited_edges

func CountRoutesBetweenTwoVerticesWithMaxEdge(g Graph, from Vertex, to Vertex, maxEdge int) int {
	all := buildAllRouteCountWithMaxEdge(g, maxEdge)
	total := 0
	for i := 1; i <= maxEdge; i++ {
		total += all[from][to][i]
	}
	return total
}
