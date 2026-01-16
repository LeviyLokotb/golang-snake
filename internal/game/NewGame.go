package game

import (
	"errors"
	"snake-game/config"
	"snake-game/internal/models"
)

func NewGame(config config.GameConfig) (GameState, error) {
	snake, errSnake := spawnSnake(config)
	if errSnake != nil {
		return GameState{}, errSnake
	}

	food, errFood := SpawnFood(config, snake.Body)
	if errFood != nil {
		return GameState{}, errFood
	}

	state := GameState{
		Snake:    &snake,
		Food:     &food,
		Score:    0,
		GameOver: false,
		Config:   config,
	}
	return state, nil
}

func spawnSnake(config config.GameConfig) (models.Snake, error) {
	W, H := config.Width, config.Heigth
	LENTH := config.InitSnakeLength

	if LENTH >= W*H {
		return models.Snake{}, errors.New("Snake so long")
	}

	headPoint := NewRandomPoint(2, W-2, 2, H-2)

	count := 0
	var nextBodyPoint models.Point
	for {
		if count >= 4 {
			return models.Snake{}, errors.New("Cannot create snake")
		}
		nextBodyPoint = headPoint.ByIntDirectionPoint(count)
		if nextBodyPoint.IsCorrect(0, W, 0, H) {
			break
		}
		count++
	}

	body := []models.Point{headPoint}

	for i := 0; i < LENTH-1; i++ {
		body = append(body, nextBodyPoint)
	}

	nextPoint, err := body[0].OppositePoint(body[1])
	if err != nil {
		return models.Snake{}, err
	}

	snake := models.Snake{
		Lenth:     LENTH,
		Body:      body,
		NextPoint: nextPoint,
	}

	return snake, nil
}

func SpawnFood(config config.GameConfig, snakeBody []models.Point) (models.Food, error) {
	W, H := config.Width, config.Heigth

	var food models.Food

FoodSpawning:
	for count := 0; ; count++ {
		if count > 2*W*H {
			return models.Food{}, errors.New("Cannot create food")
		}

		food = models.Food{
			Position: NewRandomPoint(0, W, 0, H),
		}

		for _, p := range snakeBody {
			if food.Position.Equal(p) {
				continue FoodSpawning
			}
		}

		break
	}

	return food, nil
}
