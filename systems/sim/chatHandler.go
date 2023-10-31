package simSystems

import (
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"

	UIConstants "github.com/kainn9/demo/constants/UI"
	inputConstants "github.com/kainn9/demo/constants/input"
	animUtil "github.com/kainn9/demo/systems/render/util/anim"
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

func (ChatHandlerSystem) Query() *donburi.Query {
	return donburi.NewQuery(
		filter.Contains(components.ChatStateComponent),
	)
}

func (sys ChatHandlerSystem) Run(dt float64, chatEntity *donburi.Entry) {
	config := components.ChatStateComponent.Get(chatEntity)
	interact := inputConstants.KEY_BINDS[inputConstants.KEY_BIND_INTERACT]

	popDownSprite := systemsUtil.GetChatPopDownSprite(sys.scene.World)
	popUpSprite := systemsUtil.GetChatPopUpSprite(sys.scene.World)

	if inpututil.IsKeyJustPressed(interact) && config.Active {

		// Switch to the PopDown animation.
		config.PopDownMode = true
		popDownSprite.StartTick = sys.scene.Manager.TickHandler.CurrentTick()

		// Reset the pop up animation(since it will play after popDown finishes).
		animUtil.ResetAnimationConfig(popUpSprite)

		// Reset the text counter, since it will be incremented in the popUp animation.
		config.TextAnimStartTick = 0
		config.CurrentSlideIndex++ // Increment the slide index.

	}

	// Once the pop down animation is finished, switch to the pop up animation.
	popDownFinished := sys.scene.Manager.TickHandler.TicksSinceNTicks(popDownSprite.StartTick) > UIConstants.CHAT_BOX_ANIM_SPEED*2
	if popDownFinished {
		config.PopDownMode = false
		config.PopUpMode = true
	}

	// If we are out of chat slides, time to close the chat box and
	// reset the state(incase we ever want to re-open it).
	chatFinished := config.CurrentSlideIndex > len(config.SlidesContent)-1
	if chatFinished {
		config.Active = false
		config.CurrentSlideIndex = 0
		animUtil.ResetAnimationConfig(popDownSprite)
		animUtil.ResetAnimationConfig(popUpSprite)
	}

}
