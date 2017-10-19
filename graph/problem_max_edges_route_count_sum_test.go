package graph

import (
	"testing"
)

func TestMaxEdgesRouteCountSumWithMaxEdges3(t *testing.T) {
	expected := 4
	actual := NewMaxEdgeRouteCountSum(getTestGraph(), "A", "E", 3).Sum()
	if actual != expected {
		t.Errorf("Counting is not correct.\nExpected: %v\nGot: %v", expected, actual)
	}
}

func TestMaxEdgesRouteCountSumWithMaxEdges2(t *testing.T) {
	expected := 0
	actual := NewMaxEdgeRouteCountSum(getTestGraph(), "D", "A", 10).Sum()
	if actual != expected {
		t.Errorf("Counting is not correct.\nExpected: %v\nGot: %v", expected, actual)
	}
}

func TestMaxEdgesRouteCountSumWithMaxEdges1(t *testing.T) {
	expected := 2
	actual := NewMaxEdgeRouteCountSum(getTestGraph(), "C", "C", 3).Sum()
	if actual != expected {
		t.Errorf("Counting is not correct.\nExpected: %v\nGot: %v", expected, actual)
	}
}
