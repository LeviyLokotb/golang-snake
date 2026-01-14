package render

import (
	"snake-game/internal/game"
)

func drawGameOverScreen(screen *[][]string, state game.GameState) bool {
	message := "Game Over!"

	scoreToWin := state.Config.Width*state.Config.Heigth - state.Config.InitSnakeLength

	if state.Score >= scoreToWin {
		message = "You Win!"
	}

	return drawCenterMessage(message, screen, state)
}

func drawCenterMessage(message string, screen *[][]string, state game.GameState) bool {
	W := state.Config.Width

	backgroundLenth := len(state.Config.Textures.Background)

	halfMessageLen := len(message) / (backgroundLenth * 2)
	centerWidth := W / 2

	indexStart := centerWidth - halfMessageLen
	indexStart = max(0, indexStart)

	H := state.Config.Heigth
	centerHeigth := H / 2

	return drawText(screen, message, indexStart, centerHeigth)
}
