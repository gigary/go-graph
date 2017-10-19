package graph

import (
	"math"
)

const (
	MAX_INT = math.MaxInt32
)

type (
	ShortestPathCalculator struct {
		Graph Graph
		From  Vertex
		To    Vertex

		vertices  Vertices
		distances map[Vertex]Distance
	}
)

func NewShortestPathCalculator(g Graph, from Vertex, to Vertex) *ShortestPathCalculator {
	return &ShortestPathCalculator{
		Graph: g,
		From:  from,
		To:    to,

		distances: map[Vertex]Distance{},
	}
}

func (c *ShortestPathCalculator) CalculateByAlgorithm(algorithmName string) Distance {
	c.vertices = GetAllVertices(c.Graph)

	switch algorithmName {
	case "Bellman-Ford":
		c.calculateByBellmanFord()
	default:
		c.calculateByDijkstra()
	}

	return c.distances[c.To]
}

func (c *ShortestPathCalculator) calculateByDijkstra() {
	verticesToProcess := map[Vertex]bool{}
	for _, v := range c.vertices {
		verticesToProcess[v] = true
		c.distances[v] = MAX_INT
		if w, exist := c.Graph[c.From][v]; exist {
			c.distances[v] = w
		}
	}
	minVertex := struct {
		distance Distance
		name     Vertex
	}{}
	for len(verticesToProcess) > 0 {
		minVertex.distance = MAX_INT
		for v1 := range verticesToProcess {
			if c.distances[v1] < minVertex.distance || minVertex.distance == MAX_INT {
				minVertex.distance = c.distances[v1]
				minVertex.name = v1
			}
		}
		for v2, w := range c.Graph[minVertex.name] {
			if _, exist := verticesToProcess[v2]; exist {
				if c.distances[v2] > minVertex.distance+w {
					c.distances[v2] = minVertex.distance + w
				}
			}
		}
		delete(verticesToProcess, minVertex.name)
	}
}

func (c *ShortestPathCalculator) calculateByBellmanFord() {
	for _, v := range c.vertices {
		c.distances[v] = MAX_INT
		if w, exist := c.Graph[c.From][v]; exist {
			c.distances[v] = w
		}
	}

	for i := 0; i < len(c.vertices); i++ {
		for v1, n := range c.Graph {
			for v2, w := range n {
				subDistance := c.distances[v1]
				if v1 == c.From {
					subDistance = w
				}
				if c.distances[v2]-w > subDistance {
					c.distances[v2] = c.distances[v1] + w
				}
			}
		}
	}
}
