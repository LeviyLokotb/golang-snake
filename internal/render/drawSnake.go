package render

import (
	"snake-game/internal/game"
	"snake-game/pkg/terminal"
)

func drawSnake(screen *[][]string, state game.GameState) bool {
	snake := state.Snake

	OK := true

	for i, point := range snake.Body {
		x, y := point.X, point.Y

		var (
			texture string
			ok      bool
		)
		switch i {
		case 0:
			texture, ok = randomizeTexture(state.Config.Textures.SnakeHead)
			texture = terminal.WrapTextWithStyle(texture, state.Config.Colors.SnakeHead)
		case snake.Lenth - 1:
			texture, ok = randomizeTexture(state.Config.Textures.SnakeTail)
			texture = terminal.WrapTextWithStyle(texture, state.Config.Colors.SnakeTail)
		default:
			texture, ok = randomizeTexture(state.Config.Textures.SnakeBody)
			texture = terminal.WrapTextWithStyle(texture, state.Config.Colors.SnakeBody)
		}

		if !ok {
			OK = false
		}

		OK = OK && drawPixel(screen, x, y, texture)
	}

	return OK
}
