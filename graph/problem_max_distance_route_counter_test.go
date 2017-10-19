package graph

import (
	"testing"
)

func TestMaxDistanceRouteCounterWithMaxDistance3(t *testing.T) {
	expected := 0
	actual := NewMaxDistanceRouteCounter(getTestGraph(), "E", "A", 100).Count()
	if actual != expected {
		t.Errorf("Count route with max distance wrongly.\nExpected: %v\nGot: %v", expected, actual)
	}
}

func TestMaxDistanceRouteCounterWithMaxDistance2(t *testing.T) {
	expected := 3
	actual := NewMaxDistanceRouteCounter(getTestGraph(), "A", "C", 15).Count()
	if actual != expected {
		t.Errorf("Count route with max distance wrongly.\nExpected: %v\nGot: %v", expected, actual)
	}
}

func TestMaxDistanceRouteCounterWithMaxDistance1(t *testing.T) {
	expected := 7
	actual := NewMaxDistanceRouteCounter(getTestGraph(), "C", "C", 30).Count()
	if actual != expected {
		t.Errorf("Count route with max distance wrongly.\nExpected: %v\nGot: %v", expected, actual)
	}
}
