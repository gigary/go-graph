package graph

import "strings"

type (
	Problem struct {
		Graph
		Input1  []Edge
		Input2  []Edge
		Input3  []Edge
		Input4  []Edge
		Input5  []Edge
		Input6  ProblemConstraint
		Input7  ProblemConstraint
		Input8  ProblemConstraint
		Input9  ProblemConstraint
		Input10 ProblemConstraint
	}

	Solution struct {
		Output1  Distance
		Output2  Distance
		Output3  Distance
		Output4  Distance
		Output5  string
		Output6  int
		Output7  int
		Output8  Distance
		Output9  Distance
		Output10 int
	}

	ProblemSolver struct {
		input    string
		graph    Graph
		problem  Problem
		solution Solution
	}
)

func NewProblemSolver(input string) *ProblemSolver {
	return &ProblemSolver{
		input:    input,
		solution: Solution{},
	}
}

func (s *ProblemSolver) Solve() Solution {
	s.parseInput()
	s.solution.Output1, _ = NewExactRouteDistanceCalculator(s.problem.Graph, s.problem.Input1).Calculate()
	s.solution.Output2, _ = NewExactRouteDistanceCalculator(s.problem.Graph, s.problem.Input2).Calculate()
	s.solution.Output3, _ = NewExactRouteDistanceCalculator(s.problem.Graph, s.problem.Input3).Calculate()
	s.solution.Output4, _ = NewExactRouteDistanceCalculator(s.problem.Graph, s.problem.Input4).Calculate()
	_, err := NewExactRouteDistanceCalculator(s.problem.Graph, s.problem.Input5).Calculate()
	s.solution.Output5 = err.Error()

	s.solution.Output6 = NewMaxEdgeRouteCountSum(
		s.problem.Graph,
		s.problem.Input6.From,
		s.problem.Input6.To,
		int(s.problem.Input6.Condition),
	).Sum()

	s.solution.Output7 = NewExactEdgesRouteCounter(
		s.problem.Graph,
		s.problem.Input7.From,
		s.problem.Input7.To,
		int(s.problem.Input7.Condition),
	).Count()

	s.solution.Output8 = NewShortestPathCalculator(s.problem.Graph, s.problem.Input8.From, s.problem.Input8.To).
		CalculateByAlgorithm("Bellman-Ford")

	s.solution.Output9 = NewShortestPathCalculator(s.problem.Graph, s.problem.Input9.From, s.problem.Input9.To).
		CalculateByAlgorithm("Dijkstra")

	s.solution.Output10 = NewMaxDistanceRouteCounter(
		s.problem.Graph,
		s.problem.Input10.From,
		s.problem.Input10.To,
		Distance(s.problem.Input10.Condition),
	).
		Count()

	return s.solution
}

func (s *ProblemSolver) parseInput() {
	parts := strings.Split(s.input, "|")
	s.problem = Problem{
		Graph:   ParseGraphInput(strings.Trim(parts[0], " ")),
		Input1:  ParseEdgesFromRouteInput(Route(parts[1])),
		Input2:  ParseEdgesFromRouteInput(Route(parts[2])),
		Input3:  ParseEdgesFromRouteInput(Route(parts[3])),
		Input4:  ParseEdgesFromRouteInput(Route(parts[4])),
		Input5:  ParseEdgesFromRouteInput(Route(parts[5])),
		Input6:  ParseProblemConstraintInput(parts[6]),
		Input7:  ParseProblemConstraintInput(parts[7]),
		Input8:  ParseProblemConstraintInput(parts[8]),
		Input9:  ParseProblemConstraintInput(parts[9]),
		Input10: ParseProblemConstraintInput(parts[10]),
	}
}
