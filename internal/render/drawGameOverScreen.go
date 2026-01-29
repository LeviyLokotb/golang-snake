package render

import (
	"snake-game/internal/game"
)

func (tr *TerminalRenderer) drawGameOverScreen(state game.GameState) bool {
	message := "Game Over!"

	scoreToWin := state.Config.Width*state.Config.Heigth - state.Config.InitSnakeLength

	if state.Score >= scoreToWin {
		message = "You Win!"
	}

	return tr.drawCenterMessage(message, state)
}

func (tr *TerminalRenderer) drawCenterMessage(message string, state game.GameState) bool {
	W := state.Config.Width

	backgroundLenth := len(state.Config.Textures.Background)

	halfMessageLen := len(message) / (backgroundLenth * 2)
	centerWidth := W / 2

	indexStart := centerWidth - halfMessageLen
	indexStart = max(0, indexStart)

	H := state.Config.Heigth
	centerHeigth := H / 2

	return tr.drawText(message, indexStart, centerHeigth)
}
