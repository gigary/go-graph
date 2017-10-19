package graph

import (
	"reflect"
	"testing"
)

func TestParseProblemConstraintInput(t *testing.T) {
	samples := []struct {
		input    string
		expected ProblemConstraint
	}{
		{
			input: "C-C-3",
			expected: ProblemConstraint{
				From:      Vertex("C"),
				To:        Vertex("C"),
				Condition: 3,
				input: "C-C-3",
			},
		},
		{
			input: "A-C-4",
			expected: ProblemConstraint{
				From:      Vertex("A"),
				To:        Vertex("C"),
				Condition: 4,
				input: "A-C-4",
			},
		},
		{
			input: "A-C",
			expected: ProblemConstraint{
				From: Vertex("A"),
				To:   Vertex("C"),
				input: "A-C",
			},
		},
		{
			input: "B-B",
			expected: ProblemConstraint{
				From: Vertex("B"),
				To:   Vertex("B"),
				input: "B-B",
			},
		},
		{
			input: "C-C-30",
			expected: ProblemConstraint{
				From:      Vertex("C"),
				To:        Vertex("C"),
				Condition: 30,
				input: "C-C-30",
			},
		},
	}

	for _, s := range samples {
		actual := ParseProblemConstraintInput(s.input)
		if !reflect.DeepEqual(actual, s.expected) {
			t.Errorf("Wrongly parsed problem constrant input.\nExpected: %v\nGot: %v", s.expected, actual)
		}
	}
}
