package scenesUtil

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
)

func AddChatEntity(
	scene *coldBrew.Scene,
	chatName, sceneAssetPath string,
	ticksPerWord int,
	content []components.SlidesContent,
) {

	chatEntity := scene.AddEntity(
		components.ChatStateAndConfigComponent,
		components.SpritesSliceComponent,
	)

	configAndState := components.NewChatStateAndConfig(chatName, sceneAssetPath, ticksPerWord, content)
	configAndState.State.Active = true
	configAndState.State.PopUpMode = true

	components.ChatStateAndConfigComponent.SetValue(
		chatEntity,
		*configAndState,
	)

	portraitSprites := make([]*components.Sprite, len(content))

	for i := range portraitSprites {
		portraitSprites[i] = components.NewSprite(0, 0)
	}

	components.SpritesSliceComponent.SetValue(
		chatEntity,
		portraitSprites,
	)

}
