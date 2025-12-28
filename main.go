package main

import (
	"fmt"

	"helltaker-solver/core"
)

func main() {
	level, err := core.ParseRawLevelData("data/level1.txt")
	if err != nil {
		panic("Failed to parse level data: " + err.Error())
	}

	core.PrintCoords(level.Tiles)
	fmt.Println("Player:", level.MovesLeft)
}
