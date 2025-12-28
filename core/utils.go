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
			state := RawStateToSymbols[char]
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

func PrintCoords(t Tiles) {
	for point, state := range t {
		fmt.Printf(
			"(%d, %d): %s\n",
			point.Y,
			point.X,
			OccupiedStates[state],
		)
	}
}
