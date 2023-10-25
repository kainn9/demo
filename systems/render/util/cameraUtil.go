package renderSystems

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kainn9/demo/components"
	"github.com/kainn9/demo/constants"
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

	incX := constants.CAMERA_MIN_SPEED
	incY := constants.CAMERA_MIN_SPEED

	if math.Abs(c.X-x) > constants.CAMERA_EPSILON {
		incX = constants.CAMERA_MAX_SPEED
	}

	if math.Abs(c.Y-y) > constants.CAMERA_EPSILON {
		incY = constants.CAMERA_MAX_SPEED
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
