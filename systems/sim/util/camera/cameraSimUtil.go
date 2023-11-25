package cameraSimUtil

import (
	"github.com/kainn9/demo/components"
)

func lerp(a, b, t float64) float64 {
	return a + t*(b-a)
}

func SetPositionLerp(c *components.Camera, targetX, targetY float64) {
	lerpFactor := 0.04
	c.X = lerp(c.X, targetX, lerpFactor)
	c.Y = lerp(c.Y, targetY, lerpFactor)
}

func SetPosition(c *components.Camera, targetX, targetY float64) {
	c.X = targetX
	c.Y = targetY
}
