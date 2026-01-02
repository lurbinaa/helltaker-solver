package solver

import "helltaker-solver/core"

type LevelSnapshot struct {
	Level *core.Level
	Moves []core.Direction
}

var Arrows = map[core.Direction]string{
	core.Up:    "↑",
	core.Right: "→",
	core.Down:  "↓",
	core.Left:  "←",
}
