package renderSystems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	UIGlobals "github.com/kainn9/demo/globalConfig/UI"
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
		filter.Contains(components.ChatStateAndConfigComponent),
	)
}

func (sys ChatSlidesRendererSystem) Draw(screen *ebiten.Image, chatEntity *donburi.Entry) {

	// Config/State.
	configAndState := components.ChatStateAndConfigComponent.Get(chatEntity)

	// Grab the camera.
	cameraEntity := systemsUtil.GetCameraEntity(sys.scene.World)
	camera := components.CameraComponent.Get(cameraEntity)

	// Chat Box Draw options.
	charBoxOpts := &ebiten.DrawImageOptions{}
	charBoxOpts.GeoM.Translate(10, 10)
	popDownSprite := systemsUtil.GetChatPopDownSprite(sys.scene.World)

	// Handle the pop down animation one last time,
	// when the chat box is no longer active.
	if !configAndState.State.Active && animUtil.InactiveAnimation(popDownSprite.AnimationConfig) {
		sys.renderPopDownAnimation(popDownSprite, charBoxOpts, camera)
		return
	}

	// If the chat box is not active, and the pop down animation is finished,
	// then we don't need to render anything.
	popDownFinished := sys.scene.Manager.TickHandler.TicksSinceNTicks(popDownSprite.StartTick) > UIGlobals.CHAT_BOX_ANIM_SPEED*2
	if !configAndState.State.Active && popDownFinished {
		return
	}

	// Handle the pop down animation, when the chat box is active,
	// and where switching slides(pop down then pop up).
	if configAndState.State.PopDownMode {
		sys.renderPopDownAnimation(popDownSprite, charBoxOpts, camera)
		return
	}

	// Handle the pop up animation.
	if configAndState.State.PopUpMode {
		finished := sys.renderPopUpAnimation(charBoxOpts, camera)

		if finished {
			// Handle the portrait.
			portraitSprites := components.SpritesSliceComponent.Get(chatEntity)
			sys.handlePortrait(configAndState, camera, portraitSprites)

			// Handle the text.
			slideContent := configAndState.State.SlidesContent[configAndState.State.CurrentSlideIndex]

			if configAndState.State.TextAnimStartTick == -1 {
				configAndState.State.TextAnimStartTick = sys.scene.Manager.TickHandler.CurrentTick()
			}

			textUtil.RenderTextDefault(
				slideContent.Text,
				120,
				45,
				525,
				configAndState.State.TextAnimStartTick,
				configAndState.Config.TicksPerWord,
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

	spriteAtFrameIndex := animUtil.GetAnimFrame(sys.scene.Manager, popUpSprite)

	cameraUtil.AddImage(camera, spriteAtFrameIndex, chatBoxOpts)

	finished = sys.scene.Manager.TickHandler.TicksSinceNTicks(popUpSprite.StartTick) > UIGlobals.CHAT_BOX_ANIM_SPEED*2

	return finished
}

func (sys ChatSlidesRendererSystem) renderPopDownAnimation(popDownSprite *components.Sprite, chatBoxOpts *ebiten.DrawImageOptions, camera *components.Camera) {

	spriteAtFrameIndex := animUtil.GetAnimFrame(sys.scene.Manager, popDownSprite)

	cameraUtil.AddImage(camera, spriteAtFrameIndex, chatBoxOpts)

}

func (sys ChatSlidesRendererSystem) handlePortrait(configAndState *components.ChatStateAndConfig, camera *components.Camera, portraitSprites *[]*components.Sprite) {
	portraitOpts := &ebiten.DrawImageOptions{}
	index := configAndState.State.CurrentSlideIndex
	image := (*portraitSprites)[index].Image
	imgWidth := image.Bounds().Size()
	xTrans := 35.0

	if !configAndState.State.SlidesContent[index].FacingRight {
		portraitOpts.GeoM.Scale(-1, 1)
		portraitOpts.GeoM.Translate(float64(imgWidth.X)+xTrans, 25)
	} else {
		portraitOpts.GeoM.Translate(xTrans, 25)
	}

	cameraUtil.AddImage(camera, image, portraitOpts)
}
