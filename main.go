package main

import (
	"fmt"

	"helltaker-solver/core"
	"helltaker-solver/solver"
)

func main() {
	level, err := core.ParseRawLevelData("data/level1.txt")
	if err != nil {
		panic("Failed to parse level data: " + err.Error())
	}

	fmt.Printf("Level data loaded. Available moves: %d\n", level.MovesLeft)
	fmt.Printf("Player at: (%d, %d)\n", level.PlayerPos.Y, level.PlayerPos.X)
	fmt.Println("--------------------------------------")

	moves, found := solver.Solve(&level)
	if found {
		fmt.Printf("Found solution in %d moves.\n", len(moves))
		fmt.Print("Sequence: ")
		solver.PrintSolution(moves)
	} else {
		fmt.Println("No solution found.")
	}
}
