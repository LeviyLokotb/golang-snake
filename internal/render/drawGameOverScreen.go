package render

import "snake-game/internal/game"

func drawGameOverScreen(screen *[][]string, state game.GameState) bool {
	const message = "Game Over!"

	W := state.Config.Width
	if W < len(message) {
		return false
	}

	halfMessageLen := len(message) / 2
	centerWidth := W / 2

	indexStart := centerWidth - halfMessageLen

	H := state.Config.Heigth
	centerHeigth := H / 2

	for i := range message {
		let := string(message[i])
		for len(let) < len(state.Config.Textures.Background) {
			let += " "
		}
		if !drawPixel(screen, indexStart+i, centerHeigth, let) {
			return false
		}
	}

	return true
}
