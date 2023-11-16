package components

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

var CameraComponent = donburi.NewComponentType[Camera]()

type Camera struct {
	X, Y, Zoom    float64
	Width, Height int
	Surface       *ebiten.Image
}

func NewCamera(x, y float64, w, h int, zoom float64) *Camera {
	return &Camera{
		X:       x,
		Y:       y,
		Surface: ebiten.NewImage(w, h),
		Zoom:    zoom,
	}
}
