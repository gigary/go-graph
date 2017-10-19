package graph

import (
	"sort"
	"strconv"
	"strings"
)

type (
	Vertex   string
	Distance int
	Route    string

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

func ParseGraphInput(input string) Graph {
	g := Graph{}
	edges := strings.Split(input, ",")
	for _, edge := range edges {
		edge := strings.Trim(edge, " ")
		parts := strings.Split(edge, "")
		value, ok := strconv.ParseInt(parts[2], 10, 32)
		if ok != nil {
			panic("Invalid graph input")
		}
		if g[Vertex(parts[0])] == nil {
			g[Vertex(parts[0])] = map[Vertex]Distance{}
		}
		g[Vertex(parts[0])][Vertex(parts[1])] = Distance(value)
	}
	return g
}

func ParseEdgesFromRouteInput(r Route) []Edge {
	edges := []Edge{}
	parts := strings.Split(string(r), "-")
	for i := range parts {
		if i < len(parts)-1 {
			edges = append(edges, Edge{Vertex(parts[i]), Vertex(parts[i+1])})
		}
	}
	return edges
}
