package routes_with_max_edges

import (
	"testing"
	"reflect"
)

func TestCountRoutesBetweenTwoVerticesWithMaxEdges2(t *testing.T) {
	expected := 4
	actual := CountRoutesBetweenTwoVerticesWithMaxEdge(getTestGraph(), "A", "E", 3)
	if actual != expected {
		t.Errorf("Counting is not correct.\nExpected: %v\nGot: %v", expected, actual)
	}
}

func TestCountRoutesBetweenTwoVerticesWithMaxEdges1(t *testing.T) {
	expected := 0
	actual := CountRoutesBetweenTwoVerticesWithMaxEdge(getTestGraph(), "D", "A", 10)
	if actual != expected {
		t.Errorf("Counting is not correct.\nExpected: %v\nGot: %v", expected, actual)
	}
}

func TestCountRoutesBetweenTwoVerticesWithMaxEdges(t *testing.T) {
	expected := 2
	actual := CountRoutesBetweenTwoVerticesWithMaxEdge(getTestGraph(), "C", "C", 3)
	if actual != expected {
		t.Errorf("Counting is not correct.\nExpected: %v\nGot: %v", expected, actual)
	}
}

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
	actual := buildAllRouteCountWithMaxEdge(getTestGraph(), 2)
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
	actual := buildAllRouteCountWithMaxEdge(getTestGraph(), 1)
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
	actual := buildAllRouteCountWithMaxEdge(getTestGraph(), 0)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Counting is not correct.\nExpected: %v\nGot: %v", expected, actual)
	}
}

func TestGetAllVertices(t *testing.T) {
	expected := []Vertex{"A", "B", "C", "D", "E"}
	actual := getAllVertices(getTestGraph())
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
