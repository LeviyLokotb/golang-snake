package render

import (
	"snake-game/internal/game"
	"snake-game/pkg/terminal"
)

func drawFood(screen *[][]string, state game.GameState) bool {
	food := state.Food
	x, y := food.Position.X, food.Position.Y

	texture, ok := randomizeTexture(state.Config.Textures.Food)
	if !ok {
		return false
	}

	texture = terminal.WrapTextWithStyle(texture, state.Config.Colors.Food)

	return drawPixel(screen, x, y, texture)
}
