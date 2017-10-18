package exact_route_distance

import (
	"errors"
	"strings"
	"github.com/gigary/go-graph-algorithms/graph"
)

const (
	ERROR_NO_ROUTE = "NO SUCH ROUTE"
)

func Calculate(g graph.Graph, r graph.Route) (graph.Distance, error) {
	edges := getEdgesFromRoute(r)
	distance := graph.Distance(0)
	for _, e := range edges {
		w, exist := g[e[0]][e[1]]
		if !exist {
			return 0, errors.New(ERROR_NO_ROUTE)
		}
		distance += w
	}
	return distance, nil
}

func getEdgesFromRoute(r graph.Route) []graph.Edge {
	edges := []graph.Edge{}
	parts := strings.Split(string(r), "-")
	for i := range parts {
		if i < len(parts)-1 {
			edges = append(edges, graph.Edge{graph.Vertex(parts[i]), graph.Vertex(parts[i+1])})
		}
	}
	return edges
}
