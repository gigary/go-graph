package graph

import (
	"testing"
	"reflect"
)

func TestGetAllVertices(t *testing.T) {
	expected := []Vertex{"A", "B", "C", "D", "E"}
	actual := GetAllVertices(GetTestGraph())
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Wrongly parsed route.\nExpected: %v\nGot: %v", expected, actual)
	}
}
