package loaderSystems

import (
	"log"

	"github.com/kainn9/coldBrew"
	assetComponents "github.com/kainn9/demo/components/assets"
	"github.com/kainn9/demo/constants"
	"github.com/kainn9/demo/queries"
	"github.com/yohamta/donburi"
)

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

	spritesMap := assetComponents.SpritesMapComponent.Get(entity)

	for nameKey, sprite := range *spritesMap {
		if sprite.AssetData.Loaded {
			continue
		}

		log.Println("Loading Player Sprite: " + nameKey + ".")

		path := constants.PLAYER_SPRITE_PATH + nameKey
		LoadImage(path, sprite)
	}

}
