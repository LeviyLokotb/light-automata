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

	// Создаём объекты сцены
	// sphere := materials.NewSphere(materials.GetGlass(), 80, 300, 300)
	//backWall := materials.NewRect(materials.GetWall(), 120, 125, 200, 350)
	wall0 := materials.NewRect(materials.GetWall(), 0, 49, 126, 175)
	wall1 := materials.NewRect(materials.GetWall(), 45, 65, 172, 175)
	wall2 := materials.NewRect(materials.GetWall(), 45, 65, 126, 129)

	leftWall := materials.NewRect(materials.GetWall(), 0, 0, 0, 400)
	rightWall := materials.NewRect(materials.GetWall(), 400, 400, 0, 400)
	topWall := materials.NewRect(materials.GetWall(), 0, 400, 0, 0)
	bottomWall := materials.NewRect(materials.GetWall(), 0, 400, 400, 400)

	bulb := materials.NewRect(materials.GetAir().MakeGlow(), 50, 50, 130, 170)
	prism := materials.NewSphere(materials.GetGlass(), 100, 200, 200)
	ushko1 := materials.NewTriangle(materials.GetGlass(), 200, 150, 100, 70, 100, 200)
	ushko2 := materials.NewTriangle(materials.GetGlass(), 200, 150, 300, 70, 300, 200)
	glazik1 := materials.NewSphere(materials.GetDiamond().MakeGlow(), 30, 150, 220)
	glazik2 := materials.NewSphere(materials.GetDiamond().MakeGlow(), 30, 250, 220)

	objects := []materials.Object{
		ushko1,
		ushko2,
		prism,
		bulb,
		wall0, wall1, wall2,
		leftWall,
		rightWall,
		topWall,
		bottomWall,
		glazik1,
		glazik2,
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
