package routes_with_limited_edges

import (
	"testing"
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
