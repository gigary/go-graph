package routes_with_limited_edges

import (
	graph "github.com/gigary/go-graph-algorithms"
	"reflect"
	"testing"
)

func TestBuildRouteCountWithTwoEdge(t *testing.T) {
	expected := RouteCount{
		"A": {
			"A": {0: 1, 1: 0, 2: 0},
			"B": {0: 0, 1: 1, 2: 1},
			"C": {0: 0, 1: 0, 2: 2},
			"D": {0: 0, 1: 1, 2: 0},
			"E": {0: 0, 1: 1, 2: 1},
		},
		"B": {
			"A": {0: 0, 1: 0, 2: 0},
			"B": {0: 1, 1: 0, 2: 0},
			"C": {0: 0, 1: 1, 2: 0},
			"D": {0: 0, 1: 0, 2: 1},
			"E": {0: 0, 1: 0, 2: 1},
		},
		"C": {
			"A": {0: 0, 1: 0, 2: 0},
			"B": {0: 0, 1: 0, 2: 1},
			"C": {0: 1, 1: 0, 2: 1},
			"D": {0: 0, 1: 1, 2: 0},
			"E": {0: 0, 1: 1, 2: 1},
		},
		"D": {
			"A": {0: 0, 1: 0, 2: 0},
			"B": {0: 0, 1: 0, 2: 1},
			"C": {0: 0, 1: 1, 2: 0},
			"D": {0: 1, 1: 0, 2: 1},
			"E": {0: 0, 1: 1, 2: 1},
		},
		"E": {
			"A": {0: 0, 1: 0, 2: 0},
			"B": {0: 0, 1: 1, 2: 0},
			"C": {0: 0, 1: 0, 2: 1},
			"D": {0: 0, 1: 0, 2: 0},
			"E": {0: 1, 1: 0, 2: 0},
		},
	}
	actual := buildAllRouteCountWithMaxEdge(graph.GetTestGraph(), 2)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Counting is not correct.\nExpected: %v\nGot: %v", expected, actual)
	}
}

func TestBuildRouteCountWithOneEdge(t *testing.T) {
	expected := RouteCount{
		"A": {
			"A": {0: 1, 1: 0},
			"B": {0: 0, 1: 1},
			"C": {0: 0, 1: 0},
			"D": {0: 0, 1: 1},
			"E": {0: 0, 1: 1},
		},
		"B": {
			"A": {0: 0, 1: 0},
			"B": {0: 1, 1: 0},
			"C": {0: 0, 1: 1},
			"D": {0: 0, 1: 0},
			"E": {0: 0, 1: 0},
		},
		"C": {
			"A": {0: 0, 1: 0},
			"B": {0: 0, 1: 0},
			"C": {0: 1, 1: 0},
			"D": {0: 0, 1: 1},
			"E": {0: 0, 1: 1},
		},
		"D": {
			"A": {0: 0, 1: 0},
			"B": {0: 0, 1: 0},
			"C": {0: 0, 1: 1},
			"D": {0: 1, 1: 0},
			"E": {0: 0, 1: 1},
		},
		"E": {
			"A": {0: 0, 1: 0},
			"B": {0: 0, 1: 1},
			"C": {0: 0, 1: 0},
			"D": {0: 0, 1: 0},
			"E": {0: 1, 1: 0},
		},
	}
	actual := buildAllRouteCountWithMaxEdge(graph.GetTestGraph(), 1)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Counting is not correct.\nExpected: %v\nGot: %v", expected, actual)
	}
}

func TestBuildRouteCountWithNoEdge(t *testing.T) {
	expected := RouteCount{
		"A": {
			"A": {0: 1},
			"B": {0: 0},
			"C": {0: 0},
			"D": {0: 0},
			"E": {0: 0},
		},
		"B": {
			"A": {0: 0},
			"B": {0: 1},
			"C": {0: 0},
			"D": {0: 0},
			"E": {0: 0},
		},
		"C": {
			"A": {0: 0},
			"B": {0: 0},
			"C": {0: 1},
			"D": {0: 0},
			"E": {0: 0},
		},
		"D": {
			"A": {0: 0},
			"B": {0: 0},
			"C": {0: 0},
			"D": {0: 1},
			"E": {0: 0},
		},
		"E": {
			"A": {0: 0},
			"B": {0: 0},
			"C": {0: 0},
			"D": {0: 0},
			"E": {0: 1},
		},
	}
	actual := buildAllRouteCountWithMaxEdge(graph.GetTestGraph(), 0)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Counting is not correct.\nExpected: %v\nGot: %v", expected, actual)
	}
}
