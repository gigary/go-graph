package exact_route_distance

import (
	"testing"
	"reflect"
)

func TestCalculationWithInvalidRoute(t *testing.T) {
	_, err := Calculate(getTestGraph(), "A-E-D")
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
	cases := map[Route]Weight{
		"A-B-C":     9,
		"A-D":       5,
		"A-D-C":     13,
		"A-E-B-C-D": 22,
	}
	g := getTestGraph()
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
	route := Route("A-B-C")
	expected := []Edge{
		{"A", "B"},
		{"B", "C"},
	}
	actual := getEdgesFromRoute(route)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Wrongly parsed route.\nExpected: %v\nGot: %v", expected, actual)
	}
}

func getTestGraph() Graph {
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

