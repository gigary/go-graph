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

	minVertex struct {
		distance Distance
		name     Vertex
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
	verticesToProcess := c.initVerticesToProcess()
	c.initDistancesForAllVertices()
	for len(verticesToProcess) > 0 {
		minVertex := c.findMinVertex(verticesToProcess)
		for neighbour, weight := range c.Graph[minVertex.name] {
			if !c.isVertexProcessed(neighbour, verticesToProcess) {
				c.updateDistanceIfFoundBetterValue(neighbour, minVertex.name, weight)
			}
		}
		delete(verticesToProcess, minVertex.name)
	}
}

func (c *ShortestPathCalculator) isVertexProcessed(vertex Vertex, verticesToProcess map[Vertex]bool) bool {
	_, exist := verticesToProcess[vertex]
	return !exist
}

func (c *ShortestPathCalculator) initVerticesToProcess() map[Vertex]bool {
	verticesToProcess := map[Vertex]bool{}
	for _, v := range c.vertices {
		verticesToProcess[v] = true
	}
	return verticesToProcess
}

func (c *ShortestPathCalculator) findMinVertex(verticesToProcess map[Vertex]bool) minVertex {
	minVertex := minVertex{
		distance: MAX_INT,
	}
	for v1 := range verticesToProcess {
		if c.distances[v1] < minVertex.distance || minVertex.distance == MAX_INT {
			minVertex.distance = c.distances[v1]
			minVertex.name = v1
		}
	}
	return minVertex
}

func (c *ShortestPathCalculator) calculateByBellmanFord() {
	c.initDistancesForAllVertices()
	for i := 0; i < len(c.vertices); i++ {
		for v1, n := range c.Graph {
			for v2, w := range n {
				c.updateDistanceIfFoundBetterValue(v2, v1, w)
			}
		}
	}
}

func (c *ShortestPathCalculator) initDistancesForAllVertices() {
	for _, v := range c.vertices {
		c.distances[v] = MAX_INT
		if w, exist := c.Graph[c.From][v]; exist {
			c.distances[v] = w
		}
	}
}

func (c *ShortestPathCalculator) updateDistanceIfFoundBetterValue(
	current Vertex,
	previous Vertex,
	weight Distance,
) {
	currentDistance := c.distances[previous]
	if previous == c.From {
		currentDistance = weight
	}
	if c.distances[current]-weight > currentDistance {
		c.distances[current] = c.distances[previous] + weight
	}
}
