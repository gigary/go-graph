package graph

import (
	"strconv"
	"strings"
)

type ProblemConstraint struct {
	From      Vertex
	To        Vertex
	Condition int64

	input string
}

func ParseProblemConstraintInput(input string) ProblemConstraint {
	constraint := ProblemConstraint{input: input}
	constraint.parseInput()
	return constraint
}

func (c *ProblemConstraint) parseInput() {
	parts := strings.Split(c.input, "-")
	c.From = Vertex(strings.Trim(parts[0], " "))
	c.To = Vertex(strings.Trim(parts[1], " "))

	if len(parts) == 3 {
		condition, err := strconv.ParseInt(strings.Trim(parts[2], " "), 10, 32)
		if err != nil {
			panic("Invalid test input value")
		}
		c.Condition = condition
	}
}
