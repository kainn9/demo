package cameraRenderUtil

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kainn9/demo/components"
	cameraSharedUtil "github.com/kainn9/demo/systems/util/camera"
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
// Renders the camera's surface to the screen.
// This is done in in the final render phase usually.
func Render(c *components.Camera, screen *ebiten.Image) {

	ops := &ebiten.DrawImageOptions{}
	ops.GeoM.Scale(float64(c.Zoom), float64(c.Zoom))

	tx := cameraSharedUtil.ZoomSpacingX(c.Zoom)
	ty := cameraSharedUtil.ZoomSpacingY(c.Zoom)

	ops.GeoM.Translate(tx, ty)

	screen.DrawImage(c.Surface, ops)

}

// Clears the camera's surface.
func Clear(c *components.Camera) {
	c.Surface.Clear()
}
