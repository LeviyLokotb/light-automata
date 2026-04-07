package models

import (
	"math"
	"runtime"
	"sync"

	"github.com/LeviyLokotb/light-automata/internal/config"
	"github.com/LeviyLokotb/light-automata/pkg/materials"
)

type WaveGrid struct {
	config       config.Config
	om           *materials.ObjectsManager
	grid         []*LightCell
	lightSources []*LightCell
	tick         int
}

func (g *WaveGrid) idx(x, y int) int {
	return y*g.config.WidthCells + x
}

func NewWaveGrid(config config.Config, om *materials.ObjectsManager) *WaveGrid {
	grid := make([]*LightCell, config.WidthCells*config.HeightCells)
	for i := 0; i < config.WidthCells; i++ {
		//grid[i] = make([]*LightCell, config.HeightCells)
		for j := 0; j < config.HeightCells; j++ {
			material := om.GetMaterialAt(i, j)
			cell := NewLightCell(material)
			grid[j*config.WidthCells+i] = &cell
		}
	}
	return &WaveGrid{
		config: config,
		om:     om,
		grid:   grid,
		tick:   0,
	}
}

func (g *WaveGrid) foreach(fn func(int, int)) {
	for x := 1; x < g.config.WidthCells-1; x++ {
		for y := 1; y < g.config.HeightCells-1; y++ {
			fn(x, y)
		}
	}
}

func (g *WaveGrid) foreachParallel(fn func(int, int)) {
	numWorkers := runtime.NumCPU()
	rowsPerWorker := (g.config.WidthCells - 2) / numWorkers

	var wg sync.WaitGroup
	for w := 0; w < numWorkers; w++ {
		startX := 1 + w*rowsPerWorker
		endX := startX + rowsPerWorker
		if w == numWorkers-1 {
			endX = g.config.WidthCells - 1
		}

		wg.Add(1)
		go func(startX, endX int) {
			defer wg.Done()
			for x := startX; x < endX; x++ {
				for y := 1; y < g.config.HeightCells-1; y++ {
					fn(x, y)
				}
			}
		}(startX, endX)
	}
	wg.Wait()
}

func (g *WaveGrid) foreachWithCh(fn func(int, int, int)) {
	g.foreach(func(x, y int) {
		for ch := 0; ch < 3; ch++ {
			fn(x, y, ch)
		}
	})
}

func (g *WaveGrid) foreachWithChParallel(fn func(int, int, int)) {
	g.foreachParallel(func(x, y int) {
		for ch := 0; ch < 3; ch++ {
			fn(x, y, ch)
		}
	})
}

func (g *WaveGrid) Update(dt float64) {
	g.tick++

	// Обновляем высоту (независимо)
	g.foreachWithChParallel(func(x, y, ch int) {
		cell := g.grid[g.idx(x, y)]

		cell.UpdateChanHeight(ch, dt, g.config.ExposureRate)

		if cell.IsGlow() {
			g.addWaveAt(x, y, g.tick)
		}
	})

	// Обновляем физические параметры
	w := g.config.WidthCells
	g.foreachWithChParallel(func(x, y, ch int) {
		idx := g.idx(x, y)
		laplacian := 0 +
			g.grid[idx+1].GetChanHeight(ch) +
			g.grid[idx-1].GetChanHeight(ch) +
			g.grid[idx+w].GetChanHeight(ch) +
			g.grid[idx-w].GetChanHeight(ch)

		g.grid[g.idx(x, y)].UpdateChanPhysics(ch, laplacian, dt)
	})
}

// func (g *WaveGrid) AddWaveLine(x, ymin, ymax int, tick int) {
// 	for y := ymin; y <= ymax; y++ {
// 		g.grid[g.idx(x, y)].SetHeight(math.Sin(float64(tick)*0.8) * 12)
// 	}
// }

func (g *WaveGrid) addWaveAt(x, y, tick int) {
	g.grid[g.idx(x, y)].SetHeight(math.Sin(float64(tick)*0.8) * 12)
}

func (g *WaveGrid) GetColorByAccumulated(x, y int) [3]byte {
	if x < 0 || x >= g.config.WidthCells || y < 0 || y >= g.config.HeightCells {
		return [3]byte{0, 0, 0}
	}
	return g.grid[g.idx(x, y)].GetColorByAccumulated()
}

func (g *WaveGrid) GetColorByHeight(x, y int) [3]byte {
	if x < 0 || x >= g.config.WidthCells || y < 0 || y >= g.config.HeightCells {
		return [3]byte{0, 0, 0}
	}
	return g.grid[g.idx(x, y)].GetColorByHeight()
}
