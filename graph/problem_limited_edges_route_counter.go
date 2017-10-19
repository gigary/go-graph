package graph

type (
	RouteCounts map[Vertex]map[Vertex]map[int]int

	LimitedEdgesRouteCounter struct {
		Graph    Graph
		Vertices Vertices
		Limit    int
		Counts   RouteCounts
	}
)

func NewLimitedEdgesRouteCounter(g Graph, maxEdge int) *LimitedEdgesRouteCounter {
	return &LimitedEdgesRouteCounter{
		Graph:    g,
		Vertices: Vertices{},
		Limit:    maxEdge,
	}
}

func (c *LimitedEdgesRouteCounter) Count() RouteCounts {
	c.Vertices = GetAllVertices(c.Graph)
	c.initAllCounts()

	for numberOfEdges := 0; numberOfEdges <= c.Limit; numberOfEdges++ {
		for _, begin := range c.Vertices {
			for _, end := range c.Vertices {
				if numberOfEdges == 0 {
					c.initCountsForZeroEdge(begin, end)
				} else {
					c.countRecursively(begin, end, numberOfEdges)
				}
			}
		}
	}
	return c.Counts
}

func (c *LimitedEdgesRouteCounter) initAllCounts() {
	c.Counts = RouteCounts{}
	for _, begin := range c.Vertices {
		c.Counts[begin] = map[Vertex]map[int]int{}
		for _, end := range c.Vertices {
			c.Counts[begin][end] = map[int]int{}
			for numberOfEdges := 0; numberOfEdges <= c.Limit; numberOfEdges++ {
				c.Counts[begin][end][numberOfEdges] = 0
			}
		}
	}
}

func (c *LimitedEdgesRouteCounter) initCountsForZeroEdge(begin Vertex, end Vertex) {
	n := 0
	if end == begin {
		n = 1
	}
	c.Counts[begin][end][0] += n
}

func (c *LimitedEdgesRouteCounter) countRecursively(begin Vertex, end Vertex, numberOfEdges int) {
	for _, v := range c.Vertices {
		_, exist := c.Graph[v][end]
		if exist {
			c.Counts[begin][end][numberOfEdges] += c.Counts[begin][v][numberOfEdges-1]
		}
	}
}
