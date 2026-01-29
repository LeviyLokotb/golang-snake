package config

func NewDefaultConfig() *GameConfig {
	return &GameConfig{
		Width:           30,
		Heigth:          20,
		Delay:           400,
		InitSnakeLength: 3,
		Textures: struct {
			SnakeHead  string `json:"snakeHead"`
			SnakeBody  string `json:"snakeBody"`
			SnakeTail  string `json:"snakeTail"`
			Food       string `json:"food"`
			Border     string `json:"border"`
			Background string `json:"background"`
		}{
			SnakeHead:  "✬",
			SnakeBody:  "❈❄",
			SnakeTail:  ".",
			Food:       "✦✶",
			Border:     "▩▨▧",
			Background: " ",
		},
		Colors: struct {
			SnakeHead string `json:"snakeHead"`
			SnakeBody string `json:"snakeBody"`
			SnakeTail string `json:"snakeTail"`
			Food      string `json:"food"`
			Border    string `json:"border"`
		}{
			SnakeHead: "RED",
			SnakeBody: "GREEN",
			SnakeTail: "YELLOW",
			Food:      "CYAN",
			Border:    "BOLD+MAGENTA",
		},
		Render: struct {
			PulsarMode bool `json:"pulsarMode"`
		}{
			PulsarMode: false,
		},
	}
}
