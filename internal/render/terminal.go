package render

import (
	"errors"
	"fmt"
	"snake-game/internal/game"
	"snake-game/pkg/terminal"
	"strings"
)

type TerminalRenderer struct {
	Term *terminal.Terminal
}

func NewTerminalRenderer(t *terminal.Terminal) *TerminalRenderer {
	return &TerminalRenderer{
		Term: t,
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

	if !drawBorders(&screen, state) {
		return errors.New("Rendering borders error")
	}

	if !drawScore(&screen, state) {
		return errors.New("Rendering score error")
	}

	if state.GameOver {
		if !drawGameOverScreen(&screen, state) {
			return errors.New("GameOver screen rendering error")
		}
		// В книге "Go: идиомы и паттерны проектирования" указывается,
		// насколько я понял, что goto можно использовать в таких случаях
		// (не вызывает проблем и делает код более явным, чем альтернативы)
		goto PRINT
	}

	if !drawFood(&screen, state) {
		return errors.New("Rendering food error")
	}

	if !drawSnake(&screen, state) {
		return errors.New("Rendering snake error")
	}

PRINT:
	// Переводим из raw режима для корректной отрисовки
	//tr.Term.ToNormalMode()

	printScreen(&screen)

	//tr.Term.ToRawMode()
	return nil
}

func printScreen(screen *[][]string) {
	buff := ""
	for _, symbols := range *screen {
		buff = strings.Join(symbols, "") + "\r\n" + buff
	}

	terminal.ClearScreen()
	fmt.Print(buff)
}
