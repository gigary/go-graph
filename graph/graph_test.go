package graph

import (
	"reflect"
	"testing"
)

func TestParseGraphInput(t *testing.T) {
	input := "AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7"
	expected := Graph{
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
	actual := ParseGraphInput(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Wrongly parsed graph input.\nExpected: %v\nGot: %v", expected, actual)
	}
}

func TestGetAllVertices(t *testing.T) {
	expected := []Vertex{"A", "B", "C", "D", "E"}
	actual := GetAllVertices(ParseGraphInput("AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7"))
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Wrongly parsed route.\nExpected: %v\nGot: %v", expected, actual)
	}
}

func TestParseRoute(t *testing.T) {
	route := Route("A-B-C")
	expected := []Edge{
		{"A", "B"},
		{"B", "C"},
	}
	actual := ParseEdgesFromRouteInput(route)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Wrongly parsed route.\nExpected: %v\nGot: %v", expected, actual)
	}
}
