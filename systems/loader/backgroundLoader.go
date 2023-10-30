package loaderSystems

import (
	"log"
	"strconv"
	"time"

	"github.com/kainn9/coldBrew"
	assetComponents "github.com/kainn9/demo/components/assets"
	"github.com/kainn9/demo/constants"
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

func (sys *BackgroundLoaderSystem) CustomQueryOne() *donburi.Query {
	return queries.ParallaxBackGroundLayerQuery
}

func (sys *BackgroundLoaderSystem) CustomQueryTwo() *donburi.Query {
	return queries.FrontLayerQuery
}

func (sys *BackgroundLoaderSystem) Load(entity *donburi.Entry) {

	world := sys.scene.World

	parallaxQuery := sys.CustomQueryOne()

	parallaxQuery.Each(world, func(layerEntity *donburi.Entry) {
		parallaxLoader(layerEntity)
	})

	frontLayerQuery := sys.CustomQueryTwo()

	frontLayerQuery.Each(world, func(layerEntity *donburi.Entry) {
		frontLayerLoader(layerEntity)
	})

}

func parallaxLoader(layerEntity *donburi.Entry) {
	sprite := assetComponents.SpriteComponent.Get(layerEntity)

	pLaxLayerConfig := assetComponents.ParallaxLayerConfigComponent.Get(layerEntity)

	path := constants.SCENE_ASSETS_SUB_PATH
	path += pLaxLayerConfig.SceneAssetsPath
	path += strconv.Itoa(pLaxLayerConfig.ZIndex)

	LoadImage(path, sprite)

	// A temp hack since there is so little load time
	// but we want to see the load screen for testing...
	time.Sleep(50 * time.Millisecond)
	log.Println("Loading background layer for ", pLaxLayerConfig.SceneAssetsPath, strconv.Itoa(pLaxLayerConfig.ZIndex)+".")
}

func frontLayerLoader(layerEntity *donburi.Entry) {
	sprite := assetComponents.SpriteComponent.Get(layerEntity)
	frontLayerConfig := assetComponents.FrontLayerComponent.Get(layerEntity)

	path := constants.SCENE_ASSETS_SUB_PATH
	path += frontLayerConfig.SceneAssetPath
	path += "front"

	log.Println("Loading front layer for", path+".")

	LoadImage(path, sprite)
}
