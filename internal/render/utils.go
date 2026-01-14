package render

import (
	"math/rand"
)

func drawPixel(screen *[][]string, x, y int, texture string) bool {

	if x < -1 || y < -1 ||
		y+1 > len((*screen)) ||
		x+1 > len((*screen)[0]) {
		return false
	}

	(*screen)[y+1][x+1] = texture
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

func drawText(screen *[][]string, message string, WStart, H int) bool {
	texture := ""
	CurrCellInc := 0
	for i, let := range message {
		if let == '\n' {
			if texture != "" {
				if !drawPixel(screen, WStart+CurrCellInc, H, texture) {
					return false
				}
			}
			texture = ""
			CurrCellInc = 0
			H -= 1
			continue
		}

		cellLen := len((*screen)[H+1][i+1])

		texture += string(let)
		if len(texture) < cellLen {
			continue
		}

		if !drawPixel(screen, WStart+CurrCellInc, H, texture) {
			return false
		}
		CurrCellInc++
		texture = ""
	}
	return true
}
