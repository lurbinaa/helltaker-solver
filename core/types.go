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
	SwitchHazardEven
	SwitchHazardOdd
	BoxSwitchHazardEven
	BoxSwitchHazardOdd
	SpecialItem
	BoxSpecialItem
	Key
	BoxKey
	Chest
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
	TouchHazardPushBox
	TouchHazardPunchBox
	TouchHazardAttackSkeleton
	TouchHazardPushSkeleton
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
	MovesCount            int
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
	'e': SwitchHazardEven,
	'o': SwitchHazardOdd,
	'E': BoxSwitchHazardEven,
	'O': BoxSwitchHazardOdd,
	'i': SpecialItem,
	'I': BoxSpecialItem,
	's': Skeleton,
	'k': Key,
	'K': BoxKey,
	'c': Chest,
	'g': Goal,
}

var SymbolToRawState = map[OccupiedState]rune{
	Empty:               '_',
	Player:              'p',
	Box:                 'b',
	Hazard:              'h',
	BoxHazard:           'H',
	SwitchHazardEven:    'e',
	SwitchHazardOdd:     'o',
	BoxSwitchHazardEven: 'E',
	BoxSwitchHazardOdd:  'O',
	SpecialItem:         'i',
	BoxSpecialItem:      'I',
	Key:                 'k',
	BoxKey:              'K',
	Chest:               'c',
	Skeleton:            's',
	Goal:                'g',
}

var OccupiedStateNames = map[OccupiedState]string{
	Empty:               "Empty",
	Player:              "Player",
	Box:                 "Box",
	Hazard:              "Hazard",
	BoxHazard:           "BoxHazard",
	SwitchHazardEven:    "SwitchHazardEven",
	SwitchHazardOdd:     "SwitchHazardOdd",
	BoxSwitchHazardEven: "BoxSwitchHazardEven",
	BoxSwitchHazardOdd:  "BoxSwitchHazardOdd",
	SpecialItem:         "SpecialItem",
	BoxSpecialItem:      "BoxSpecialItem",
	BoxKey:              "BoxKey",
	Key:                 "Key",
	Chest:               "Chest",
	Skeleton:            "Skeleton",
	Goal:                "Goal",
}

var ActionNames = map[Action]string{
	Move:                      "Move",
	PushBox:                   "PushBox",
	PunchBox:                  "PunchBox",
	AttackSkeleton:            "AttackSkeleton",
	PushSkeleton:              "PushSkeleton",
	TouchHazard:               "TouchHazard",
	TouchHazardPushBox:        "TouchHazardPushBox",
	TouchHazardPunchBox:       "TouchHazardPunchBox",
	TouchHazardAttackSkeleton: "TouchHazardAttackSkeleton",
	TouchHazardPushSkeleton:   "TouchHazardPushSkeleton",
	SpecialItemCollect:        "SpecialItemCollect",
	CollectKey:                "CollectKey",
	OpenChest:                 "OpenChest",
	PunchChest:                "PunchChest",
	Win:                       "Win",
	Unknown:                   "Unknown",
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
