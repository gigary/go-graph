package exact_route_distance

import (
	"github.com/gigary/go-graph-algorithms/graph"
	"reflect"
	"testing"
)

func TestCalculationWithInvalidRoute(t *testing.T) {
	_, err := Calculate(graph.ParseGraphInput("AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7"), graph.ParseEdgesFromRouteInput("A-E-D"))
	if err == nil {
		t.Errorf("Calculation should return error if route is invalid")
	}
	expected := ERROR_NO_ROUTE
	actual := err.Error()
	if actual != expected {
		t.Errorf("Invalid calculation.\nExpected: %v\nGot: %v", expected, actual)
	}
}

func TestCalculateDistanceWithValidRoute(t *testing.T) {
	cases := map[graph.Route]graph.Distance{
		"A-B-C":     9,
		"A-D":       5,
		"A-D-C":     13,
		"A-E-B-C-D": 22,
	}
	g := graph.ParseGraphInput("AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7")
	for r, expected := range cases {
		actual, err := Calculate(g, graph.ParseEdgesFromRouteInput(r))
		if err != nil {
			t.Errorf("Calculation should NOT return error if route is valid")
		}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Invalid calculation.\nExpected: %v\nGot: %v", expected, actual)
		}
	}
}
