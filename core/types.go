package core

type OccupiedState int
type Direction int
type Action int

const (
	Empty OccupiedState = iota
	Player
	Box
	BoxSpecialItem
	SpecialItem
	Skeleton
	Goal
)

const (
	Up Direction = iota
	Right
	Down
	Left
)

const (
	Move Action = iota
	// Does nothing
	PunchBox
	PushBox
	AttackSkeleton
	PushSkeleton
	SpecialItemCollect
	Win
	OutOfMoves
	Unknown
)

type Point struct {
	Y, X int
}

type Level struct {
	PlayerPos             Point
	MovesLeft             int
	SpecialItems          int
	SpecialItemsCollected int
	Tiles                 Tiles
}

type Tiles = map[Point]OccupiedState
type Directions = map[Direction]Point

var RawStateToSymbols = map[rune]OccupiedState{
	'_': Empty,
	'p': Player,
	'b': Box,
	'B': BoxSpecialItem,
	'i': SpecialItem,
	's': Skeleton,
	'g': Goal,
}

var OccupiedStates = map[OccupiedState]string{
	Empty:    "Empty",
	Player:   "Player",
	Box:      "Box",
	Skeleton: "Skeleton",
	Goal:     "Goal",
}

var DirectionOffsets = Directions{
	Up:    {-1, 0},
	Right: {0, 1},
	Down:  {1, 0},
	Left:  {0, -1},
}
