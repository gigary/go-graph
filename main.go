package main

import (
	"github.com/gigary/go-graph-algorithms/graph"
	"os"
	"bufio"
	"fmt"
)

func main() {
	textFile := os.Args[1]
	file, err := os.Open(textFile)
	if err != nil {
		panic("Unable to open the file!")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		solution := graph.NewProblemSolver(scanner.Text()).Solve()
		fmt.Println(fmt.Sprintf("%+v", solution))
	}
}
