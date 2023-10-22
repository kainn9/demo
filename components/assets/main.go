package assetComponents

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type AssetData struct {
	Loaded bool
	Name   string
}

type Sprite struct {
	*ebiten.Image
	*AssetData
}

type Sound struct {
	*AssetData
}

var SpriteComponent = donburi.NewComponentType[Sprite]()
var SpriteComponents = donburi.NewComponentType[[]*Sprite]()

func NewSprite(name string) *Sprite {
	return &Sprite{
		Image: &ebiten.Image{},
		AssetData: &AssetData{
			Name: name,
		},
	}
}
