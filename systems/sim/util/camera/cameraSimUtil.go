package cameraSimUtil

import (
	"math"

	"github.com/kainn9/demo/components"
	cameraGlobals "github.com/kainn9/demo/globalConfig/camera"
)

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

	incX := cameraGlobals.CAMERA_MIN_SPEED
	incY := cameraGlobals.CAMERA_MIN_SPEED

	if math.Abs(c.X-x) > float64(cameraGlobals.CAMERA_EPSILON) {
		incX = cameraGlobals.CAMERA_MAX_SPEED
	}

	if math.Abs(c.Y-y) > float64(cameraGlobals.CAMERA_EPSILON) {
		incY = cameraGlobals.CAMERA_MAX_SPEED
	}

	c.X = smooth(c.X, x, incX)
	c.Y = smooth(c.Y, y, incY)
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
