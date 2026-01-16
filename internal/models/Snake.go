package models

import (
	"slices"
	"snake-game/internal/input"
)

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

		if newPoint.Equal(snake.Body[1]) {
			return
		}

		snake.NextPoint = newPoint
	}
}

func (snake *Snake) Move(isFoodEaten bool) {

	snake.Body = slices.Insert(snake.Body, 0, snake.NextPoint)
	snake.Lenth += 1

	if !isFoodEaten {
		snake.Lenth -= 1
		snake.Body = snake.Body[:len(snake.Body)-1]
	}

	direction, errDirection := snake.Body[1].GetDirectionIndex(snake.Body[0])
	if errDirection != nil {
		direction = 0
	}

	snake.NextPoint = snake.Body[0].ByIntDirectionPoint(direction)
}

func (snake Snake) IsUroboros() bool {
	head := snake.Body[0]
	for _, p := range snake.Body[1:] {
		if head.Equal(p) {
			return true
		}
	}
	return false
}
