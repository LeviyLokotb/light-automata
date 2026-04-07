package render

import (
	"image/color"

	"github.com/LeviyLokotb/light-automata/internal/config"
	"github.com/LeviyLokotb/light-automata/internal/models"
	"github.com/LeviyLokotb/light-automata/pkg/materials"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	grid      *models.WaveGrid
	width     int
	height    int
	pixelSize int
	delta     float64
	tick      int
	maxTicks  int
	waveMode  bool
}

func NewGame(conf config.Config, om *materials.ObjectsManager, maxTicks int) *Game {
	return &Game{
		grid:      models.NewWaveGrid(conf, om),
		width:     conf.WidthCells,
		height:    conf.HeightCells,
		pixelSize: conf.PixelSize,
		delta:     1,
		tick:      0,
		maxTicks:  maxTicks,
		waveMode:  conf.WaveMode,
	}
}

func (g *Game) Update() error {
	//log.Printf("Physics tick %d/%d", g.tick, g.maxTicks)

	if g.tick > g.maxTicks {
		return ebiten.Termination
	}

	// if g.tick < 300 {
	// 	g.grid.AddWaveSource(50, g.height*47/100, g.height*53/100, g.tick)
	// }

	g.grid.Update(g.delta)
	g.tick++

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	img := ebiten.NewImage(g.width, g.height)

	for x := 0; x < g.width; x++ {
		for y := 0; y < g.height; y++ {
			var rgb [3]byte
			if g.waveMode {
				rgb = g.grid.GetColorByHeight(x, y)
			} else {
				rgb = g.grid.GetColorByAccumulated(x, y)
			}
			img.Set(x, y, color.RGBA{
				R: rgb[0],
				G: rgb[1],
				B: rgb[2],
				A: 255,
			})
		}
	}

	// Масштабирование
	options := &ebiten.DrawImageOptions{}
	scaleX := float64(screen.Bounds().Dx()) / float64(g.width)
	scaleY := float64(screen.Bounds().Dy()) / float64(g.height)
	options.GeoM.Scale(scaleX, scaleY)

	// Отрисовка
	screen.DrawImage(img, options)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Возвращаем логический размер экрана
	return g.width * g.pixelSize, g.height * g.pixelSize
}
