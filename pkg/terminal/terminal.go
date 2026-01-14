package terminal

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

type Terminal struct {
	fd       int
	oldState *term.State
	isRaw    bool
}

func CurrentTerminal() *Terminal {
	return &Terminal{
		fd:    int(os.Stdin.Fd()),
		isRaw: false,
	}
}

func (t *Terminal) ToRawMode() {
	oldState, err := term.MakeRaw(t.fd)
	if err != nil {
		return
	}
	t.oldState = oldState
	t.isRaw = true
}

func (t *Terminal) ToNormalMode() {
	term.Restore(t.fd, t.oldState)
	t.isRaw = false
}

func ClearScreen() {
	// CSI H -- переместить курсор в начало
	// CSI 2J -- очистить окно терминала
	fmt.Print("\033[H\033[2J")
}

func HideCursor() {
	fmt.Print("\033[?25l")
}

func ShowCursor() {
	fmt.Print("\033[?25h")
}
