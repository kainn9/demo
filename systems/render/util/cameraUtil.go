package renderSystems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kainn9/demo/components"
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

func SetPosition(c *components.Camera, x, y float64) {
	c.X = x
	c.Y = y
}

func Clear(c *components.Camera) {
	c.Surface.Clear()
}
