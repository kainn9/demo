package simSystems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"

	assetComponents "github.com/kainn9/demo/components/assets"
	"github.com/kainn9/demo/systems/render/util/animUtil"
	systemsUtil "github.com/kainn9/demo/systems/util"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type ChatHandlerSystem struct {
	scene *coldBrew.Scene
}

func NewChatHandler(scene *coldBrew.Scene) *ChatHandlerSystem {
	return &ChatHandlerSystem{
		scene: scene,
	}
}

func (*ChatHandlerSystem) Query() *donburi.Query {
	return donburi.NewQuery(
		filter.Contains(components.ChatStateComponent),
	)
}

func (sys *ChatHandlerSystem) Run(dt float64, chatEntity *donburi.Entry) {
	config := components.ChatStateComponent.Get(chatEntity)

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) && config.Active {
		config.PopDownMode = true

		popUpEntity := systemsUtil.GetChatPopUpEntity(sys.scene.World)
		popUpSprite := assetComponents.SpriteComponent.Get(popUpEntity)
		animUtil.ResetAnimationConfig(popUpSprite)

		popDownEntity := systemsUtil.GetChatPopDownEntity(sys.scene.World)
		popDownSprite := assetComponents.SpriteComponent.Get(popDownEntity)
		animUtil.ResetAnimationConfig(popDownSprite)

		config.CurrentSlideIndex++
	}
}
