package routes_with_max_distance

import (
	"github.com/gigary/go-graph-algorithms/graph"
	"testing"
)

func TestCountRoutesWithMaxDistance3(t *testing.T) {
	expected := 0
	actual := CountRoutesWithMaxDistance(graph.GetTestGraph(), "E", "A", 100)
	if actual != expected {
		t.Errorf("Count route with max distance wrongly.\nExpected: %v\nGot: %v", expected, actual)
	}
}

func TestCountRoutesWithMaxDistance2(t *testing.T) {
	expected := 3
	actual := CountRoutesWithMaxDistance(graph.GetTestGraph(), "A", "C", 15)
	if actual != expected {
		t.Errorf("Count route with max distance wrongly.\nExpected: %v\nGot: %v", expected, actual)
	}
}

func TestCountRoutesWithMaxDistance(t *testing.T) {
	expected := 7
	actual := CountRoutesWithMaxDistance(graph.GetTestGraph(), "C", "C", 30)
	if actual != expected {
		t.Errorf("Count route with max distance wrongly.\nExpected: %v\nGot: %v", expected, actual)
	}
}
