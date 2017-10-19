package graph

type (
	MaxEdgeRouteCountSum struct {
		LimitedEdgesRouteCounter
		From Vertex
		To   Vertex
	}
)

func NewMaxEdgeRouteCountSum(g Graph, from Vertex, to Vertex, maxEdge int) *MaxEdgeRouteCountSum {
	return &MaxEdgeRouteCountSum{
		LimitedEdgesRouteCounter: LimitedEdgesRouteCounter{
			Graph: g,
			Limit: maxEdge,
		},
		From: from,
		To:   to,
	}
}

func (s *MaxEdgeRouteCountSum) Sum() int {
	counts := s.Count()
	total := 0
	for i := 1; i <= s.Limit; i++ {
		total += counts[s.From][s.To][i]
	}
	return total
}
