package render

import (
	"errors"
	"fmt"
	"snake-game/config"
	"snake-game/internal/game"
	"snake-game/pkg/terminal"
	"strings"
)

type TerminalRenderer struct {
	pulsarMode bool

	snakeHeadTexture string
	snakeBodyTexture []string
	snakeTailTexture string

	foodTexture   []string
	borderTexture []string

	screen *[][]string
}

func NewTerminalRenderer(config config.GameConfig) *TerminalRenderer {

	snakeHeadTexture, _ := randomizeTexture(config.Textures.SnakeHead)
	snakeTailTexture, _ := randomizeTexture(config.Textures.SnakeTail)

	var (
		snakeBodyTexture []string
		borderTexture    []string
		foodTexture      []string
	)
	for i := 0; i < (config.Width*config.Heigth + 2); i++ {
		nextSnakeCell, _ := randomizeTexture(config.Textures.SnakeBody)
		snakeBodyTexture = append(snakeBodyTexture, nextSnakeCell)

		nextBorder, _ := randomizeTexture(config.Textures.Border)
		borderTexture = append(borderTexture, nextBorder)

		nextFood, _ := randomizeTexture(config.Textures.Food)
		foodTexture = append(foodTexture, nextFood)
	}

	return &TerminalRenderer{
		pulsarMode: config.Render.PulsarMode,

		snakeHeadTexture: snakeHeadTexture,
		snakeBodyTexture: snakeBodyTexture,
		snakeTailTexture: snakeTailTexture,

		foodTexture:   foodTexture,
		borderTexture: borderTexture,

		screen: nil,
	}
}

func (tr TerminalRenderer) RenderGame(state game.GameState) error {
	terminal.HideCursor()

	// +2 для границ, ещё +1 в высоту для счёта
	w, h := state.Config.Width+2, state.Config.Heigth+3

	screen := make([][]string, h)
	for i := range screen {
		screen[i] = make([]string, w)
		for j := range screen[i] {
			screen[i][j] = state.Config.Textures.Background
		}
	}

	tr.screen = &screen

	if !tr.drawBorders(state) {
		return errors.New("Rendering borders error")
	}

	if !tr.drawScore(state) {
		return errors.New("Rendering score error")
	}

	if state.GameOver {
		if !tr.drawGameOverScreen(state) {
			return errors.New("GameOver screen rendering error")
		}
		// В книге "Go: идиомы и паттерны проектирования" указывается,
		// насколько я понял, что goto можно использовать в таких случаях
		// (не вызывает проблем и делает код более явным, чем альтернативы)
		goto PRINT
	}

	if !tr.drawFood(state) {
		return errors.New("Rendering food error")
	}

	if !tr.drawSnake(state) {
		return errors.New("Rendering snake error")
	}

PRINT:
	tr.printScreen()
	return nil
}

func (tr *TerminalRenderer) printScreen() {
	buff := ""
	for _, symbols := range *tr.screen {
		buff = strings.Join(symbols, "") + "\r\n" + buff
	}

	terminal.ClearScreen()
	fmt.Print(buff)
}
