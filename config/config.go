package config

type GameConfig struct {
	Width           int `json:"width"`
	Heigth          int `json:"heigth"`
	Delay           int `json:"delay"` // in milliseconds
	InitSnakeLength int `json:"initSnakeLenth"`
	Textures        struct {
		SnakeHead  string `json:"snakeHead"`
		SnakeBody  string `json:"snakeBody"`
		SnakeTail  string `json:"snakeTail"`
		Food       string `json:"food"`
		Border     string `json:"border"`
		Background string `json:"background"`
	} `json:"textures"`
	Colors struct {
		SnakeHead string `json:"snakeHead"`
		SnakeBody string `json:"snakeBody"`
		SnakeTail string `json:"snakeTail"`
		Food      string `json:"food"`
		Border    string `json:"border"`
	} `json:"colors"`
}
