package game

import (
	"snake-game/config"
	"snake-game/internal/models"
)

func updateGameState(state *GameState) error {

	if state.GameOver {
		return nil
	}

	snake := state.Snake
	food := state.Food
	score := state.Score
	config := state.Config

	isFoodEaten, gameOverBySnake := updateSnake(snake, *food, config)

	if isFoodEaten {
		score += 1

		errFood := spawnNewFood(food, *snake, config)
		if errFood != nil {
			return errFood
		}
	}

	state.Score = score
	state.GameOver = gameOverBySnake

	return nil
}

func isPointOutOfBorder(p models.Point, config config.GameConfig) bool {
	return !p.IsCorrect(0, config.Width, 0, config.Heigth)
}

func updateSnake(snake *models.Snake, food models.Food, config config.GameConfig) (isFoodEaten, gameOver bool) {
	if isPointOutOfBorder(snake.NextPoint, config) || snake.IsUroboros() {
		return false, true
	}

	isFoodEaten = snake.NextPoint.Equal(food.Position)

	snake.Move(isFoodEaten)

	return isFoodEaten, false
}

func spawnNewFood(food *models.Food, snake models.Snake, config config.GameConfig) error {
	newFood, err := SpawnFood(config, snake.Body)
	if err != nil {
		return err
	}

	*food = newFood
	return nil
}
