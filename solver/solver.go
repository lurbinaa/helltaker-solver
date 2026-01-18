package solver

import (
	"fmt"
	"maps"
	"sort"
	"strings"

	"helltaker-solver/core"
)

func Solve(level *core.Level) (sols []Solution, iters uint) {
	visited := make(map[string]bool)
	queue := []LevelSnapshot{{
		Level: CloneLevel(level),
		Moves: []core.Direction{},
	}}
	iters = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		iters += 1

		key := SerializeLevel(current.Level)
		if visited[key] {
			continue
		}
		visited[key] = true

		if current.Level.MovesLeft <= 0 {
			continue
		}

		for _, d := range current.Level.CheckAllAvailableMoves() {
			level := CloneLevel(current.Level)
			action := level.HandleInput(d)

			moves := make([]core.Direction, len(current.Moves)+1)
			copy(moves, current.Moves)
			moves[len(current.Moves)] = d

			if action == core.Win {
				sols = append(sols, Solution{Moves: moves})
				continue
			}

			queue = append(queue, LevelSnapshot{
				Level: level,
				Moves: moves,
			})
		}
	}

	return sols, iters
}

// Serializes the level into a string with sorted coordinates
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

	// Include all state that affects gameplay uniqueness
	// Player position + what's under them is already in Tiles map
	// But UnderPlayer matters because it can be SpecialItem vs Empty
	fmt.Fprintf(&sb,
		"%v|%v|%v|%v|",
		l.KeyCollected,
		l.UnderPlayer,
		l.SpecialItemsCollected,
		l.MovesLeft,
	)
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
		UnderPlayer:           l.UnderPlayer,
		MovesLeft:             l.MovesLeft,
		SpecialItems:          l.SpecialItems,
		SpecialItemsCollected: l.SpecialItemsCollected,
		KeyCollected:          l.KeyCollected,
		MovesCount:            l.MovesCount,
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
