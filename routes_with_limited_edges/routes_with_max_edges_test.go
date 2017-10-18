package routes_with_limited_edges

import (
	"testing"
	"github.com/gigary/go-graph-algorithms/graph"
)

func TestCountRoutesBetweenTwoVerticesWithMaxEdges2(t *testing.T) {
	expected := 4
	actual := CountRoutesBetweenTwoVerticesWithMaxEdge(graph.ParseGraphInput("AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7"), "A", "E", 3)
	if actual != expected {
		t.Errorf("Counting is not correct.\nExpected: %v\nGot: %v", expected, actual)
	}
}

func TestCountRoutesBetweenTwoVerticesWithMaxEdges1(t *testing.T) {
	expected := 0
	actual := CountRoutesBetweenTwoVerticesWithMaxEdge(graph.ParseGraphInput("AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7"), "D", "A", 10)
	if actual != expected {
		t.Errorf("Counting is not correct.\nExpected: %v\nGot: %v", expected, actual)
	}
}

func TestCountRoutesBetweenTwoVerticesWithMaxEdges(t *testing.T) {
	expected := 2
	actual := CountRoutesBetweenTwoVerticesWithMaxEdge(graph.ParseGraphInput("AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7"), "C", "C", 3)
	if actual != expected {
		t.Errorf("Counting is not correct.\nExpected: %v\nGot: %v", expected, actual)
	}
}
