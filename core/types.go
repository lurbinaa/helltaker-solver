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
	Key
	Chest
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
	PushBox
	PunchBox
	AttackSkeleton
	PushSkeleton
	TouchHazard
	SpecialItemCollect
	CollectKey
	PunchChest
	OpenChest
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
	KeyCollected          bool
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
	'k': Key,
	'c': Chest,
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
	Key:            'k',
	Chest:          'c',
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
	Key:            "Key",
	Chest:          "Chest",
	Goal:           "Goal",
}

var ActionNames = map[Action]string{
	Move:               "Move",
	PushBox:            "PushBox",
	PunchBox:           "PunchBox",
	AttackSkeleton:     "AttackSkeleton",
	PushSkeleton:       "PushSkeleton",
	TouchHazard:        "TouchHazard",
	SpecialItemCollect: "SpecialItemCollect",
	CollectKey:         "CollectKey",
	OpenChest:          "OpenChest",
	PunchChest:         "PunchChest",
	Win:                "Win",
	OutOfMoves:         "OutOfMoves",
	Unknown:            "Unknown",
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
