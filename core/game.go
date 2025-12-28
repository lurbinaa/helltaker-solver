// Core logic for recreating Helltaker game mechanics.
package core

func (l *Level) CalculateOffset(d Direction) (target Point) {
	offset := DirectionOffsets[d]
	return Point{l.PlayerPos.Y + offset.Y, l.PlayerPos.X + offset.X}
}

func (l *Level) CheckCollisionBehind(d Direction) bool {
	offset := DirectionOffsets[d]
	target := Point{l.PlayerPos.Y + 2*offset.Y, l.PlayerPos.X + 2*offset.X}
	_, exists := l.Tiles[target]
	return exists
}

func (l *Level) IsValidInput(d Direction) bool {
	target := l.CalculateOffset(d)
	_, exists := l.Tiles[target]
	return exists
}

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
	}
	l.Tiles[target] = o
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
		if l.MovesLeft < 3 && l.CheckWin() {
			a = Win
		} else {
			a = Move
		}
		l.MovePlayerTo(d)
	case SpecialItem:
		a = SpecialItemCollect
		l.SpecialItemsCollected += 1
		l.MovePlayerTo(d)
	case Box:
	case BoxSpecialItem:
		if !l.CheckCollisionBehind(d) {
			a = PushBox
			l.PushCollidable(d, occupiedState)
		} else {
			a = PunchBox
			// Do nothing
		}
	case Skeleton:
		a = AttackSkeleton
		if !l.CheckCollisionBehind(d) {
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
