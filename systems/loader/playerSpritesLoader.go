package loaderSystems

import (
	"log"

	"github.com/kainn9/coldBrew"

	"github.com/kainn9/demo/components"
	clientConstants "github.com/kainn9/demo/constants/client"
	"github.com/kainn9/demo/queries"
	"github.com/yohamta/donburi"
)

type PlayerSpritesLoaderSystem struct{}

func NewPlayerSpritesLoader(scene *coldBrew.Scene) *PlayerSpritesLoaderSystem {
	return &PlayerSpritesLoaderSystem{}
}

func (sys PlayerSpritesLoaderSystem) Query() *donburi.Query {
	return queries.PlayerQuery
}

func (sys PlayerSpritesLoaderSystem) Load(entity *donburi.Entry) {

	spritesMap := components.PlayerSpritesAnimMapComponent.Get(entity)

	for nameKey, sprite := range *spritesMap {
		if sprite.AssetData.Loaded {
			continue
		}

		log.Println("Loading Player Sprite: " + nameKey + ".")

		path := string(clientConstants.PLAYER_ASSETS_SUB_PATH + nameKey)
		LoadImage(path, sprite)
	}

}
