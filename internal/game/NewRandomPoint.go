package game

import (
	"math/rand"
	"snake-game/internal/models"
)

func NewRandomPoint(minW, W, minH, H int) models.Point {
	x := rand.Intn(W-minW) + minW
	y := rand.Intn(H-minH) + minH
	return models.Point{
		X: x,
		Y: y,
	}
}
