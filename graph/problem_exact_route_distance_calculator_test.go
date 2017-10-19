package graph

import (
	"reflect"
	"testing"
)

func TestExactRouteDistanceCalculatorWithInvalidRoute(t *testing.T) {
	_, err := NewExactRouteDistanceCalculator(getTestGraph(), ParseEdgesFromRouteInput("A-E-D")).Calculate()
	if err == nil {
		t.Errorf("Calculation should return error if route is invalid")
	}
	expected := ERROR_NO_ROUTE
	actual := err.Error()
	if actual != expected {
		t.Errorf("Invalid calculation.\nExpected: %v\nGot: %v", expected, actual)
	}
}

func TestExactRouteDistanceCalculatorWithValidRoute(t *testing.T) {
	cases := map[Route]Distance{
		"A-B-C":     9,
		"A-D":       5,
		"A-D-C":     13,
		"A-E-B-C-D": 22,
	}
	g := getTestGraph()
	for r, expected := range cases {
		actual, err := NewExactRouteDistanceCalculator(g, ParseEdgesFromRouteInput(r)).Calculate()
		if err != nil {
			t.Errorf("Calculation should NOT return error if route is valid")
		}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Invalid calculation.\nExpected: %v\nGot: %v", expected, actual)
		}
	}
}
