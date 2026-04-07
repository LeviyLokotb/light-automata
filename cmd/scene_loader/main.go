package main

import (
	"fmt"
	"log"
	"os"

	"github.com/LeviyLokotb/light-automata/internal/config"
	"github.com/LeviyLokotb/light-automata/internal/render"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s /path/to/config.yaml\n", os.Args[0])
		os.Exit(1)
	}

	path := os.Args[1]
	// Загружаем параметры и сцену из файла
	conf, om, err := config.LoadSuperConfigFromYaml(path)
	if err != nil {
		log.Fatal(err)
	}

	game := render.NewGame(*conf, om, 10000)

	// Устанавливаем параметры движка
	ebiten.SetWindowSize(800, 800)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	// Запускаем
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
