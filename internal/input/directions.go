package input

import (
	"unicode"
)

type Direction int

const (
	NONE Direction = iota
	UP
	LEFT
	DOWN
	RIGHT
)

func GetDirectionFromRune(char rune) Direction {
	switch unicode.ToLower(char) {
	case 'w', 'ц', 'k', 'л':
		return UP
	case 'a', 'ф', 'h', 'р':
		return LEFT
	case 's', 'ы', 'j', 'о':
		return DOWN
	case 'd', 'в', 'l', 'д':
		return RIGHT
	}
	return NONE
}
