package shortest_route

import (
	"testing"
	graph "github.com/gigary/go-graph-algorithms"
)

func TestCalculateShortestRouteV2(t *testing.T) {
	samples := []struct {
		from     graph.Vertex
		to       graph.Vertex
		expected graph.Distance
	}{
		{
			from:     "A",
			to:       "C",
			expected: 9,
		},
		{
			from:     "B",
			to:       "B",
			expected: 9,
		},
	}

	for _, s := range samples {
		actual := CalculateShortestRouteV2(getTestGraph(), s.from, s.to)
		if actual != s.expected {
			t.Errorf("Invalid shortest rout.\nExpected: %v\nGot: %v", s.expected, actual)
		}
	}
}

func TestCalculateShortestRouteV1(t *testing.T) {
	samples := []struct {
		from     graph.Vertex
		to       graph.Vertex
		expected graph.Distance
	}{
		{
			from:     "A",
			to:       "C",
			expected: 9,
		},
		{
			from:     "B",
			to:       "B",
			expected: 9,
		},
	}

	for _, s := range samples {
		actual := CalculateShortestRouteV1(getTestGraph(), s.from, s.to)
		if actual != s.expected {
			t.Errorf("Invalid shortest rout.\nExpected: %v\nGot: %v", s.expected, actual)
		}
	}
}

func getTestGraph() graph.Graph {
	return graph.Graph{
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
}
