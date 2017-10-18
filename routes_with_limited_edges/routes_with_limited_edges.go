package routes_with_limited_edges

import (
	"github.com/gigary/go-graph-algorithms/graph"
)

type (
	RouteCount map[graph.Vertex]map[graph.Vertex]map[int]int
)

func buildAllRouteCountWithMaxEdge(g graph.Graph, maxEdge int) RouteCount {
	c := RouteCount{}
	vertices := graph.GetAllVertices(g)
	for _, begin := range vertices {
		c[begin] = map[graph.Vertex]map[int]int{}
		for _, end := range vertices {
			c[begin][end] = map[int]int{}
		}
	}

	for i := 0; i <= maxEdge; i++ {
		for _, begin := range vertices {
			for _, end := range vertices {
				c[begin][end][i] = 0
				if i == 0 {
					n := 0
					if end == begin {
						n = 1
					}
					c[begin][end][i] += n
				} else {
					for _, v := range vertices {
						_, exist := g[v][end]
						if exist {
							c[begin][end][i] += c[begin][v][i-1]
						}
					}
				}
			}
		}
	}
	return c
}
