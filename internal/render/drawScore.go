package render

import (
	"fmt"
	"snake-game/internal/game"
)

func (tr *TerminalRenderer) drawScore(state game.GameState) bool {
	score := fmt.Sprint(state.Score)

	foodTexture := tr.getTexture(state.Config.Textures.Food, tr.foodTexture[0])

	text := "Score: " + score + " " + foodTexture

	return tr.drawText(text, -1, -2)
}
