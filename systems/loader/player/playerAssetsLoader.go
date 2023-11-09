package loaderPlayerSystems

import (
	"log"

	"github.com/kainn9/coldBrew"

	loaderUtil "github.com/kainn9/demo/systems/loader/util"

	"github.com/kainn9/demo/components"
	clientGlobals "github.com/kainn9/demo/globalConfig/client"
	"github.com/kainn9/demo/queries"
	"github.com/yohamta/donburi"
)

type PlayerAssetsLoaderSystem struct{}

func NewPlayerAssetsLoader(scene *coldBrew.Scene) *PlayerAssetsLoaderSystem {
	return &PlayerAssetsLoaderSystem{}
}

func (sys PlayerAssetsLoaderSystem) Query() *donburi.Query {
	return queries.PlayerQuery
}

func (sys PlayerAssetsLoaderSystem) Load(entity *donburi.Entry) {

	spritesMap := components.SpritesAnimMapComponent.Get(entity)

	for nameKey, sprite := range *spritesMap {
		if sprite.AssetData.Loaded {
			continue
		}

		path := string(clientGlobals.PLAYER_ASSETS_SUB_PATH + nameKey)

		log.Println("Loading Player Sprite at", path)
		loaderUtil.LoadImage(path, sprite)
	}

}
