package main

import (
	"log"

	"github.com/LeviyLokotb/light-automata/internal/config"
	"github.com/LeviyLokotb/light-automata/internal/render"
	"github.com/LeviyLokotb/light-automata/pkg/materials"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	// Параметры
	conf := config.NewDefault()
	conf.WaveMode = true

	// Создаём объекты сцены
	bulb := materials.NewRect(materials.GetAir().MakeGlow(), 0, 1, 0, 400)
	wall := materials.NewRect(materials.GetWall(), 30, 34, 0, 400)
	gap1 := materials.NewRect(materials.GetAir(), 30, 34, 120, 140)
	gap2 := materials.NewRect(materials.GetAir(), 30, 34, 260, 280)

	leftWall := materials.NewRect(materials.GetWall(), 0, 0, 0, 400)
	rightWall := materials.NewRect(materials.GetWall(), 400, 400, 0, 400)
	topWall := materials.NewRect(materials.GetWall(), 0, 400, 0, 0)
	bottomWall := materials.NewRect(materials.GetWall(), 0, 400, 400, 400)

	objects := []materials.Object{
		leftWall,
		rightWall,
		topWall,
		bottomWall,
		bulb,
		wall,
		gap1,
		gap2,
	}

	om := materials.NewObjectsManager(objects)

	game := render.NewGame(conf, om, 10000)

	// Устанавливаем параметры движка
	ebiten.SetWindowSize(800, 800)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	// Запускаем
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
