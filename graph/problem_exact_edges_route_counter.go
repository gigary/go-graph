package graph

type (
	ExactEdgesRouteCounter struct {
		LimitedEdgesRouteCounter
		From Vertex
		To   Vertex
	}
)

func NewExactEdgesRouteCounter(g Graph, from Vertex, to Vertex, numberOfEdges int) *ExactEdgesRouteCounter {
	return &ExactEdgesRouteCounter{
		LimitedEdgesRouteCounter: LimitedEdgesRouteCounter{
			Graph: g,
			Limit: numberOfEdges,
		},
		From: from,
		To:   to,
	}
}

func (c *ExactEdgesRouteCounter) Count() int {
	counts := c.LimitedEdgesRouteCounter.Count()
	return counts[c.From][c.To][c.Limit]
}
