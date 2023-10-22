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
type ParallaxBackgroundLoaderSystem struct {
	scene *coldBrew.Scene
}

func NewParallaxBackgroundLoader(scene *coldBrew.Scene) *ParallaxBackgroundLoaderSystem {
	return &ParallaxBackgroundLoaderSystem{
		scene: scene,
	}
}

func (sys *ParallaxBackgroundLoaderSystem) Query() *donburi.Query {
	return queries.ParallaxBackGroundLayerQuery
}

func (sys *ParallaxBackgroundLoaderSystem) Load(entity *donburi.Entry) {

	sprite := assetComponents.SpriteComponent.Get(entity)

	pLaxLayerConfig := assetComponents.ParallaxLayerConfigComponent.Get(entity)

	path := constants.SCENE_SUB_DIR
	path += pLaxLayerConfig.SubPath
	path += strconv.Itoa(pLaxLayerConfig.ZIndex)

	LoadImage(path, sprite)

	// A temp hack since there is so little load time
	// but we want to see the load screen for testing...
	time.Sleep(50 * time.Millisecond)
	log.Println("Loading layer", strconv.Itoa(pLaxLayerConfig.ZIndex), ".")

}
