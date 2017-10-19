package graph

import (
	"errors"
)

const (
	ERROR_NO_ROUTE = "NO SUCH ROUTE"
)

type (
	ExactRouteDistanceCalculator struct {
		Graph Graph
		Route []Edge
	}
)

func NewExactRouteDistanceCalculator(g Graph, edges []Edge) *ExactRouteDistanceCalculator {
	return &ExactRouteDistanceCalculator{
		Graph: g,
		Route: edges,
	}
}

func (calculator *ExactRouteDistanceCalculator) Calculate() (Distance, error) {
	distance := Distance(0)
	for _, edge := range calculator.Route {
		w, exist := calculator.Graph[edge[0]][edge[1]]
		if !exist {
			return 0, errors.New(ERROR_NO_ROUTE)
		}
		distance += w
	}
	return distance, nil
}
