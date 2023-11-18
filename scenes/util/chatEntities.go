package scenesUtil

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	UIGlobals "github.com/kainn9/demo/globalConfig/UI"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
	"github.com/yohamta/donburi"
)

func AddBasicChatEntity(
	scene *coldBrew.Scene,
	chatName string,
	content []components.SlidesContent,
	auto bool,
) *donburi.Entry {

	chatEntity := scene.AddEntity(
		components.ChatStateAndConfigComponent,
		components.SpritesSliceComponent,
	)

	configAndState := components.NewChatStateAndConfig(chatName, content)
	configAndState.State.Active = auto
	configAndState.State.PopUpMode = auto

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

	return chatEntity

}

func AddOnCollideChatEntity(
	scene *coldBrew.Scene,
	chatName string,
	content []components.SlidesContent,
	x, y, width, height float64,
) {

	// Chat Entity.
	chatEntity := scene.AddEntity(
		components.ChatStateAndConfigComponent,
		components.SpritesSliceComponent,
		components.RigidBodyComponent,
	)

	configAndState := components.NewChatStateAndConfig(chatName, content)

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

	body := tBokiComponents.NewRigidBodyBox(x, y, width, height, 0, false)
	components.RigidBodyComponent.SetValue(chatEntity, *body)

}

func AddOnInteractChatEntity(
	scene *coldBrew.Scene,
	chatName string,
	content []components.SlidesContent,
	x, y, width, height float64,
) {

	// Chat Entity.
	chatEntity := scene.AddEntity(
		components.ChatStateAndConfigComponent,
		components.SpritesSliceComponent,
		components.RigidBodyComponent,
		components.IndicatorStateAndConfigComponent,
	)

	// Chat State.
	configAndState := components.NewChatStateAndConfig(chatName, content)
	components.ChatStateAndConfigComponent.SetValue(
		chatEntity,
		*configAndState,
	)

	// Sprites.
	portraitSprites := make([]*components.Sprite, len(content))

	for i := range portraitSprites {
		portraitSprites[i] = components.NewSprite(0, 0)
	}

	components.SpritesSliceComponent.SetValue(
		chatEntity,
		portraitSprites,
	)

	// Body.
	body := tBokiComponents.NewRigidBodyBox(x, y, width, height, 0, false)
	components.RigidBodyComponent.SetValue(chatEntity, *body)

	// Indicator State
	offX := UIGlobals.IndicatorPlayerOffsets[UIGlobals.CurrentLayout][UIGlobals.INDICATOR_INTERACT].X
	offY := UIGlobals.IndicatorPlayerOffsets[UIGlobals.CurrentLayout][UIGlobals.INDICATOR_INTERACT].Y

	indicatorState := components.NewIndicatorStateAndConfig(
		offX, offY,
		false,
		true,
		UIGlobals.INDICATOR_INTERACT,
	)

	components.IndicatorStateAndConfigComponent.SetValue(chatEntity, *indicatorState)

}
