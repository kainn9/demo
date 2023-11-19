package loaderSystems

import (
	"log"
	"strconv"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	clientGlobals "github.com/kainn9/demo/globalConfig/client"
	"github.com/kainn9/demo/queries"
	loaderUtil "github.com/kainn9/demo/systems/loader/util"
	"github.com/yohamta/donburi"
)

// Todo: add some kind of sorting based on ZIndex.
type BackgroundLoaderSystem struct {
	scene *coldBrew.Scene
}

func NewBackgroundLoader(scene *coldBrew.Scene) *BackgroundLoaderSystem {
	return &BackgroundLoaderSystem{
		scene: scene,
	}
}

func (sys BackgroundLoaderSystem) parallaxBGQuery() *donburi.Query {
	return queries.ParallaxBackGroundLayerQuery
}

func (sys BackgroundLoaderSystem) frontLayerBGQuery() *donburi.Query {
	return queries.FrontLayerQuery
}

func (sys BackgroundLoaderSystem) bgSoundQuery() *donburi.Query {
	return queries.BackgroundSoundQuery
}

func (sys BackgroundLoaderSystem) Load(entity *donburi.Entry) {

	world := sys.scene.World

	sys.parallaxBGQuery().Each(world, func(layerEntity *donburi.Entry) {
		sys.parallaxLoader(layerEntity)
	})

	sys.frontLayerBGQuery().Each(world, func(layerEntity *donburi.Entry) {
		sys.frontLayerLoader(layerEntity)
	})

	sys.bgSoundQuery().Each(world, func(soundEntity *donburi.Entry) {
		sys.bgSoundLoader(soundEntity)
	})

}

func (sys BackgroundLoaderSystem) parallaxLoader(layerEntity *donburi.Entry) {
	sprite := components.SpriteComponent.Get(layerEntity)

	pLaxLayerConfig := components.ParallaxLayerConfigComponent.Get(layerEntity)

	path := clientGlobals.SCENE_ASSETS_SUB_PATH
	path += pLaxLayerConfig.SceneAssetsPath
	path += strconv.Itoa(pLaxLayerConfig.ZIndex)

	loaderUtil.LoadImage(path, sprite)

	log.Println("Loading background layer for ", pLaxLayerConfig.SceneAssetsPath, strconv.Itoa(pLaxLayerConfig.ZIndex)+".")
}

func (sys BackgroundLoaderSystem) frontLayerLoader(layerEntity *donburi.Entry) {
	sprite := components.SpriteComponent.Get(layerEntity)
	frontLayerConfig := components.FrontLayerComponent.Get(layerEntity)

	path := clientGlobals.SCENE_ASSETS_SUB_PATH
	path += frontLayerConfig.SceneAssetPath
	path += "front"

	log.Println("Loading front layer for", path+".")

	loaderUtil.LoadImage(path, sprite)
}

func (sys BackgroundLoaderSystem) bgSoundLoader(soundEntity *donburi.Entry) {

	config := components.BgSoundConfigComponent.Get(soundEntity)

	path := clientGlobals.SCENE_ASSETS_SUB_PATH + config.SceneAssetsPath + clientGlobals.BG_SOUND_NAME
	sound := components.SoundComponent.Get(soundEntity)

	loaderUtil.LoadSound(path, sound)
}
