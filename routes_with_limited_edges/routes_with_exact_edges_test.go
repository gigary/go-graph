package routes_with_limited_edges

import (
	"github.com/gigary/go-graph-algorithms/graph"
	"testing"
)

func TestCountRoutesFromSourceToDestinationWithExactEdges2(t *testing.T) {
	expected := 0
	actual := CountRoutesBetweenTwoVerticesWithExactNumberOfEdges(graph.ParseGraphInput("AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7"), "D", "A", 5)
	if actual != expected {
		t.Errorf("Counting is not correct.\nExpected: %v\nGot: %v", expected, actual)
	}
}

func TestCountRoutesFromSourceToDestinationWithExactEdges1(t *testing.T) {
	expected := 2
	actual := CountRoutesBetweenTwoVerticesWithExactNumberOfEdges(graph.ParseGraphInput("AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7"), "A", "E", 3)
	if actual != expected {
		t.Errorf("Counting is not correct.\nExpected: %v\nGot: %v", expected, actual)
	}
}

func TestCountRoutesFromSourceToDestinationWithExactEdges(t *testing.T) {
	expected := 3
	actual := CountRoutesBetweenTwoVerticesWithExactNumberOfEdges(graph.ParseGraphInput("AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7"), "A", "C", 4)
	if actual != expected {
		t.Errorf("Counting is not correct.\nExpected: %v\nGot: %v", expected, actual)
	}
}
