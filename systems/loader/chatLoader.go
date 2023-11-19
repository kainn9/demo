package loaderSystems

import (
	"log"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	clientGlobals "github.com/kainn9/demo/globalConfig/client"
	loaderUtil "github.com/kainn9/demo/systems/loader/util"
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
		filter.Contains(components.ChatStateAndConfigComponent),
	)
}

func (sys ChatLoaderSystem) Load(entity *donburi.Entry) {
	configAndState := components.ChatStateAndConfigComponent.Get(entity)
	portraitSprites := components.SpritesSliceComponent.Get(entity)

	sys.loadPortraitSprites(configAndState, portraitSprites)

}

func (sys ChatLoaderSystem) loadPortraitSprites(configAndState *components.ChatStateAndConfig, portraitSprites *[]*components.Sprite) {

	for i, data := range configAndState.State.SlidesContent {

		if (*portraitSprites)[i].AssetData.Loaded {
			continue
		}

		path := clientGlobals.CHARACTER_ASSETS_PREFIX_PATH
		path += data.PortraitName + "/" + clientGlobals.SPRITES_ASSETS_PREFIX_PATH
		path += "portrait"

		log.Println("Loading Portrait Sprite:", path)
		loaderUtil.LoadImage(path, (*portraitSprites)[i])

	}
}
