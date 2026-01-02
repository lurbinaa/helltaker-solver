package solver

import (
	"fmt"
	"maps"
	"sort"
	"strings"

	"helltaker-solver/core"
)

func Solve(level *core.Level) ([]core.Direction, bool) {
	visited := make(map[string]bool)

	queue := []LevelSnapshot{{
		Level: CloneLevel(level),
		Moves: []core.Direction{},
	}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		key := SerializeLevel(current.Level)
		if visited[key] {
			continue
		}
		visited[key] = true

		if current.Level.CheckWin() {
			return current.Moves, true
		}

		if current.Level.MovesLeft <= 0 {
			continue
		}

		for _, ds := range current.Level.CheckAllAvailableMoves() {
			level := CloneLevel(current.Level)
			action := level.HandleInput(ds)

			if action == core.Win {
				return append(current.Moves, ds), true
			}

			moves := make([]core.Direction, len(current.Moves)+1)
			copy(moves, current.Moves)
			moves[len(current.Moves)] = ds

			queue = append(queue, LevelSnapshot{
				Level: level,
				Moves: moves,
			})
		}
	}

	return nil, false
}

// serialize the level into a string with sorted coordinates
// to check for repeated sequences
func SerializeLevel(l *core.Level) string {
	keys := make([]core.Point, 0, len(l.Tiles))
	for p := range l.Tiles {
		keys = append(keys, p)
	}

	sort.Slice(keys, func(i, j int) bool {
		if keys[i].Y != keys[j].Y {
			return keys[i].Y < keys[j].Y
		}
		return keys[i].X < keys[j].X
	})

	var sb strings.Builder
	for _, p := range keys {
		fmt.Fprintf(&sb, "%d,%d,%v|", p.Y, p.X, l.Tiles[p])
	}
	return sb.String()
}

func CloneLevel(l *core.Level) *core.Level {
	tiles := make(core.Tiles)
	maps.Copy(tiles, l.Tiles)

	return &core.Level{
		PlayerPos:             l.PlayerPos,
		MovesLeft:             l.MovesLeft,
		SpecialItems:          l.SpecialItems,
		SpecialItemsCollected: l.SpecialItemsCollected,
		Tiles:                 tiles,
	}
}

func PrintSolution(moves []core.Direction) {
	for i, m := range moves {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(Arrows[m])
	}
	fmt.Println()
}
