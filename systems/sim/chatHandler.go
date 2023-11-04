package simSystems

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"

	UIConstants "github.com/kainn9/demo/constants/UI"
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
		filter.Contains(components.ChatStateAndConfigComponent),
	)
}

func (sys ChatHandlerSystem) Run(dt float64, chatEntity *donburi.Entry) {

	playerEntity := systemsUtil.GetPlayerEntity(sys.scene.World)
	playerState := components.PlayerStateComponent.Get(playerEntity)

	configAndState := components.ChatStateAndConfigComponent.Get(chatEntity)

	popDownSprite := systemsUtil.GetChatPopDownSprite(sys.scene.World)
	popUpSprite := systemsUtil.GetChatPopUpSprite(sys.scene.World)

	// Handle Key Press.
	if playerState.IsInteracting && configAndState.State.Active {
		playerState.IsInteracting = false
		sys.handleNextSlide(configAndState, popDownSprite, popUpSprite)
	}

	sys.handleTransitionState(configAndState, popDownSprite)
	sys.handleClose(configAndState, popDownSprite, popUpSprite)

}

func (sys ChatHandlerSystem) handleNextSlide(
	configAndState *components.ChatStateAndConfig,
	popDownSprite, popUpSprite *components.Sprite,
) {

	// Switch to the PopDown animation.
	configAndState.State.PopDownMode = true
	popDownSprite.StartTick = sys.scene.Manager.TickHandler.CurrentTick()

	// Reset the pop up animation(since it will play after popDown finishes).
	animUtil.ResetAnimationConfig(popUpSprite)

	// Reset the text counter, since it will be incremented in the popUp animation.
	configAndState.State.TextAnimStartTick = -1
	configAndState.State.CurrentSlideIndex++ // Increment the slide index.
}

func (sys ChatHandlerSystem) handleTransitionState(configAndState *components.ChatStateAndConfig, popDownSprite *components.Sprite) {
	// Once the pop down animation is finished, switch to the pop up animation.
	popDownFinished := sys.scene.Manager.TickHandler.TicksSinceNTicks(popDownSprite.StartTick) > UIConstants.CHAT_BOX_ANIM_SPEED*2
	if popDownFinished {
		configAndState.State.PopDownMode = false
		configAndState.State.PopUpMode = true
	}
}

func (sys ChatHandlerSystem) handleClose(configAndState *components.ChatStateAndConfig, popDownSprite, popUpSprite *components.Sprite) {
	// If we are out of chat slides, time to close the chat box and
	// reset the state(incase we ever want to re-open it).
	chatFinished := configAndState.State.CurrentSlideIndex > len(configAndState.State.SlidesContent)-1
	if chatFinished {
		configAndState.State.Active = false
		configAndState.State.CurrentSlideIndex = 0
		animUtil.ResetAnimationConfig(popDownSprite)
		animUtil.ResetAnimationConfig(popUpSprite)
	}
}
