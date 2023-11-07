package loaderSystems

import (
	"log"

	"github.com/kainn9/coldBrew"

	"github.com/kainn9/demo/components"
	clientConstants "github.com/kainn9/demo/constants/client"
	"github.com/kainn9/demo/queries"
	"github.com/yohamta/donburi"
)

type PlayerCarLoaderSystem struct{}

func NewPlayerCarLoader(scene *coldBrew.Scene) *PlayerCarLoaderSystem {
	return &PlayerCarLoaderSystem{}
}

func (sys PlayerCarLoaderSystem) Query() *donburi.Query {

	return queries.PlayerCarQuery

}

func (sys PlayerCarLoaderSystem) Load(entity *donburi.Entry) {

	spritesMap := components.SpritesMapComponent.Get(entity)

	for nameKey, sprite := range *spritesMap {
		if sprite.AssetData.Loaded {
			continue
		}

		log.Println("Loading Car Sprite: " + nameKey + ".")

		path := string(clientConstants.PLAYER_CAR_ASSETS_SUBPATH + nameKey)
		LoadImage(path, sprite)
	}

}
