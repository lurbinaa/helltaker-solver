// Core logic for recreating Helltaker game mechanics.
package core

func (l *Level) CalculateOffset(d Direction) (target Point) {
	offset := DirectionOffsets[d]
	return Point{l.PlayerPos.Y + offset.Y, l.PlayerPos.X + offset.X}
}

func (l *Level) CanPushTo(d Direction) bool {
	offset := DirectionOffsets[d]
	behind := Point{l.PlayerPos.Y + 2*offset.Y, l.PlayerPos.X + 2*offset.X}
	state, exists := l.Tiles[behind]
	return exists && state == Empty
}

func (l *Level) IsValidInput(d Direction) bool {
	target := l.CalculateOffset(d)
	_, exists := l.Tiles[target]
	return exists
}

// returns true if the player is adjacent to the goal
func (l *Level) CheckWin() bool {
	for d := range DirectionOffsets {
		target := l.CalculateOffset(d)
		if l.Tiles[target] == Goal {
			return true
		}
	}
	return false
}

func (l *Level) CheckAllAvailableMoves() (ds []Direction) {
	for d := range DirectionOffsets {
		if l.IsValidInput(d) {
			ds = append(ds, d)
		}
	}
	return ds
}

func (l *Level) MovePlayerTo(d Direction) {
	target := l.CalculateOffset(d)

	l.Tiles[l.PlayerPos] = Empty
	l.Tiles[target] = Player
	l.PlayerPos = target
}

func (l *Level) PushCollidable(d Direction, o OccupiedState) {
	offset := DirectionOffsets[d]
	current := l.CalculateOffset(d)
	// Tile behind
	target := Point{current.Y + offset.Y, current.X + offset.X}
	if o == BoxSpecialItem {
		l.Tiles[current] = SpecialItem
		l.Tiles[target] = Box
	} else {
		l.Tiles[current] = Empty
		l.Tiles[target] = o
	}
}

func (l *Level) AttackSkeleton(d Direction) {
	target := l.CalculateOffset(d)
	l.Tiles[target] = Empty
}

func (l *Level) HandleInput(d Direction) (a Action) {
	target := l.CalculateOffset(d)
	occupiedState := l.Tiles[target]

	switch occupiedState {
	case Empty:
		l.MovePlayerTo(d)
		if l.CheckWin() {
			a = Win
		} else {
			a = Move
		}
	case SpecialItem:
		a = SpecialItemCollect
		l.SpecialItemsCollected += 1
		l.MovePlayerTo(d)
	case Box, BoxSpecialItem:
		if l.CanPushTo(d) {
			a = PushBox
			l.PushCollidable(d, occupiedState)
		} else {
			a = PunchBox
			// Do nothing
		}
	case Skeleton:
		if l.CanPushTo(d) {
			a = PushSkeleton
			l.PushCollidable(d, occupiedState)
		} else {
			a = AttackSkeleton
			l.AttackSkeleton(d)
		}
	}
	l.MovesLeft -= 1
	return a
}
