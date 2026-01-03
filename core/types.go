package core

type OccupiedState int
type Direction int
type Action int

const (
	Empty OccupiedState = iota
	Player
	Box
	Hazard
	BoxHazard
	SpecialItem
	BoxSpecialItem
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
	TouchHazard
	SpecialItemCollect
	Win
	OutOfMoves
	Unknown
)

type Point struct {
	Y, X int
}

type Tiles map[Point]OccupiedState
type Directions map[Direction]Point
type Level struct {
	PlayerPos             Point
	UnderPlayer           OccupiedState
	MovesLeft             int
	SpecialItems          int
	SpecialItemsCollected int
	Tiles                 Tiles
}

var RawStateToSymbol = map[rune]OccupiedState{
	'_': Empty,
	'p': Player,
	'b': Box,
	'h': Hazard,
	'H': BoxHazard,
	'i': SpecialItem,
	'I': BoxSpecialItem,
	's': Skeleton,
	'g': Goal,
}

var SymbolToRawState = map[OccupiedState]rune{
	Empty:          '_',
	Player:         'p',
	Box:            'b',
	Hazard:         'h',
	BoxHazard:      'H',
	SpecialItem:    'i',
	BoxSpecialItem: 'I',
	Skeleton:       's',
	Goal:           'g',
}

var OccupiedStateNames = map[OccupiedState]string{
	Empty:          "Empty",
	Player:         "Player",
	Box:            "Box",
	Hazard:         "Hazard",
	BoxHazard:      "BoxHazard",
	BoxSpecialItem: "BoxSpecialItem",
	SpecialItem:    "SpecialItem",
	Skeleton:       "Skeleton",
	Goal:           "Goal",
}

var DirectionNames = map[Direction]string{
	Up:    "Up",
	Right: "Right",
	Down:  "Down",
	Left:  "Left",
}

var DirectionOffsets = Directions{
	Up:    {-1, 0},
	Right: {0, 1},
	Down:  {1, 0},
	Left:  {0, -1},
}
