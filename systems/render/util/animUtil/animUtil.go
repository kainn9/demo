package animUtil

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kainn9/coldBrew"
	assetComponents "github.com/kainn9/demo/components/assets"
)

func ResetAnimationConfig(spriteSheet *assetComponents.Sprite) {
	if spriteSheet == nil {
		return
	}

	spriteSheet.AnimationConfig.StartTick = 0
}

func GetFrame(m *coldBrew.Manager, spriteSheet *assetComponents.Sprite) *ebiten.Image {
	currentTick := m.TickHandler.CurrentTick()

	animConfig := spriteSheet.AnimationConfig

	// Anim has just started playing.
	if animConfig.StartTick == 0 {

		animConfig.StartTick = currentTick
	}

	// Anim has been played before, but has finished.
	totalAnimationTicks := animConfig.FrameCount * animConfig.AnimationFramesPerTick
	ticksSinceAnimationStart := m.TickHandler.TicksSinceNTicks(animConfig.StartTick)

	// If animation has finished, and does not have freeze bool,
	// allow the animation to loop.
	var frameIndex int
	animationFinished := ticksSinceAnimationStart >= totalAnimationTicks

	if animationFinished && animConfig.Freeze {
		frameIndex = animConfig.FrameCount - 1
	} else {
		frameIndex = (ticksSinceAnimationStart / animConfig.AnimationFramesPerTick) % animConfig.FrameCount
	}

	sx, sy := (0)+frameIndex*(animConfig.FrameWidth), (0)

	rect := image.Rect(sx, sy, sx+(animConfig.FrameWidth), animConfig.FrameHeight)

	return spriteSheet.SubImage(rect).(*ebiten.Image)
}
