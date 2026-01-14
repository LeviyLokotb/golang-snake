package models

import "snake-game/internal/input"

type Snake struct {
	Body      []Point
	NextPoint Point
	Lenth     int
}

func (snake *Snake) SwitchDirection(direction input.Direction) {
	if direction == input.NONE {
		return
	}
	head := snake.Body[0]
	if snake.Lenth >= 2 {
		newPoint := head.ByDirectionPoint(direction)

		if newPoint.IsEqual(snake.Body[1]) {
			return
		}

		snake.NextPoint = newPoint
	}
}

func (snake Snake) IsUroboros() bool {
	head := snake.Body[0]
	for _, p := range snake.Body[1:] {
		if head.IsEqual(p) {
			return true
		}
	}
	return false
}
