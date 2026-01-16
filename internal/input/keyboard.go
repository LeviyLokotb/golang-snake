package input

import (
	"bufio"
	"os"
	"snake-game/pkg/terminal"
	"time"
	"unicode"
)

type InputHandler struct {
	direction chan Direction
	pause     chan bool
	restart   chan bool
	quit      chan bool
	Term      *terminal.Terminal
}

func NewInputHandler(t *terminal.Terminal) *InputHandler {
	return &InputHandler{
		direction: make(chan Direction, 1),
		pause:     make(chan bool, 1),
		restart:   make(chan bool, 1),
		quit:      make(chan bool),
		Term:      t,
	}
}

func (ih *InputHandler) Start() {
	go func() {
		// Переводим терминал в "сырой" режим чтобы считывать руны небуферизированно
		ih.Term.ToRawMode()
		defer ih.Term.ToNormalMode()

		reader := bufio.NewReader(os.Stdin)

		for {
			// select выбирает первый канал который отправит данные
			select {
			case <-time.After(10 * time.Millisecond):
			// Событие: выход
			case <-ih.quit:
				return
			// Остальные события
			default:
				char, _, err := reader.ReadRune()

				if err != nil {
					continue
				}

				// Обработка escape-последовательностей
				if char == '\033' {
					nextChar, _, err := reader.ReadRune()
					if nextChar != '[' || err != nil {
						continue
					}

					escapeChar, _, err := reader.ReadRune()
					if err != nil {
						continue
					}

					switch escapeChar {
					case 'A':
						char = 'w'
					case 'B':
						char = 's'
					case 'C':
						char = 'd'
					case 'D':
						char = 'a'
					}
				}

				switch unicode.ToLower(char) {
				case 'q', 'й', '`', 'ё':
					ih.Stop()
					return
				case 'r', 'к', '\r':
					select {
					case ih.restart <- true:
					default:
						<-ih.restart
						ih.restart <- true
					}
					continue
				case ' ', 'e', 'у', 'p', 'з':
					select {
					case ih.pause <- true:
					default:
						<-ih.pause
						ih.pause <- true
					}
				default:
					// Можно добавить другие сочетания
				}

				dir := GetDirectionFromRune(char)
				if dir != NONE {
					select {
					case ih.direction <- dir:
						//
					default:
						// Нужно очистить канал
						<-ih.direction
						ih.direction <- dir
					}
				}

				//time.Sleep(10 * time.Millisecond)
			}
		}
	}()
}

// Возвращает readonly канал (если бы мы возвращали значение,
// поток бы блокировался до чтения из канала)

func (ih *InputHandler) Direction() <-chan Direction {
	return ih.direction
}

func (ih *InputHandler) Pause() <-chan bool {
	return ih.pause
}

func (ih *InputHandler) Restart() <-chan bool {
	return ih.restart
}

func (ih *InputHandler) Quit() <-chan bool {
	return ih.quit
}

func (ih *InputHandler) Stop() {
	close(ih.quit)
}
