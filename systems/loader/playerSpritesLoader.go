package loaderSystems

import (
	"log"

	"github.com/kainn9/coldBrew"
	assetComponents "github.com/kainn9/demo/components/assets"
	"github.com/kainn9/demo/constants"
	"github.com/kainn9/demo/queries"
	"github.com/yohamta/donburi"
)

// Todo: add some kind of sorting based on ZIndex.
type PlayerSpritesLoaderSystem struct {
	scene *coldBrew.Scene
}

func NewPlayerSpritesLoader(scene *coldBrew.Scene) *PlayerSpritesLoaderSystem {
	return &PlayerSpritesLoaderSystem{
		scene: scene,
	}
}

func (sys *PlayerSpritesLoaderSystem) Query() *donburi.Query {
	return queries.PlayerQuery
}

func (sys *PlayerSpritesLoaderSystem) Load(entity *donburi.Entry) {

	sprites := assetComponents.SpriteComponents.Get(entity)

	for _, sprite := range *sprites {
		if sprite.AssetData.Loaded {
			continue
		}

		path := constants.PLAYER_SPRITE_PATH + sprite.AssetData.Name
		LoadImage(path, sprite)
	}

	log.Println("Loading Player Sprites.")

}
