package models

import (
	"math"

	"github.com/LeviyLokotb/light-automata/internal/models/mycolors"
	"github.com/LeviyLokotb/light-automata/pkg/materials"
)

type LightCell struct {
	materials.Material
	chans []*ColorChan
}

func NewLightCell(material materials.Material) LightCell {
	r := NewColorChan(mycolors.RED)
	g := NewColorChan(mycolors.GREEN)
	b := NewColorChan(mycolors.BLUE)
	return LightCell{
		Material: material,
		chans:    []*ColorChan{&r, &g, &b},
	}
}

func (l *LightCell) UpdateChanPhysics(ch int, laplacian float64, dt float64) {
	if l == nil {
		return
	}

	l.chans[ch].UpdatePhysics(laplacian, dt, l.Mass)
}

func (l *LightCell) UpdateChanHeight(ch int, dt, exposureRate float64) {
	if l == nil {
		return
	}

	l.chans[ch].UpdateHeight(dt, exposureRate)
}

func (l LightCell) GetChanHeight(ch int) float64 {
	return l.chans[ch].height
}

func (l *LightCell) SetHeight(h float64) {
	for i := range l.chans {
		l.chans[i].SetHeight(h)
	}
}

func (l LightCell) GetColorByAccumulated() [3]byte {
	color := [3]byte{}
	for i, ch := range l.chans {
		cv := math.Min(1.0, ch.GetAccumulatedLight())
		cv = cv * cv * 255
		if l.Mass < 1 {
			cv = math.Min(cv+float64(l.Color[i]), 255)
		}
		color[i] = byte(cv)
	}
	return color
}

func (l LightCell) GetColorByHeight() [3]byte {
	color := [3]byte{}
	for i, ch := range l.chans {
		cv := 0
		if ch.height > 0 {
			cv = int(ch.height * 100)
		}
		color[i] = byte(cv)
	}
	return color
}
