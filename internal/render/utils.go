package render

import (
	"math/rand"
)

func (tr *TerminalRenderer) drawPixel(x, y int, texture string) bool {

	if x < -1 || y < -2 ||
		y+2 >= len((*tr.screen)) ||
		x+1 >= len((*tr.screen)[y+2]) {
		return false
	}

	(*tr.screen)[y+2][x+1] = texture
	return true
}

func randomizeTexture(texture string) (string, bool) {
	symbols := []rune(texture)
	n := len(symbols)

	if n < 1 {
		return "?", false
	}

	if n > 1 {
		idx := rand.Intn(n)
		symb := symbols[idx]
		texture = string(symb)
	}

	return texture, true
}

func (tr *TerminalRenderer) getTexture(configTexture, defaultTexture string) string {
	if tr.pulsarMode {
		textture, ok := randomizeTexture(configTexture)
		if !ok {
			return "?"
		}
		return textture
	}
	return defaultTexture
}

func (tr *TerminalRenderer) drawText(message string, WStart, H int) bool {
	texture := ""
	CurrCellInc := 0
	for _, let := range message {
		if let == '\n' {
			if let != '\n' {
				texture += string(let)
			}
			if texture != "" {
				if !tr.drawPixel(WStart+CurrCellInc, H, texture) {
					return false
				}
			}
			texture = ""
			CurrCellInc = 0
			H -= 1
			continue
		}

		if WStart+CurrCellInc+1 >= len((*tr.screen)[H+2]) {
			continue
		}

		//debug_tools.AddToLog(H)
		cellLen := len((*tr.screen)[H+2][WStart+CurrCellInc+1])

		texture += string(let)
		if len(texture) < cellLen {
			continue
		}

		if !tr.drawPixel(WStart+CurrCellInc, H, texture) {
			return false
		}
		CurrCellInc++
		texture = ""
	}

	if texture != "" {
		if !tr.drawPixel(WStart+CurrCellInc, H, texture) {
			return false
		}
	}
	return true
}
