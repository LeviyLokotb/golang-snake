package game

import (
	"snake-game/config"
	"snake-game/internal/input"
	"time"
)

type Renderer interface {
	RenderGame(GameState) error
}

type InputHandler interface {
	Start()
	Direction() <-chan input.Direction
	Pause() <-chan bool
	Restart() <-chan bool
	Quit() <-chan bool
	Stop()
}

type Game struct {
	state          *GameState
	renderer       Renderer
	inputHandler   InputHandler
	gamePaused     bool
	gameStop       bool
	StopGameSignal chan bool
}

func (game *Game) StopGame() {
	game.gameStop = true
	close(game.StopGameSignal)
}

func (game *Game) GameStop() <-chan bool {
	return game.StopGameSignal
}

func StartGame(config config.GameConfig, renderer Renderer, inputHandler *input.InputHandler) (*Game, error) {
	state, err := newGameState(config)
	if err != nil {
		return nil, err
	}

	game := Game{
		state:          &state,
		gamePaused:     false,
		gameStop:       false,
		renderer:       renderer,
		inputHandler:   inputHandler,
		StopGameSignal: make(chan bool),
	}

	game.inputHandler.Start()

	delay := time.Duration(config.Delay * int(time.Millisecond))

	go func() {
	MainCycle:
		for !game.gameStop {
			select {
			//case <-time.After(delay):

			case <-game.inputHandler.Quit():
				break MainCycle

			case <-game.inputHandler.Pause():
				game.gamePaused = !game.gamePaused

			case <-game.inputHandler.Restart():
				new_state, err := newGameState(config)
				if err != nil {
					break MainCycle
				}
				game.state = &new_state

			case dir := <-game.inputHandler.Direction():
				game.state.Snake.SwitchDirection(dir)

				if !game.gamePaused {
					err := updateGameState(game.state)
					if err != nil {
						break MainCycle
					}
				}
				time.Sleep(delay)

			default:
				if !game.gamePaused {
					err := updateGameState(game.state)
					if err != nil {
						break MainCycle
					}
				}
				time.Sleep(delay)
			}
		}
		game.StopGame()
	}()

	go func() {
		for !game.gameStop {
			err = game.renderer.RenderGame(*game.state)
			if err != nil {
				return
			}
			render_delay := delay / 10

			if render_delay <= 0 {
				render_delay = 1 * time.Millisecond
			}

			time.Sleep(render_delay)
		}
	}()

	return &game, nil
}
