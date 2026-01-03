package main

import (
	"fmt"

	"helltaker-solver/core"
	"helltaker-solver/solver"
)

func main() {
	level, err := core.ParseRawLevelData("data/level3.txt")
	if err != nil {
		panic("Failed to parse level data: " + err.Error())
	}

	fmt.Printf("Level data loaded. Available moves: %d\n", level.MovesLeft)
	fmt.Printf("Player at: (%d, %d)\n", level.PlayerPos.Y, level.PlayerPos.X)
	fmt.Println("--------------------------------------")

	moves, found, iters := solver.Solve(&level)
	if found {
		fmt.Printf(
			"Found solution in %d moves. %d iterations in total.\n",
			len(moves),
			iters,
		)
		fmt.Print("Sequence: ")
		solver.PrintSolution(moves)
	} else {
		fmt.Printf("No solution found in %d iterations.\n", iters)
	}
}
