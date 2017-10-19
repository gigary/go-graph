package graph

import (
	"testing"
)

func TestShortestRouteCalculatorByDijstraAlgorithm(t *testing.T) {
	samples := []struct {
		from     Vertex
		to       Vertex
		expected Distance
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
		actual := NewShortestPathCalculator(getTestGraph(), s.from, s.to).CalculateByAlgorithm("Dijkstra")
		if actual != s.expected {
			t.Errorf("Invalid shortest route.\nExpected: %v\nGot: %v", s.expected, actual)
		}
	}
}

func TestShortestRouteCalculatorByBellmanFordAlgorithm(t *testing.T) {
	samples := []struct {
		from     Vertex
		to       Vertex
		expected Distance
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
		actual := NewShortestPathCalculator(getTestGraph(), s.from, s.to).CalculateByAlgorithm("Bellman-Ford")
		if actual != s.expected {
			t.Errorf("Invalid shortest route.\nExpected: %v\nGot: %v", s.expected, actual)
		}
	}
}
