package renderSystems

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	UIGlobals "github.com/kainn9/demo/globalConfig/UI"
	animUtil "github.com/kainn9/demo/systems/render/util/anim"
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

	boxX := 80.0
	boxY := 265.0
	// Chat Box Draw options.
	charBoxOpts := &ebiten.DrawImageOptions{}
	charBoxOpts.GeoM.Translate(boxX, boxY)
	popDownSprite := systemsUtil.GetChatPopDownSprite(sys.scene.World)

	// Handle the pop down animation one last time,
	// when the chat box is no longer active.
	if !configAndState.State.Active && animUtil.InactiveAnimation(popDownSprite.AnimationConfig) {
		sys.renderPopDownAnimation(popDownSprite, charBoxOpts, screen)
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
		sys.renderPopDownAnimation(popDownSprite, charBoxOpts, screen)
		return
	}

	// Handle the pop up animation.
	if configAndState.State.PopUpMode {
		finished := sys.renderPopUpAnimation(charBoxOpts, screen)

		if finished {
			// Handle the portrait.
			portraitSprites := components.SpritesSliceComponent.Get(chatEntity)
			sys.handlePortrait(configAndState, portraitSprites, screen)

			// Handle the text.
			slideContent := configAndState.State.SlidesContent[configAndState.State.CurrentSlideIndex]

			if configAndState.State.TextAnimStartTick == -1 {
				configAndState.State.TextAnimStartTick = sys.scene.Manager.TickHandler.CurrentTick()
			}

			if configAndState.State.NameTextAnimStartTick == -1 {
				configAndState.State.NameTextAnimStartTick = sys.scene.Manager.TickHandler.CurrentTick()
			}

			paddingX := 72.0
			paddingY := 35.0

			// Text.
			textUtil.RenderTextDefault(
				slideContent.Text,
				boxX+paddingX,
				boxY+paddingY,
				490,
				configAndState.State.TextAnimStartTick,
				configAndState.Config.TicksPerWord,
				1,
				&sys.scene.World,
				sys.scene.Manager.TickHandler,
				screen,
			)

			nameX := 88.0
			nameY := 345.0
			nameText := configAndState.State.SlidesContent[configAndState.State.CurrentSlideIndex].CharName

			// Name.
			textUtil.RenderTextDefault(
				nameText,
				nameX,
				nameY,
				210,
				configAndState.State.NameTextAnimStartTick,
				configAndState.Config.TicksPerWord,
				1,
				&sys.scene.World,
				sys.scene.Manager.TickHandler,
				screen,
			)

			nameWidth := float64(len(nameText) * 7)
			nameLineSX := float32(nameX - 6)
			nameLineEndX := float32(nameX+nameWidth) + 2
			nameLineY := float32(nameY - 5)

			cyan := color.RGBA{0, 254, 254, 255}

			// Horizontal line.
			vector.StrokeLine(screen, nameLineSX, nameLineY, nameLineEndX, nameLineY, 2, cyan, false)

			// Vertical line.
			vector.StrokeLine(screen, nameLineEndX, nameLineY, nameLineEndX, nameLineY+18, 2, cyan, false)
		}

		return
	}

}

func (sys ChatSlidesRendererSystem) renderPopUpAnimation(chatBoxOpts *ebiten.DrawImageOptions, screen *ebiten.Image) (finished bool) {
	popUpSprite := systemsUtil.GetChatPopUpSprite(sys.scene.World)

	spriteAtFrameIndex := animUtil.GetAnimFrame(sys.scene.Manager.TickHandler, popUpSprite)

	screen.DrawImage(spriteAtFrameIndex, chatBoxOpts)

	finished = sys.scene.Manager.TickHandler.TicksSinceNTicks(popUpSprite.StartTick) > UIGlobals.CHAT_BOX_ANIM_SPEED*2

	return finished
}

func (sys ChatSlidesRendererSystem) renderPopDownAnimation(popDownSprite *components.Sprite, chatBoxOpts *ebiten.DrawImageOptions, screen *ebiten.Image) {

	spriteAtFrameIndex := animUtil.GetAnimFrame(sys.scene.Manager.TickHandler, popDownSprite)

	screen.DrawImage(spriteAtFrameIndex, chatBoxOpts)

}

func (sys ChatSlidesRendererSystem) handlePortrait(configAndState *components.ChatStateAndConfig, portraitSprites *[]*components.Sprite, screen *ebiten.Image) {
	portraitOpts := &ebiten.DrawImageOptions{}
	index := configAndState.State.CurrentSlideIndex
	image := (*portraitSprites)[index].Image
	imgWidth := image.Bounds().Size()

	x := 80.0
	y := 263.0

	if !configAndState.State.SlidesContent[index].FacingRight {
		portraitOpts.GeoM.Scale(-1, 1)
		portraitOpts.GeoM.Translate(float64(imgWidth.X)+x, y)
	} else {
		portraitOpts.GeoM.Translate(x, y)
	}

	screen.DrawImage(image, portraitOpts)
}
