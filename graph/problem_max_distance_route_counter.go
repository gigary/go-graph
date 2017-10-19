package graph

type (
	MaxDistanceRouteCounter struct {
		Graph Graph
		From  Vertex
		To    Vertex
		Limit Distance
	}
)

func NewMaxDistanceRouteCounter(g Graph, from Vertex, to Vertex, maxDistance Distance) *MaxDistanceRouteCounter {
	return &MaxDistanceRouteCounter{
		Graph: g,
		From:  from,
		To:    to,
		Limit: maxDistance,
	}
}

func (c *MaxDistanceRouteCounter) Count() int {
	r := 0
	type vertexToProcess struct {
		distance Distance
		vertex   Vertex
	}
	verticesToProcess := []vertexToProcess{
		{distance: 0, vertex: c.From},
	}
	for len(verticesToProcess) > 0 {
		v := verticesToProcess[0]
		verticesToProcess = verticesToProcess[1:]
		for neighbor, distance := range c.Graph[v.vertex] {
			if v.distance+distance < c.Limit {
				verticesToProcess = append(verticesToProcess, vertexToProcess{
					distance: v.distance + distance,
					vertex:   neighbor,
				})
				if neighbor == c.To {
					r++
				}
			}
		}
	}
	return r
}
