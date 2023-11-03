package components

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type AnimState string

type Sprite struct {
	OffSetX, OffSetY float64
	*ebiten.Image
	*AssetData

	// This is optional, only needed for sprites that are animated(aka spriteSheets).
	*AnimationConfig
}

var SpriteComponent = donburi.NewComponentType[Sprite]()
var SpritesMapComponent = donburi.NewComponentType[map[string]*Sprite]()
var SpritesSliceComponent = donburi.NewComponentType[[]*Sprite]()
var PlayerSpritesAnimMapComponent = donburi.NewComponentType[map[AnimState]*Sprite]()

// This should work, but its not needed for now.
// var MultipleSpritesSliceComponent = donburi.NewComponentType[map[string][]*Sprite]()

func NewSprite(offX, offY float64) *Sprite {
	return &Sprite{
		OffSetX:   offX,
		OffSetY:   offY,
		Image:     &ebiten.Image{},
		AssetData: &AssetData{},
	}
}
