package terminal

import (
	"fmt"
	"strings"
)

func GetStyle(style string) string {
	result := ""
	styles := strings.Split(style, "+")
	for _, s := range styles {
		s = strings.ToUpper(s)
		var (
			code string
			ok   bool
		)

		if strings.HasPrefix(s, "#") && len(s) == 7 {
			r, g, b, err := HexToByte(s)
			if err != nil {
				ok = false
			}
			code = fmt.Sprintf("38;2;%d;%d;%d", r, g, b)
			ok = true
		} else {
			code, ok = colors[s]
		}

		if ok {
			result += "\033[" + code + "m"
		}
	}
	return result
}

func WrapTextWithStyle(text, style string) string {
	return GetStyle(style) + text + GetStyle("RESET")
}

var colors map[string]string = map[string]string{
	"RESET":      "0",
	"BOLD":       "1",
	"BLINK":      "5",
	"FAST_BLINK": "5",
	"NEGATIVE":   "7",
	"BACKGROUND": "49",
	"BLACK":      "30",
	"RED":        "31",
	"GREEN":      "32",
	"YELLOW":     "33",
	"BLUE":       "34",
	"MAGENTA":    "35",
	"CYAN":       "36",
	"WHITE":      "37",
}
