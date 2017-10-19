package graph

import (
	"reflect"
	"testing"
)

func TestProblemSover(t *testing.T) {
	input := "AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7|A-B-C|A-D|A-D-C|A-E-B-C-D|A-E-D|C-C-3|A-C-4|A-C|B-B|C-C-30"
	expected := Solution{
		Output1:  9,
		Output2:  5,
		Output3:  13,
		Output4:  22,
		Output5:  "NO SUCH ROUTE",
		Output6:  2,
		Output7:  3,
		Output8:  9,
		Output9:  9,
		Output10: 7,
	}
	actual := NewProblemSolver(input).Solve()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Incorrect problem solution.\nExpected: %v\nGot: %v", expected, actual)
	}
}
