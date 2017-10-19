package graph

type (
	RouteCounts map[Vertex]map[Vertex]map[int]int

	LimitedEdgesRouteCounter struct {
		Graph Graph
		Limit int
	}
)

func NewLimitedEdgesRouteCounter(g Graph, maxEdge int) *LimitedEdgesRouteCounter {
	return &LimitedEdgesRouteCounter{
		Graph: g,
		Limit: maxEdge,
	}
}

func (c *LimitedEdgesRouteCounter) Count() RouteCounts {
	vertices := GetAllVertices(c.Graph)
	counts := RouteCounts{}
	for _, begin := range vertices {
		counts[begin] = map[Vertex]map[int]int{}
		for _, end := range vertices {
			counts[begin][end] = map[int]int{}
		}
	}

	for i := 0; i <= c.Limit; i++ {
		for _, begin := range vertices {
			for _, end := range vertices {
				counts[begin][end][i] = 0
				if i == 0 {
					n := 0
					if end == begin {
						n = 1
					}
					counts[begin][end][i] += n
				} else {
					for _, v := range vertices {
						_, exist := c.Graph[v][end]
						if exist {
							counts[begin][end][i] += counts[begin][v][i-1]
						}
					}
				}
			}
		}
	}
	return counts
}
