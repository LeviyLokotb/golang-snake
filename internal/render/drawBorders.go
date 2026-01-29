package render

import (
	"snake-game/internal/game"
	"snake-game/internal/models"
	"snake-game/pkg/terminal"
)

func (tr *TerminalRenderer) drawBorders(state game.GameState) bool {
	W, H := state.Config.Width, state.Config.Heigth

	var plan []models.Point

	for i := -1; i <= W; i++ {
		plan = append(plan, models.Point{X: i, Y: -1})
		plan = append(plan, models.Point{X: i, Y: H})
	}

	for j := -1; j <= H; j++ {
		plan = append(plan, models.Point{X: -1, Y: j})
		plan = append(plan, models.Point{X: W, Y: j})
	}

	OK := true
	for i, point := range plan {
		x, y := point.X, point.Y

		texture := tr.getTexture(state.Config.Textures.Border, tr.borderTexture[i])

		texture = terminal.WrapTextWithStyle(texture, state.Config.Colors.Border)

		OK = OK && tr.drawPixel(x, y, texture)
	}

	return OK
}
