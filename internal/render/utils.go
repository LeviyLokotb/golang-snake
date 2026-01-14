package render

import "math/rand"

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
