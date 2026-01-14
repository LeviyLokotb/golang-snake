package game

import (
	"snake-game/config"
	"snake-game/internal/models"
)

type GameState struct {
	Snake    *models.Snake
	Food     *models.Food
	Score    int
	GameOver bool
	Config   config.GameConfig
}
