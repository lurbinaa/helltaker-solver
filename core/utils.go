package core

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseRawLevelData(path string) (l Level, err error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return l, fmt.Errorf("Failed to read level data: %w", err)
	}

	lines := strings.Split(string(data), "\n")
	m, err := strconv.Atoi(strings.TrimSpace(lines[0]))
	if err != nil {
		panic("Invalid level data: first line must be an integer")
	}

	l.MovesLeft = m
	l.Tiles = make(Tiles)
	for y, line := range lines[1:] {
		for x, char := range line {
			if char == ' ' {
				continue // out of bound
			}

			point := Point{Y: y, X: x}
			state := RawStateToSymbol[char]
			l.Tiles[point] = state

			switch state {
			case Player:
				l.PlayerPos = point
			case BoxSpecialItem:
				l.SpecialItems += 1
			}
		}
	}
	return l, nil
}

func PrintRawLevelData(l *Level) {
	// check borders
	maxY, maxX := 0, 0
	for p := range l.Tiles {
		if p.Y > maxY {
			maxY = p.Y
		}
		if p.X > maxX {
			maxX = p.X
		}
	}

	fmt.Println(l.MovesLeft)

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			state, exists := l.Tiles[Point{Y: y, X: x}]
			if !exists {
				fmt.Print(" ")
			} else {
				fmt.Print(string(SymbolToRawState[state]))
			}
		}
		fmt.Println()
	}
}

func PrintTiles(t Tiles) {
	for point, state := range t {
		fmt.Printf(
			"(%d, %d): %s\n",
			point.Y,
			point.X,
			OccupiedStateNames[state],
		)
	}
}
