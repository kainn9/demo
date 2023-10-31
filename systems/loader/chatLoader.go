package loaderSystems

import (
	"log"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	assetComponents "github.com/kainn9/demo/components/assets"
	clientConstants "github.com/kainn9/demo/constants/client"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type ChatLoaderSystem struct {
	scene *coldBrew.Scene
}

func NewChatLoader(scene *coldBrew.Scene) *ChatLoaderSystem {
	return &ChatLoaderSystem{
		scene: scene,
	}
}

func (sys ChatLoaderSystem) Query() *donburi.Query {
	return donburi.NewQuery(
		filter.Contains(components.ChatStateComponent),
	)
}

func (sys ChatLoaderSystem) Load(entity *donburi.Entry) {
	config := components.ChatStateComponent.Get(entity)
	portraitSprites := assetComponents.SpritesSliceComponent.Get(entity)

	// sys.loadChatPreReqAssets(sys.scene)
	sys.loadPortraitSprites(config, portraitSprites)

}

func (sys ChatLoaderSystem) loadPortraitSprites(config *components.ChatStateAndConfig, portraitSprites *[]*assetComponents.Sprite) {

	for i, data := range config.SlidesContent {

		if (*portraitSprites)[i].AssetData.Loaded {
			continue
		}

		path := clientConstants.CHARACTER_ASSETS_SUB_PATH
		path += data.PortraitName + "/"
		path += "portrait"

		log.Println("Loading Portrait Sprite:", path)
		LoadImage(path, (*portraitSprites)[i])

	}
}
