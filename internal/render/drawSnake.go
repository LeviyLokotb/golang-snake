package render

import (
	"snake-game/internal/game"
	"snake-game/pkg/terminal"
)

func (tr *TerminalRenderer) drawSnake(state game.GameState) bool {
	snake := state.Snake

	OK := true

	for i, point := range snake.Body {
		x, y := point.X, point.Y

		var texture string
		switch i {
		case 0:
			texture = tr.getTexture(state.Config.Textures.SnakeHead, tr.snakeHeadTexture)
			texture = terminal.WrapTextWithStyle(texture, state.Config.Colors.SnakeHead)
		case snake.Lenth - 1:
			texture = tr.getTexture(state.Config.Textures.SnakeTail, tr.snakeTailTexture)
			texture = terminal.WrapTextWithStyle(texture, state.Config.Colors.SnakeTail)
		default:
			texture = tr.getTexture(state.Config.Textures.SnakeBody, tr.snakeBodyTexture[i])
			texture = terminal.WrapTextWithStyle(texture, state.Config.Colors.SnakeBody)
		}

		OK = OK && tr.drawPixel(x, y, texture)
	}

	return OK
}
