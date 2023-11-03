package animUtil

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
)

func ResetAnimationConfig(spriteSheet *components.Sprite) {
	if spriteSheet == nil {
		return
	}

	spriteSheet.AnimationConfig.StartTick = -1
}

func ActiveAnimation(animConfig *components.AnimationConfig) bool {
	return animConfig.StartTick != -1
}

func InactiveAnimation(spriteSheet *components.AnimationConfig) bool {
	return !ActiveAnimation(spriteSheet)
}

// PlayAnim returns the current frame of the animation.
// If the Animation has not been "initialized", it will initialize it,
// by setting the start tick to the current tick.
func PlayAnim(m *coldBrew.Manager, spriteSheet *components.Sprite) *ebiten.Image {

	currentTick := m.TickHandler.CurrentTick()
	animConfig := spriteSheet.AnimationConfig

	// Anim has just started playing.
	if InactiveAnimation(animConfig) {
		animConfig.StartTick = currentTick
	}

	// Anim has been played before, but has finished.
	totalAnimationTicks := animConfig.FrameCount * animConfig.AnimationFramesPerTick
	ticksSinceAnimationStart := m.TickHandler.TicksSinceNTicks(animConfig.StartTick)

	// If animation has finished, and does not have freeze bool,
	// allow the animation to loop. Otherwise, freeze the animation
	// (render the last frame).
	var frameIndex int
	animationFinished := ticksSinceAnimationStart >= totalAnimationTicks

	if animationFinished && animConfig.Freeze {
		frameIndex = animConfig.FrameCount - 1
	} else {
		frameIndex = (ticksSinceAnimationStart / animConfig.AnimationFramesPerTick) % animConfig.FrameCount
	}

	sx, sy := frameIndex*(animConfig.FrameWidth), 0

	rect := image.Rect(sx, sy, sx+(animConfig.FrameWidth), animConfig.FrameHeight)

	return spriteSheet.SubImage(rect).(*ebiten.Image)
}
