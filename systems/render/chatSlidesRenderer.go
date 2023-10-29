package renderSystems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	assetComponents "github.com/kainn9/demo/components/assets"
	"github.com/kainn9/demo/constants"
	"github.com/kainn9/demo/systems/render/util/animUtil"
	cameraUtil "github.com/kainn9/demo/systems/render/util/camera"
	systemsUtil "github.com/kainn9/demo/systems/util"
)

type ChatSlidesRendererSystem struct {
	scene *coldBrew.Scene
}

func NewChatSlidesRenderer(scene *coldBrew.Scene) *ChatSlidesRendererSystem {
	return &ChatSlidesRendererSystem{
		scene: scene,
	}
}

func (*ChatSlidesRendererSystem) Query() *donburi.Query {
	return donburi.NewQuery(
		filter.Contains(components.ChatStateComponent),
	)
}

func (sys *ChatSlidesRendererSystem) Draw(screen *ebiten.Image, chatEntity *donburi.Entry) {

	slideSprites := assetComponents.SpritesSliceComponent.Get(chatEntity)
	config := components.ChatStateComponent.Get(chatEntity)

	indexOutOfSlideRange := config.CurrentSlideIndex >= len(*slideSprites)
	if indexOutOfSlideRange {
		config.Active = false
		return
	}

	// Get the portrait.
	portraitSpritesMap := assetComponents.SpritesMapComponent.Get(chatEntity)
	portraitKey := config.PortraitNames[config.CurrentSlideIndex]
	portraitSprite := (*portraitSpritesMap)[portraitKey]

	// Grab the camera.
	cameraEntity := systemsUtil.GetCameraEntity(sys.scene.World)
	camera := components.CameraComponent.Get(cameraEntity)

	// Draw options.
	charBubbleOpts := &ebiten.DrawImageOptions{}
	charBubbleOpts.GeoM.Translate(10, 10)

	// Handle the pop up animation.
	if config.PopUpMode {

		popUpEntity := systemsUtil.GetChatPopUpEntity(sys.scene.World)
		popUpSprite := assetComponents.SpriteComponent.Get(popUpEntity)

		spriteAtFrameIndex := animUtil.GetFrame(sys.scene.Manager, popUpSprite)
		popUpExpired := sys.scene.Manager.TickHandler.TicksSinceNTicks(popUpSprite.StartTick) > constants.CHAT_POP_UP_TICKS

		if popUpExpired {
			config.PopUpMode = false
		}

		cameraUtil.AddImage(camera, spriteAtFrameIndex, charBubbleOpts)
		return
	}

	// Handle the pop down animation.
	if config.PopDownMode {

		popDownEntity := systemsUtil.GetChatPopDownEntity(sys.scene.World)
		popDownSprite := assetComponents.SpriteComponent.Get(popDownEntity)

		spriteAtFrameIndex := animUtil.GetFrame(sys.scene.Manager, popDownSprite)
		popUpExpired := sys.scene.Manager.TickHandler.TicksSinceNTicks(popDownSprite.StartTick) > constants.CHAT_POP_UP_TICKS

		// If the pop down animation is over, stage the pop up animation
		// for the next slide.
		if popUpExpired {
			config.PopDownMode = false
			config.PopUpMode = true
		}

		// No need to stage the pop up animation, if theres no next slide.
		if indexOutOfSlideRange {
			config.PopUpMode = false
		}

		cameraUtil.AddImage(camera, spriteAtFrameIndex, charBubbleOpts)
		return
	}

	// If the chat is not active, don't render anything.
	// We allow the pop down animation to finish though.
	// So its down here.
	if !config.Active {
		return
	}

	// Handle the slide/text animation.
	activeSlideSprite := (*slideSprites)[config.CurrentSlideIndex]

	spriteAtFrameIndex := animUtil.GetFrame(sys.scene.Manager, activeSlideSprite)

	cameraUtil.AddImage(camera, spriteAtFrameIndex, charBubbleOpts)

	// Handle the portrait.
	portraitOpts := &ebiten.DrawImageOptions{}
	portraitOpts.GeoM.Translate(35, 25)
	cameraUtil.AddImage(camera, portraitSprite.Image, portraitOpts)

}
