package loaderSystems

import (
	"log"
	"strconv"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	clientConstants "github.com/kainn9/demo/constants/client"
	"github.com/kainn9/demo/queries"
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

func (sys BackgroundLoaderSystem) ParallaxBGQuery() *donburi.Query {
	return queries.ParallaxBackGroundLayerQuery
}

func (sys BackgroundLoaderSystem) FrontLayerBGQuery() *donburi.Query {
	return queries.FrontLayerQuery
}

func (sys BackgroundLoaderSystem) Load(entity *donburi.Entry) {

	world := sys.scene.World

	parallaxQuery := sys.ParallaxBGQuery()

	parallaxQuery.Each(world, func(layerEntity *donburi.Entry) {
		sys.parallaxLoader(layerEntity)
	})

	frontLayerQuery := sys.FrontLayerBGQuery()

	frontLayerQuery.Each(world, func(layerEntity *donburi.Entry) {
		sys.frontLayerLoader(layerEntity)
	})

}

func (sys BackgroundLoaderSystem) parallaxLoader(layerEntity *donburi.Entry) {
	sprite := components.SpriteComponent.Get(layerEntity)

	pLaxLayerConfig := components.ParallaxLayerConfigComponent.Get(layerEntity)

	path := clientConstants.SCENE_ASSETS_SUB_PATH
	path += pLaxLayerConfig.SceneAssetsPath
	path += strconv.Itoa(pLaxLayerConfig.ZIndex)

	LoadImage(path, sprite)

	log.Println("Loading background layer for ", pLaxLayerConfig.SceneAssetsPath, strconv.Itoa(pLaxLayerConfig.ZIndex)+".")
}

func (sys BackgroundLoaderSystem) frontLayerLoader(layerEntity *donburi.Entry) {
	sprite := components.SpriteComponent.Get(layerEntity)
	frontLayerConfig := components.FrontLayerComponent.Get(layerEntity)

	path := clientConstants.SCENE_ASSETS_SUB_PATH
	path += frontLayerConfig.SceneAssetPath
	path += "front"

	log.Println("Loading front layer for", path+".")

	LoadImage(path, sprite)
}
