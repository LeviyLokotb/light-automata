package models

import (
	"math"

	"github.com/LeviyLokotb/light-automata/internal/models/mycolors"
)

type ColorChan struct {
	height           float64
	velocity         float64
	accumulatedLight float64
	colorShift       mycolors.ColorShift
}

func NewColorChan(shift mycolors.ColorShift) ColorChan {
	return ColorChan{
		height:           0,
		velocity:         0,
		accumulatedLight: 0,
		colorShift:       shift,
	}
}

func (c *ColorChan) UpdatePhysics(laplacian float64, dt float64, mass float64) {
	if c == nil || mass < 0 {
		return
	}
	// ∂²u/∂t² = c²·∇²u
	vel := mass - float64(c.colorShift)
	vel = math.Max(vel, 0)

	acc := (laplacian/4 - c.height) * vel
	c.velocity += acc * dt
}

func (c *ColorChan) UpdateHeight(dt, exposureRate float64) {
	if c == nil {
		return
	}
	c.height += c.velocity * dt

	c.accumulatedLight += math.Abs(c.height) * exposureRate
	c.accumulatedLight *= 0.998
}

func (c *ColorChan) SetHeight(h float64) {
	if c == nil {
		return
	}
	c.height = h
}

func (c *ColorChan) GetAccumulatedLight() float64 {
	if c == nil {
		return 0
	}
	return c.accumulatedLight
}
