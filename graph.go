package go_graph_algorithms

import "sort"

type (
	Vertex string
	Distance int
	Route  string

	Vertices []Vertex
	Edge     [2]Vertex
	Graph    map[Vertex]map[Vertex]Distance
)

func GetAllVertices(g Graph) []Vertex {
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

func GetTestGraph() Graph {
	return Graph{
		"A": {
			"B": 5,
			"D": 5,
			"E": 7,
		},
		"B": {
			"C": 4,
		},
		"C": {
			"D": 8,
			"E": 2,
		},
		"D": {
			"C": 8,
			"E": 6,
		},
		"E": {
			"B": 3,
		},
	}
}
