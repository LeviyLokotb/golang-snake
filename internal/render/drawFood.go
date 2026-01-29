package render

import (
	"snake-game/internal/game"
	"snake-game/pkg/terminal"
)

func (tr *TerminalRenderer) drawFood(state game.GameState) bool {
	food := state.Food
	x, y := food.Position.X, food.Position.Y

	currentFoodTexture := tr.foodTexture[state.Score]

	texture := tr.getTexture(state.Config.Textures.Food, currentFoodTexture)

	texture = terminal.WrapTextWithStyle(texture, state.Config.Colors.Food)

	return tr.drawPixel(x, y, texture)
}
