package graph

import (
	"testing"
)

func TestExactEdgesRouteCounterWithExactEdges3(t *testing.T) {
	expected := 0
	actual := NewExactEdgesRouteCounter(getTestGraph(), "D", "A", 5).Count()
	if actual != expected {
		t.Errorf("Counting is not correct.\nExpected: %v\nGot: %v", expected, actual)
	}
}

func TestExactEdgesRouteCounterWithExactEdges2(t *testing.T) {
	expected := 2
	actual := NewExactEdgesRouteCounter(getTestGraph(), "A", "E", 3).Count()
	if actual != expected {
		t.Errorf("Counting is not correct.\nExpected: %v\nGot: %v", expected, actual)
	}
}

func TestExactEdgesRouteCounterWithExactEdges1(t *testing.T) {
	expected := 3
	actual := NewExactEdgesRouteCounter(getTestGraph(), "A", "C", 4).Count()
	if actual != expected {
		t.Errorf("Counting is not correct.\nExpected: %v\nGot: %v", expected, actual)
	}
}
