package game

import (
	"snake-game/config"
	"snake-game/internal/models"
)

func UpdateGame(state *GameState) error {

	if state.GameOver {
		return nil
	}

	oldSnake := state.Snake
	oldFood := state.Food
	score := state.Score
	config := state.Config

	newSnake, gameOverBySnake := nextSnake(oldSnake, oldFood, config)

	newFood := oldFood
	if newSnake.Lenth > oldSnake.Lenth {
		score += 1
		var errFood error
		newFood, errFood = spawnNewFood(oldFood, newSnake, config)
		if errFood != nil {
			return errFood
		}
	}

	*state = GameState{
		Snake:    newSnake,
		Food:     newFood,
		Score:    score,
		GameOver: gameOverBySnake,
		Config:   config,
	}
	return nil
}

func isPointOutOfBorder(p models.Point, config config.GameConfig) bool {
	return !p.IsCorrect(0, config.Width, 0, config.Heigth)
}

func nextSnake(oldSnake *models.Snake, oldFood *models.Food, config config.GameConfig) (newSnake *models.Snake, gameOver bool) {
	if isPointOutOfBorder(oldSnake.NextPoint, config) || oldSnake.IsUroboros() {
		return oldSnake, true
	}

	newBody := []models.Point{oldSnake.NextPoint}
	newBody = append(newBody, oldSnake.Body...)
	newLenth := oldSnake.Lenth + 1

	if !newBody[0].IsEqual(oldFood.Position) {
		newLenth -= 1
		newBody = newBody[:len(newBody)-1]
	}

	direction, errDirection := newBody[1].GetDirectionIndex(newBody[0])
	if errDirection != nil {
		direction = 0
	}

	newNextPoint := newBody[0].ByIntDirectionPoint(direction)

	newSnake = &models.Snake{
		Body:      newBody,
		NextPoint: newNextPoint,
		Lenth:     newLenth,
	}
	return newSnake, false
}

func spawnNewFood(oldFood *models.Food, snake *models.Snake, config config.GameConfig) (*models.Food, error) {
	food, err := SpawnFood(config, snake.Body)

	for food.Position.IsEqual(oldFood.Position) {
		food, err = SpawnFood(config, snake.Body)
		if err != nil {
			return nil, err
		}
	}
	return &food, err
}
