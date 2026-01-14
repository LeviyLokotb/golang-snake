package render

import "snake-game/internal/game"

func drawGameOverScreen(screen *[][]string, state game.GameState) bool {
	const message = "Game Over!"

	W := state.Config.Width
	if W < len(message) {
		return false
	}

	backgroundLenth := len(state.Config.Textures.Background)

	halfMessageLen := len(message) / (backgroundLenth * 2)
	centerWidth := W / 2

	indexStart := centerWidth - halfMessageLen

	H := state.Config.Heigth
	centerHeigth := H / 2

	return drawText(screen, message, indexStart, centerHeigth)
}
