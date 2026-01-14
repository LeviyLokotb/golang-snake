package render

import (
	"math/rand"
)

func drawPixel(screen *[][]string, x, y int, texture string) bool {

	if x < -1 || y < -1 ||
		y+1 >= len((*screen)) ||
		x+1 >= len((*screen)[y+1]) {
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
	// debug_tools.AddToLog(fmt.Sprint(WStart, " | ", H))
	texture := ""
	CurrCellInc := 0
	for _, let := range message {
		// debug_tools.AddToLog(fmt.Sprint(WStart+CurrCellInc, " | ", len((*screen)[H+1])))
		if let == '\n' {
			if let != '\n' {
				texture += string(let)
			}
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

		if WStart+CurrCellInc+1 >= len((*screen)[H+1]) {
			continue
		}

		//debug_tools.AddToLog(H)
		cellLen := len((*screen)[H+1][WStart+CurrCellInc+1])

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

	if texture != "" {
		if !drawPixel(screen, WStart+CurrCellInc, H, texture) {
			return false
		}
	}
	return true
}
