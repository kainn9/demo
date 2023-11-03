package cameraUtil

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kainn9/demo/components"
	cameraConstants "github.com/kainn9/demo/constants/camera"
)

// Translates a image relative to the camera via its draw options(mutates), and x/y coordinates(args).
func Translate(c *components.Camera, ops *ebiten.DrawImageOptions, x, y float64) *ebiten.DrawImageOptions {
	ops.GeoM.Translate(-c.X+x, -c.Y+y)
	return ops
}

// Adds an image to the camera's surface. You probably want to use Translate() before this,
// and want to use the mutated/translated draw options.
func AddImage(c *components.Camera, img *ebiten.Image, options *ebiten.DrawImageOptions) {
	c.Surface.DrawImage(img, options)
}

// Renders the camera's surface to the screen.
// This is done in in the final render phase usually.
func Render(c *components.Camera, screen *ebiten.Image) {
	screen.DrawImage(c.Surface, &ebiten.DrawImageOptions{})
}

// Sets the camera's position to the given x/y coordinates.
// Has a smooth camera option, which will move the camera at a
// more constant speed that increases if the player gets too far.
// We prefer constant speeds over "lerping" because  lerping
// will slow down the camera as it approaches the target, which
// can make the parallax effect look weird when it gets really
// close.
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

// Clears the camera's surface.
func Clear(c *components.Camera) {
	c.Surface.Clear()
}

// Private "smoothing" function for the cam setPos.
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
