package scenesUtil

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	assetComponents "github.com/kainn9/demo/components/assets"
)

func AddChatEntity(
	scene *coldBrew.Scene,
	chatName, sceneAssetPath string,
	ticksPerWord int,
	content []components.SlidesContent,
) {

	chatEntity := scene.AddEntity(
		components.ChatStateComponent,
		assetComponents.SpritesSliceComponent,
	)

	chatState := components.NewChatStateAndConfig(chatName, sceneAssetPath, ticksPerWord, content)
	chatState.Active = true
	chatState.PopUpMode = true

	components.ChatStateComponent.SetValue(
		chatEntity,
		*chatState,
	)

	portraitSprites := make([]*assetComponents.Sprite, len(content))

	for i := range portraitSprites {
		portraitSprites[i] = assetComponents.NewSprite(0, 0)
	}

	assetComponents.SpritesSliceComponent.SetValue(
		chatEntity,
		portraitSprites,
	)

}
