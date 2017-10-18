package routes_with_limited_edges

import "sort"

type (
	Vertex string
	Weight int
	Route  string

	Vertices   []Vertex
	Edge       [2]Vertex
	Graph      map[Vertex]map[Vertex]Weight
	RouteCount map[Vertex]map[Vertex]map[int]int
)

func buildAllRouteCountWithMaxEdge(g Graph, maxEdge int) RouteCount {
	c := RouteCount{}
	vertices := getAllVertices(g)
	for _, begin := range vertices {
		c[begin] = map[Vertex]map[int]int{}
		for _, end := range vertices {
			c[begin][end] = map[int]int{}
		}
	}

	for i := 0; i <= maxEdge; i++ {
		for _, begin := range vertices {
			for _, end := range vertices {
				c[begin][end][i] = 0
				if i == 0 {
					n := 0
					if end == begin {
						n = 1
					}
					c[begin][end][i] += n
				} else {
					for _, v := range vertices {
						_, exist := g[v][end]
						if exist {
							c[begin][end][i] += c[begin][v][i-1]
						}
					}
				}
			}
		}
	}
	return c
}

func getAllVertices(g Graph) []Vertex {
	vertices := Vertices{}
	flags := map[Vertex]bool{}
	for begin, neighbors := range g {
		flags[begin] = true
		for end := range neighbors {
			flags[end] = true
		}
	}
	for v := range flags {
		vertices = append(vertices, v)
	}
	sort.Sort(vertices)
	return vertices
}

func (v Vertices) Len() int {
	return len(v)
}

func (v Vertices) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v Vertices) Less(i, j int) bool {
	return string(v[i]) < string(v[j])
}
