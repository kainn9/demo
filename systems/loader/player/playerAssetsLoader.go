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

func (sys PlayerAssetsLoaderSystem) Load(playerEntity *donburi.Entry) {
	sys.loadSprites(playerEntity)
	sys.loadSounds(playerEntity)
}

func (sys PlayerAssetsLoaderSystem) loadSprites(playerEntity *donburi.Entry) {

	spritesMap := components.SpritesCharStateMapComponent.Get(playerEntity)

	for nameKey, sprite := range *spritesMap {
		if sprite.AssetData.Loaded {
			continue
		}

		path := string(clientGlobals.CHARACTER_ASSETS_PREFIX_PATH + clientGlobals.PLAYER_PREFIX_PATH + clientGlobals.SPRITES_ASSETS_PREFIX_PATH + nameKey)

		log.Println("Loading Player Sprite at", path)
		loaderUtil.LoadImage(path, sprite)
	}
}

func (sys PlayerAssetsLoaderSystem) loadSounds(playerEntity *donburi.Entry) {

	soundsMap := components.SoundCharStateMapComponent.Get(playerEntity)

	for nameKey, sound := range *soundsMap {
		if sound.AssetData.Loaded {
			continue
		}

		path := string(clientGlobals.CHARACTER_ASSETS_PREFIX_PATH + clientGlobals.PLAYER_PREFIX_PATH + clientGlobals.SOUNDS_ASSETS_PREFIX_PATH + nameKey)

		log.Println("Loading Player Sound at", path)
		loaderUtil.LoadSound(path, sound)
	}
}
