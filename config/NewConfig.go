package config

type Modifier func(*GameConfig)

func NewConfig(modifiers ...Modifier) *GameConfig {
	conf := NewDefaultConfig()

	for _, modify := range modifiers {
		modify(conf)
	}

	return conf
}

func SetShape(width, heigth int) Modifier {
	return func(gc *GameConfig) {
		gc.Width = width
		gc.Heigth = heigth
	}
}

func SetSpeed(speed int) Modifier {
	return func(gc *GameConfig) {
		gc.Delay = speed
	}
}

func SetInitialSnakeLenth(lenth int) Modifier {
	return func(gc *GameConfig) {
		gc.InitSnakeLength = lenth
	}
}
