package renderSystems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	assetComponents "github.com/kainn9/demo/components/assets"
	UIConstants "github.com/kainn9/demo/constants/UI"
	animUtil "github.com/kainn9/demo/systems/render/util/anim"
	cameraUtil "github.com/kainn9/demo/systems/render/util/camera"
	textUtil "github.com/kainn9/demo/systems/render/util/text"
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

func (ChatSlidesRendererSystem) Query() *donburi.Query {
	return donburi.NewQuery(
		filter.Contains(components.ChatStateComponent),
	)
}

func (sys ChatSlidesRendererSystem) Draw(screen *ebiten.Image, chatEntity *donburi.Entry) {

	// Config/State.
	config := components.ChatStateComponent.Get(chatEntity)

	// Grab the camera.
	cameraEntity := systemsUtil.GetCameraEntity(sys.scene.World)
	camera := components.CameraComponent.Get(cameraEntity)

	// Chat Box Draw options.
	charBoxOpts := &ebiten.DrawImageOptions{}
	charBoxOpts.GeoM.Translate(10, 10)
	popDownSprite := systemsUtil.GetChatPopDownSprite(sys.scene.World)

	// Handle the pop down animation one last time,
	// when the chat box is no longer active.
	if !config.Active && popDownSprite.StartTick == 0 {
		sys.renderPopDownAnimation(popDownSprite, charBoxOpts, camera, config)
		return
	}

	// If the chat box is not active, and the pop down animation is finished,
	// then we don't need to render anything.
	popDownFinished := sys.scene.Manager.TickHandler.TicksSinceNTicks(popDownSprite.StartTick) > UIConstants.CHAT_BOX_ANIM_SPEED*2
	if !config.Active && popDownFinished {
		return
	}

	// Handle the pop down animation, when the chat box is active,
	// and where switching slides(pop down then pop up).
	if config.PopDownMode {
		sys.renderPopDownAnimation(popDownSprite, charBoxOpts, camera, config)
		return
	}

	// Handle the pop up animation.
	if config.PopUpMode {
		finished := sys.renderPopUpAnimation(charBoxOpts, camera)

		if finished {
			// Handle the portrait.
			portraitSprites := assetComponents.SpritesSliceComponent.Get(chatEntity)
			sys.handlePortrait(camera, config, portraitSprites)

			// Handle the text.
			slideContent := config.SlidesContent[config.CurrentSlideIndex]

			if config.TextAnimStartTick == 0 {
				config.TextAnimStartTick = sys.scene.Manager.TickHandler.CurrentTick()
			}

			textUtil.RenderTextDefault(
				slideContent.Text,
				120,
				45,
				525,
				config.TextAnimStartTick,
				config.TicksPerWord,
				camera,
				&sys.scene.World,
				sys.scene.Manager,
			)
		}

		return
	}

}

func (sys ChatSlidesRendererSystem) renderPopUpAnimation(chatBoxOpts *ebiten.DrawImageOptions, camera *components.Camera) (finished bool) {

	popUpSprite := systemsUtil.GetChatPopUpSprite(sys.scene.World)

	spriteAtFrameIndex := animUtil.PlayAnim(sys.scene.Manager, popUpSprite)

	cameraUtil.AddImage(camera, spriteAtFrameIndex, chatBoxOpts)

	finished = sys.scene.Manager.TickHandler.TicksSinceNTicks(popUpSprite.StartTick) > UIConstants.CHAT_BOX_ANIM_SPEED*2

	return finished
}

func (sys ChatSlidesRendererSystem) renderPopDownAnimation(popDownSprite *assetComponents.Sprite, chatBoxOpts *ebiten.DrawImageOptions, camera *components.Camera, config *components.ChatStateAndConfig) {

	spriteAtFrameIndex := animUtil.PlayAnim(sys.scene.Manager, popDownSprite)

	cameraUtil.AddImage(camera, spriteAtFrameIndex, chatBoxOpts)

}

func (sys ChatSlidesRendererSystem) handlePortrait(camera *components.Camera, config *components.ChatStateAndConfig, portraitSprites *[]*assetComponents.Sprite) {
	portraitOpts := &ebiten.DrawImageOptions{}
	portraitOpts.GeoM.Translate(35, 25)
	cameraUtil.AddImage(camera, (*portraitSprites)[config.CurrentSlideIndex].Image, portraitOpts)
}
