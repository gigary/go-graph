package routes_with_max_distance

import (
	"github.com/gigary/go-graph-algorithms/graph"
	"testing"
)

func TestCountRoutesWithMaxDistance3(t *testing.T) {
	expected := 0
	actual := CountRoutesWithMaxDistance(graph.ParseGraphInput("AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7"), "E", "A", 100)
	if actual != expected {
		t.Errorf("Count route with max distance wrongly.\nExpected: %v\nGot: %v", expected, actual)
	}
}

func TestCountRoutesWithMaxDistance2(t *testing.T) {
	expected := 3
	actual := CountRoutesWithMaxDistance(graph.ParseGraphInput("AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7"), "A", "C", 15)
	if actual != expected {
		t.Errorf("Count route with max distance wrongly.\nExpected: %v\nGot: %v", expected, actual)
	}
}

func TestCountRoutesWithMaxDistance(t *testing.T) {
	expected := 7
	actual := CountRoutesWithMaxDistance(graph.ParseGraphInput("AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7"), "C", "C", 30)
	if actual != expected {
		t.Errorf("Count route with max distance wrongly.\nExpected: %v\nGot: %v", expected, actual)
	}
}
