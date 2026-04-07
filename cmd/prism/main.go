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

	leftWall := materials.NewRect(materials.GetWall(), 0, 2, 0, 400)
	rightWall := materials.NewRect(materials.GetWall(), 398, 400, 0, 400)
	topWall := materials.NewRect(materials.GetWall(), 0, 400, 0, 2)
	bottomWall := materials.NewRect(materials.GetWall(), 0, 400, 398, 400)

	bulb := materials.NewRect(materials.GetAir().MakeGlow(), 50, 50, 130, 170)
	// Примерно правильный треугольник
	ax, ay, bx, by, cx, cy := 200.0, 80.0, 300.0, 255.0, 100.0, 255.0
	axr, ayr, bxr, byr, cxr, cyr := materials.RotateTriangle(ax, ay, bx, by, cx, cy, 0.5)
	prism := materials.NewTriangle(materials.GetGlass(), int(axr), int(ayr), int(bxr), int(byr), int(cxr), int(cyr))
	objects := []materials.Object{
		prism,
		bulb,
		wall0, wall1, wall2,
		leftWall,
		rightWall,
		topWall,
		bottomWall,
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
