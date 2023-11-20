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

type NpcAssetLoaderSystem struct {
	loadedNpcs map[components.NpcName]*donburi.Entry
}

func NewNpcAssetLoader(scene *coldBrew.Scene) *NpcAssetLoaderSystem {
	return &NpcAssetLoaderSystem{
		loadedNpcs: make(map[components.NpcName]*donburi.Entry),
	}
}

func (sys NpcAssetLoaderSystem) Query() *donburi.Query {

	return queries.NpcQuery
}

func (sys NpcAssetLoaderSystem) Load(npcEntity *donburi.Entry) {

	spritesMap := components.SpritesCharStateMapComponent.Get(npcEntity)
	config := components.NpcConfigComponent.Get(npcEntity)

	if sys.reuseExistingNpcSprites(config, spritesMap) {
		return
	}

	sys.loadSprites(config, spritesMap)

	sys.loadedNpcs[config.Name] = npcEntity
}

func (sys NpcAssetLoaderSystem) reuseExistingNpcSprites(config *components.NpcConfig, spritesMap *map[components.CharState]*components.Sprite) bool {
	if matchedNpcEntity, ok := sys.loadedNpcs[config.Name]; ok {

		matchedSpritesMap := components.SpritesCharStateMapComponent.Get(matchedNpcEntity)

		for nameKey, sprite := range *matchedSpritesMap {
			(*spritesMap)[nameKey].Image = sprite.Image
		}

		return true
	}

	return false
}

func (sys NpcAssetLoaderSystem) loadSprites(config *components.NpcConfig, spritesMap *map[components.CharState]*components.Sprite) {
	for nameKey, sprite := range *spritesMap {
		if sprite.AssetData.Loaded {
			continue
		}

		path := string(clientGlobals.CHARACTER_ASSETS_PREFIX_PATH)
		path += string(config.Name) + "/"
		path += string(clientGlobals.SPRITES_ASSETS_PREFIX_PATH)
		path += string(nameKey)

		log.Println("Loading Npc Sprite at", path)

		loaderUtil.LoadImage(path, sprite)
	}

}
