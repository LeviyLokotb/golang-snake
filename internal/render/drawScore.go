package render

import (
	"fmt"
	"snake-game/internal/game"
)

func drawScore(screen *[][]string, state game.GameState) bool {
	score := fmt.Sprint(state.Score)

	foodTexture, ok := randomizeTexture(state.Config.Textures.Food)
	if !ok {
		return false
	}

	text := "Score: " + score + " " + foodTexture

	return drawText(screen, text, -1, -2)
}
