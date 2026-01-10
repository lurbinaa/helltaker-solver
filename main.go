package main

import (
	"fmt"

	"helltaker-solver/core"
	"helltaker-solver/solver"
)

func main() {
	level, err := core.ParseRawLevelData("data/level4.txt")
	if err != nil {
		panic("Failed to parse level data: " + err.Error())
	}

	fmt.Printf("Level data loaded. Available moves: %d\n", level.MovesLeft)
	fmt.Printf("Player at: (%d, %d)\n", level.PlayerPos.Y, level.PlayerPos.X)

	solutions, iters := solver.Solve(&level)
	if len(solutions) > 0 {
		for _, sol := range solutions {
			core.PrintLine()
			fmt.Print("Sequence: ")
			solver.PrintSolution(sol.Moves)
			new, _ := core.ParseRawLevelData("data/level4.txt")
			core.DebugMovements(&new, sol.Moves)
		}

		core.PrintLine()

		fmt.Printf(
			"Found %d solutions in %d iterations.\n",
			len(solutions),
			iters,
		)
	} else {
		core.PrintLine()

		fmt.Printf("No solution found in %d iterations.\n", iters)
	}
}
