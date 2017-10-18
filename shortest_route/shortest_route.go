package shortest_route

import (
	graph "github.com/gigary/go-graph-algorithms"
	"math"
)

const (
	MAX_INT = math.MaxInt32
)

func CalculateShortestRouteV2(g graph.Graph, from graph.Vertex, to graph.Vertex) graph.Distance {
	vertices := graph.GetAllVertices(g)
	d := map[graph.Vertex]graph.Distance{}

	calculateShortestRouteByDijkstra(vertices, d, g, from)

	return d[to]
}

func CalculateShortestRouteV1(g graph.Graph, from graph.Vertex, to graph.Vertex) graph.Distance {
	vertices := graph.GetAllVertices(g)
	d := map[graph.Vertex]graph.Distance{}

	calculateShortestRouteByBellmanFord(vertices, g, d, from)

	return d[to]
}

func calculateShortestRouteByDijkstra(vertices graph.Vertices, d map[graph.Vertex]graph.Distance, g graph.Graph, from graph.Vertex) {
	verticesToProcess := map[graph.Vertex]bool{}
	for _, v := range vertices {
		verticesToProcess[v] = true
		d[v] = MAX_INT
		if w, exist := g[from][v]; exist {
			d[v] = w
		}
	}
	minVertex := struct {
		distance graph.Distance
		name     graph.Vertex
	}{}
	for len(verticesToProcess) > 0 {
		minVertex.distance = MAX_INT
		for v1 := range verticesToProcess {
			if d[v1] < minVertex.distance || minVertex.distance == MAX_INT {
				minVertex.distance = d[v1]
				minVertex.name = v1
			}
		}
		for v2, w := range g[minVertex.name] {
			if _, exist := verticesToProcess[v2]; exist {
				if d[v2] > minVertex.distance+w {
					d[v2] = minVertex.distance + w
				}
			}
		}
		delete(verticesToProcess, minVertex.name)
	}
}

func calculateShortestRouteByBellmanFord(vertices graph.Vertices, g graph.Graph, d map[graph.Vertex]graph.Distance, from graph.Vertex) {
	for _, v := range vertices {
		d[v] = MAX_INT
		if w, exist := g[from][v]; exist {
			d[v] = w
		}
	}

	for i := 0; i < len(vertices); i++ {
		for v1, n := range g {
			for v2, w := range n {
				subDistance := d[v1]
				if v1 == from {
					subDistance = w
				}
				if d[v2]-w > subDistance {
					d[v2] = d[v1] + w
				}
			}
		}
	}
}
