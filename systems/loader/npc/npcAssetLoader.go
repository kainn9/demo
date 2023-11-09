package loaderNpcSystems

import (
	"log"

	"github.com/kainn9/coldBrew"

	"github.com/kainn9/demo/components"
	clientGlobals "github.com/kainn9/demo/globalConfig/client"
	"github.com/kainn9/demo/queries"
	loaderUtil "github.com/kainn9/demo/systems/loader/util"
	"github.com/yohamta/donburi"
)

type NpcAssetLoaderSystem struct{}

func NewNpcAssetLoader(scene *coldBrew.Scene) *NpcAssetLoaderSystem {
	return &NpcAssetLoaderSystem{}
}

func (sys NpcAssetLoaderSystem) Query() *donburi.Query {

	return queries.NpcQuery
}

func (sys NpcAssetLoaderSystem) Load(npcEntity *donburi.Entry) {

	spritesMap := components.SpritesAnimMapComponent.Get(npcEntity)
	config := components.NpcConfigComponent.Get(npcEntity)

	for nameKey, sprite := range *spritesMap {
		if sprite.AssetData.Loaded {
			continue
		}

		path := string(clientGlobals.CHARACTER_ASSETS_SUB_PATH)
		path += string(config.Name) + "/"
		path += string(nameKey)

		log.Println("Loading Npc Sprite at", path)

		loaderUtil.LoadImage(path, sprite)
	}

}
