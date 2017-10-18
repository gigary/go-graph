package routes_with_limited_edges

import (
	"testing"
	"github.com/gigary/go-graph-algorithms/graph"
)

func TestCountRoutesBetweenTwoVerticesWithMaxEdges2(t *testing.T) {
	expected := 4
	actual := CountRoutesBetweenTwoVerticesWithMaxEdge(graph.GetTestGraph(), "A", "E", 3)
	if actual != expected {
		t.Errorf("Counting is not correct.\nExpected: %v\nGot: %v", expected, actual)
	}
}

func TestCountRoutesBetweenTwoVerticesWithMaxEdges1(t *testing.T) {
	expected := 0
	actual := CountRoutesBetweenTwoVerticesWithMaxEdge(graph.GetTestGraph(), "D", "A", 10)
	if actual != expected {
		t.Errorf("Counting is not correct.\nExpected: %v\nGot: %v", expected, actual)
	}
}

func TestCountRoutesBetweenTwoVerticesWithMaxEdges(t *testing.T) {
	expected := 2
	actual := CountRoutesBetweenTwoVerticesWithMaxEdge(graph.GetTestGraph(), "C", "C", 3)
	if actual != expected {
		t.Errorf("Counting is not correct.\nExpected: %v\nGot: %v", expected, actual)
	}
}
