// Core logic for recreating Helltaker game mechanics
package core

func IsHazard(s OccupiedState) bool {
	return s == Hazard ||
		s == SwitchHazardEven ||
		s == SwitchHazardOdd
}

func (l *Level) CalculateOffset(d Direction) (target Point) {
	offset := DirectionOffsets[d]
	return Point{l.PlayerPos.Y + offset.Y, l.PlayerPos.X + offset.X}
}

func (l *Level) CanPushTo(d Direction) bool {
	offset := DirectionOffsets[d]
	behind := Point{l.PlayerPos.Y + 2*offset.Y, l.PlayerPos.X + 2*offset.X}
	state, exists := l.Tiles[behind]
	return exists && (state == Empty || IsHazard(state) || state == Key)
}

func (l *Level) IsValidInput(d Direction) bool {
	target := l.CalculateOffset(d)
	_, exists := l.Tiles[target]
	return exists
}

// Returns true if the player is adjacent to the goal
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
	directions := []Direction{Up, Right, Down, Left}
	for _, d := range directions {
		if l.IsValidInput(d) {
			ds = append(ds, d)
		}
	}
	return ds
}

func (l *Level) CheckHazardActiveness(o OccupiedState) bool {
	isEven := l.MovesCount%2 == 0

	switch o {
	case Hazard:
		return true
	case SwitchHazardEven:
		return isEven
	case SwitchHazardOdd:
		return !isEven
	default:
		return false
	}
}

func (l *Level) MovePlayerTo(d Direction) {
	target := l.CalculateOffset(d)

	l.Tiles[l.PlayerPos] = l.UnderPlayer
	l.UnderPlayer = l.Tiles[target]

	l.Tiles[target] = Player
	l.PlayerPos = target
}

func (l *Level) PushCollidable(d Direction, o OccupiedState) {
	offset := DirectionOffsets[d]
	current := l.CalculateOffset(d)
	// Tile behind
	target := Point{current.Y + offset.Y, current.X + offset.X}
	targetState := l.Tiles[target]

	switch o {
	case Box:
		l.Tiles[current] = Empty
	case BoxHazard:
		l.Tiles[current] = Hazard
	case BoxSwitchHazardEven:
		l.Tiles[current] = SwitchHazardEven
	case BoxSwitchHazardOdd:
		l.Tiles[current] = SwitchHazardOdd
	case BoxSpecialItem:
		l.Tiles[current] = SpecialItem
	case BoxKey:
		l.Tiles[current] = Key
	case Skeleton:
		l.Tiles[current] = Empty
		if !IsHazard(targetState) {
			l.Tiles[target] = o
		}
		return
	default:
		panic("Unhandled occupied state in PushCollidable")
	}

	switch targetState {
	case Hazard:
		l.Tiles[target] = BoxHazard
	case SwitchHazardEven:
		l.Tiles[target] = BoxSwitchHazardEven
	case SwitchHazardOdd:
		l.Tiles[target] = BoxSwitchHazardOdd
	case Key:
		l.Tiles[target] = BoxKey
	default:
		l.Tiles[target] = Box
	}
}

func (l *Level) AttackSkeleton(t Point) {
	l.Tiles[t] = Empty
	if l.CheckHazardActiveness(l.UnderPlayer) {
		l.MovesLeft -= 1
	}
}

func (l *Level) HandleInput(d Direction) (a Action) {
	target := l.CalculateOffset(d)
	occupiedState := l.Tiles[target]

	switch occupiedState {
	case Empty:
		l.MovePlayerTo(d)

		a = Move

	case Hazard, SwitchHazardEven, SwitchHazardOdd:
		if l.CheckHazardActiveness(occupiedState) {
			a = TouchHazard
			l.MovesLeft -= 1
		} else {
			a = Move
		}
		l.MovePlayerTo(d)
	case SpecialItem:
		a = SpecialItemCollect
		l.SpecialItemsCollected += 1
		l.MovePlayerTo(d)
	case Key:
		a = CollectKey
		l.KeyCollected = true
		l.MovePlayerTo(d)
	case Chest:
		if l.KeyCollected {
			a = OpenChest
			l.Tiles[target] = Empty
			l.MovePlayerTo(d)
		} else {
			a = PunchChest
		}
	case Box,
		BoxSpecialItem,
		BoxHazard,
		BoxSwitchHazardEven,
		BoxSwitchHazardOdd,
		BoxKey:
		if l.CanPushTo(d) {
			a = PushBox
			l.PushCollidable(d, occupiedState)

			if l.CheckHazardActiveness(l.UnderPlayer) {
				a = TouchHazardPushBox
				l.MovesLeft -= 1
			}
		} else {
			// Punching a box does nothing in this game
			a = PunchBox

			if l.CheckHazardActiveness(l.UnderPlayer) {
				a = TouchHazardPunchBox
				l.MovesLeft -= 1
			}
		}
	case Skeleton:
		if l.CanPushTo(d) {
			a = PushSkeleton
			l.PushCollidable(d, occupiedState)

			if l.CheckHazardActiveness(l.UnderPlayer) {
				a = TouchHazardPushSkeleton
				l.MovesLeft -= 1
			}
		} else {
			a = AttackSkeleton
			l.AttackSkeleton(target)

			if l.CheckHazardActiveness(l.UnderPlayer) {
				a = TouchHazardAttackSkeleton
				l.MovesLeft -= 1
			}
		}
	}

	if l.CheckWin() {
		a = Win
	}

	l.MovesLeft -= 1
	l.MovesCount += 1
	return a
}
