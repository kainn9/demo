package clientUISystems

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"

	UIGlobals "github.com/kainn9/demo/globalConfig/UI"
	soundUtil "github.com/kainn9/demo/systems/client/util/sound"
	animUtil "github.com/kainn9/demo/systems/render/util/anim"
	systemsUtil "github.com/kainn9/demo/systems/util"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type ChatHandlerSystem struct {
	scene *coldBrew.Scene

	CallBackSystems []ChatCallBack
}

type ChatCallBack interface {
	SlideIndex() int
	Callback(*coldBrew.Scene)
	ChatName() string
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
	world := sys.scene.World
	playerEntity := systemsUtil.PlayerEntity(world)
	playerState := components.PlayerStateComponent.Get(playerEntity)

	configAndState := components.ChatStateAndConfigComponent.Get(chatEntity)

	popDownSprite := systemsUtil.ChatPopDownSprite(world)
	popUpSprite := systemsUtil.ChatPopUpSprite(world)

	globalSoundsEntity := systemsUtil.UISoundsSingletonEntity(world)
	globalSounds := components.SoundsMapComponent.Get(globalSoundsEntity)

	audContext := components.AudioContextComponent.Get(globalSoundsEntity).Context

	chatNewSound := (*globalSounds)[UIGlobals.CHAT_BOX_NEW_SOUND_NAME]

	if configAndState.State.JustOpened {
		soundUtil.PlaySound(audContext, chatNewSound)
	}

	if configAndState.State.Active {
		configAndState.State.JustOpened = false
	}

	// Handle Key Press.
	if playerState.IsInteracting && configAndState.State.Active {
		playerState.IsInteracting = false
		sys.handleNextSlide(configAndState, popDownSprite, popUpSprite)
		soundUtil.PlaySound(audContext, chatNewSound)
	}

	sys.handleTransitionState(configAndState, popDownSprite)
	sys.handleClose(configAndState, popDownSprite, popUpSprite, playerState)

	if configAndState.State.Active {
		sys.handleCallback(configAndState)
	}
}

func (sys ChatHandlerSystem) handleNextSlide(
	configAndState *components.ChatStateAndConfig,
	popDownSprite, popUpSprite *components.Sprite,
) {
	// Reset the text counter, since it will be incremented in the popUp animation.
	configAndState.State.TextAnimStartTick = -1
	configAndState.State.CurrentSlideIndex++ // Increment the slide index.

	// Exit early and don't reset slide animations if speaker does not change.
	slideIndex := configAndState.State.CurrentSlideIndex
	prevSlideIndex := slideIndex - 1
	slides := configAndState.State.SlidesContent

	if prevSlideIndex >= 0 && slideIndex < len(slides) {

		currentSlidePortName := slides[slideIndex].PortraitName
		prevSlideName := slides[prevSlideIndex].PortraitName

		if currentSlidePortName == prevSlideName {
			return
		}

	}

	configAndState.State.NameTextAnimStartTick = -1

	// Switch to the PopDown animation.
	configAndState.State.PopDownMode = true
	popDownSprite.StartTick = sys.scene.Manager.TickHandler.CurrentTick()

	// Reset the pop up animation(since it will play after popDown finishes).
	animUtil.ResetAnimationConfig(popUpSprite)

}

func (sys ChatHandlerSystem) handleCallback(stateAndConfig *components.ChatStateAndConfig) {
	if sys.CallBackSystems == nil {
		return
	}

	for _, callback := range sys.CallBackSystems {
		matchingNames := callback.ChatName() == stateAndConfig.Config.ChatName
		matchingIndex := callback.SlideIndex() == stateAndConfig.State.CurrentSlideIndex

		if matchingNames && matchingIndex {
			callback.Callback(sys.scene)
		}
	}

}

func (sys ChatHandlerSystem) handleTransitionState(configAndState *components.ChatStateAndConfig, popDownSprite *components.Sprite) {
	// Once the pop down animation is finished, switch to the pop up animation.
	popDownFinished := sys.scene.Manager.TickHandler.TicksSinceNTicks(popDownSprite.StartTick) > UIGlobals.CHAT_BOX_ANIM_SPEED*3

	if popDownFinished {
		configAndState.State.PopDownMode = false
	}
}

func (sys ChatHandlerSystem) handleClose(
	configAndState *components.ChatStateAndConfig,
	popDownSprite,
	popUpSprite *components.Sprite,
	playerState *components.PlayerState,
) {
	// If we are out of chat slides, time to close the chat box and
	// reset the state(incase we ever want to re-open it).
	chatFinished := configAndState.State.CurrentSlideIndex > len(configAndState.State.SlidesContent)-1
	if chatFinished {
		// Also handle the callback when chat is finished.
		// For post chat callbacks(aka slide length index).
		sys.handleCallback(configAndState)

		configAndState.State.Active = false
		configAndState.State.CurrentSlideIndex = 0

		configAndState.State.PopUpMode = false
		animUtil.ResetAnimationConfig(popDownSprite)
		animUtil.ResetAnimationConfig(popUpSprite)

	}
}
