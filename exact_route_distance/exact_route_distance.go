package exact_route_distance

import (
	"errors"
	"github.com/gigary/go-graph-algorithms/graph"
)

const (
	ERROR_NO_ROUTE = "NO SUCH ROUTE"
)

func Calculate(g graph.Graph, edges []graph.Edge) (graph.Distance, error) {
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
