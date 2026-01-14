package terminal

import (
	"errors"
	"strconv"
	"strings"
)

func HexToByte(hex string) (r, g, b byte, _ error) {
	if !strings.HasPrefix(hex, "#") || len(hex) != 7 {
		return 255, 255, 255, errors.New("Invalid format")
	}

	r64, err := strconv.ParseUint(hex[1:3], 16, 8)
	if err != nil {
		return 0, 0, 0, err
	}

	g64, err := strconv.ParseUint(hex[3:5], 16, 8)
	if err != nil {
		return 0, 0, 0, err
	}

	b64, err := strconv.ParseUint(hex[5:7], 16, 8)
	if err != nil {
		return 0, 0, 0, err
	}

	return byte(r64), byte(g64), byte(b64), nil
}
