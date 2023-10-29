package animUtil

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kainn9/coldBrew"
	assetComponents "github.com/kainn9/demo/components/assets"
)

func ResetAnimationData(spriteSheet *assetComponents.Sprite) {
	if spriteSheet == nil {
		return
	}

	spriteSheet.AnimationData.StartTick = 0
}

func GetFrame(m *coldBrew.Manager, spriteSheet *assetComponents.Sprite) *ebiten.Image {
	currentTick := m.TickHandler.CurrentTick()

	animData := spriteSheet.AnimationData

	// Anim has just started playing.
	if animData.StartTick == 0 {

		animData.StartTick = currentTick
	}

	// Anim has been played before, but has finished.
	totalAnimationTicks := animData.FrameCount * animData.AnimationFramesPerTick
	ticksSinceAnimationStart := m.TickHandler.TicksSinceNTicks(animData.StartTick)

	// If animation has finished, and does not have freeze bool,
	// allow the animation to loop.
	var frameIndex int
	animationFinished := ticksSinceAnimationStart >= totalAnimationTicks

	if animationFinished && animData.Freeze {
		frameIndex = animData.FrameCount - 1
	} else {
		frameIndex = (ticksSinceAnimationStart / animData.AnimationFramesPerTick) % animData.FrameCount
	}

	sx, sy := (0)+frameIndex*(animData.FrameWidth), (0)

	rect := image.Rect(sx, sy, sx+(animData.FrameWidth), animData.FrameHeight)

	return spriteSheet.SubImage(rect).(*ebiten.Image)
}
