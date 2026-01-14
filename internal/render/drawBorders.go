package render

import (
	"snake-game/internal/game"
	"snake-game/internal/models"
	"snake-game/pkg/terminal"
)

func drawBorders(screen *[][]string, state game.GameState) bool {
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
	for _, point := range plan {
		x, y := point.X, point.Y

		texture, ok := randomizeTexture(state.Config.Textures.Border)
		if !ok {
			OK = false
		}

		texture = terminal.WrapTextWithStyle(texture, state.Config.Colors.Border)

		OK = OK && drawPixel(screen, x, y, texture)
	}

	return OK
}
