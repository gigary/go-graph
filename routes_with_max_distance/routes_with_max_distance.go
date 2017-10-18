package routes_with_max_distance

import (
	"github.com/gigary/go-graph-algorithms/graph"
)

func CountRoutesWithMaxDistance(g graph.Graph, from graph.Vertex, to graph.Vertex, maxDistance graph.Distance) int {
	r := 0
	type vertexToProcess struct {
		distance graph.Distance
		vertex   graph.Vertex
	}
	verticesToProcess := []vertexToProcess{
		{distance: 0, vertex: from},
	}
	for len(verticesToProcess) > 0 {
		v := verticesToProcess[0]
		verticesToProcess = verticesToProcess[1:]
		for neighbor, distance := range g[v.vertex] {
			if v.distance+distance < maxDistance {
				verticesToProcess = append(verticesToProcess, vertexToProcess{
					distance: v.distance + distance,
					vertex:   neighbor,
				})
				if neighbor == to {
					r++
				}
			}
		}
	}
	return r

}
