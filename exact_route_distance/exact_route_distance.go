package exact_route_distance

import (
	"errors"
	"strings"
)

const (
	ERROR_NO_ROUTE = "NO SUCH ROUTE"
)

type (
	Vertex string
	Weight int
	Route  string

	Edge  [2]Vertex
	Graph map[Vertex]map[Vertex]Weight
)

func Calculate(g Graph, r Route) (Weight, error) {
	edges := getEdgesFromRoute(r)
	distance := Weight(0)
	for _, e := range edges {
		w, exist := g[e[0]][e[1]]
		if !exist {
			return 0, errors.New(ERROR_NO_ROUTE)
		}
		distance += w
	}
	return distance, nil
}

func getEdgesFromRoute(r Route) []Edge {
	edges := []Edge{}
	parts := strings.Split(string(r), "-")
	for i := range parts {
		if i < len(parts)-1 {
			edges = append(edges, Edge{Vertex(parts[i]), Vertex(parts[i+1])})
		}
	}
	return edges
}
