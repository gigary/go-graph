package graph

type (
	vertexToProcess struct {
		distance Distance
		vertex   Vertex
	}

	MaxDistanceRouteCounter struct {
		Graph Graph
		From  Vertex
		To    Vertex
		Limit Distance

		counter           int
		verticesToProcess []vertexToProcess
	}
)

func NewMaxDistanceRouteCounter(g Graph, from Vertex, to Vertex, maxDistance Distance) *MaxDistanceRouteCounter {
	return &MaxDistanceRouteCounter{
		Graph: g,
		From:  from,
		To:    to,
		Limit: maxDistance,

		counter: 0,
		verticesToProcess: []vertexToProcess{
			{distance: 0, vertex: from},
		},
	}
}

func (c *MaxDistanceRouteCounter) Count() int {
	for len(c.verticesToProcess) > 0 {
		v := c.popFirstVertexFromTheQueue()
		for neighbor, distance := range c.Graph[v.vertex] {
			newDistance := v.distance + distance
			if c.isDistanceSatisfied(newDistance) {
				c.queueNeighborVertexToProcessLater(neighbor, newDistance)
				c.increaseCounterIfSeeTargetVertex(neighbor)
			}
		}
	}
	return c.counter
}

func (c *MaxDistanceRouteCounter) popFirstVertexFromTheQueue() vertexToProcess {
	v := c.verticesToProcess[0]
	c.verticesToProcess = c.verticesToProcess[1:]
	return v
}

func (c *MaxDistanceRouteCounter) isDistanceSatisfied(newDistance Distance) bool {
	return newDistance < c.Limit
}

func (c *MaxDistanceRouteCounter) queueNeighborVertexToProcessLater(neighbor Vertex, newDistance Distance) {
	c.verticesToProcess = append(c.verticesToProcess, vertexToProcess{
		distance: newDistance,
		vertex:   neighbor,
	})
}

func (c *MaxDistanceRouteCounter) increaseCounterIfSeeTargetVertex(v Vertex) {
	if v == c.To {
		c.counter++
	}
}
