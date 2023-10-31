package cameraUtil

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kainn9/demo/components"
	cameraConstants "github.com/kainn9/demo/constants/camera"
)

func Translate(c *components.Camera, ops *ebiten.DrawImageOptions, x, y float64) *ebiten.DrawImageOptions {
	ops.GeoM.Translate(-c.X+x, -c.Y+y)
	return ops
}

func AddImage(c *components.Camera, img *ebiten.Image, options *ebiten.DrawImageOptions) {
	c.Surface.DrawImage(img, options)
}

func Render(c *components.Camera, screen *ebiten.Image) {
	screen.DrawImage(c.Surface, &ebiten.DrawImageOptions{})
}

func SetPosition(c *components.Camera, x, y float64, smoothCam bool) {

	if !smoothCam {
		c.X = x
		c.Y = y
		return
	}

	incX := cameraConstants.CAMERA_MIN_SPEED
	incY := cameraConstants.CAMERA_MIN_SPEED

	if math.Abs(c.X-x) > cameraConstants.CAMERA_EPSILON {
		incX = cameraConstants.CAMERA_MAX_SPEED
	}

	if math.Abs(c.Y-y) > cameraConstants.CAMERA_EPSILON {
		incY = cameraConstants.CAMERA_MAX_SPEED
	}

	c.X = smooth(c.X, x, incX)
	c.Y = smooth(c.Y, y, incY)
}

func Clear(c *components.Camera) {
	c.Surface.Clear()
}

func smooth(current, target, speed float64) float64 {
	// Calculate the change in position based on the speed.
	if current < target {
		current += speed
		if current > target {
			current = target
		}
	} else if current > target {
		current -= speed
		if current < target {
			current = target
		}
	}

	return current
}
