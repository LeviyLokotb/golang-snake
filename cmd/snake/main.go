package main

import (
	"flag"
	"snake-game/config"
	"snake-game/internal/game"
	"snake-game/internal/input"
	"snake-game/internal/render"
	debug_tools "snake-game/pkg/debug"
	"snake-game/pkg/terminal"
)

func main() {
	// Очищаем логи
	debug_tools.ClearLog()

	// Флаги
	homeDir, _ := debug_tools.GetHomeDir()
	config_name := flag.String("config", homeDir+"/"+".config/snake-go/config", "Game config file (json)")
	flag.Parse()

	// Загрузка захардкоженного конфига по умолчанию
	// config := config.NewDefaultConfig()

	// Загрузка конфига из json
	conf, err := config.LoadConfigFromJSON(*config_name)
	if err != nil {
		debug_tools.AddToLog("Error while loading config")
		conf = config.NewDefaultConfig()
	}

	// Объект, отвечающий за контроль терминала
	term := terminal.CurrentTerminal()
	// Объект, отвечающий за рендер графики
	renderer := render.NewTerminalRenderer(term)
	// Объект, отвечающий за отслеживание ввода
	ih := input.NewInputHandler(term)

	// Запуск
	game, err := game.StartGame(*conf, renderer, ih)
	if err != nil {
		panic(err)
	}
	// Ожидание завершения
	<-game.GameStop()
}
