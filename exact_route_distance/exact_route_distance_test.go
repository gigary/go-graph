package exact_route_distance

import (
	"reflect"
	"testing"
	"github.com/gigary/go-graph-algorithms/graph"
)

func TestCalculationWithInvalidRoute(t *testing.T) {
	_, err := Calculate(graph.GetTestGraph(), "A-E-D")
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
	g := graph.GetTestGraph()
	for r, expected := range cases {
		actual, err := Calculate(g, r)
		if err != nil {
			t.Errorf("Calculation should NOT return error if route is valid")
		}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Invalid calculation.\nExpected: %v\nGot: %v", expected, actual)
		}
	}
}

func TestParseRoute(t *testing.T) {
	route := graph.Route("A-B-C")
	expected := []graph.Edge{
		{"A", "B"},
		{"B", "C"},
	}
	actual := getEdgesFromRoute(route)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Wrongly parsed route.\nExpected: %v\nGot: %v", expected, actual)
	}
}
