package main

import (
	"flag"
	"fmt"
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
	config_name := flag.String("config", "snake_config.json", "Game config file (json)")
	flag.Parse()

	// Загрузка захардкоженного конфига по умолчанию
	// config := config.NewDefaultConfig()

	// Загрузка конфига из json
	config, err := config.LoadConfigFromJSON(*config_name)
	if err != nil {
		fmt.Println("Error while loading config")
		panic(err)
	}

	// Объект, отвечающий за контроль терминала
	term := terminal.CurrentTerminal()
	// Объект, отвечающий за рендер графики
	renderer := render.NewTerminalRenderer(term)
	// Объект, отвечающий за отслеживание ввода
	ih := input.NewInputHandler(term)

	// Запуск
	game, err := game.StartGame(*config, renderer, ih)
	if err != nil {
		panic(err)
	}
	// Ожидание завершения
	<-game.GameStop()
}
