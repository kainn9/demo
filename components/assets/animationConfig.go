package assetComponents

type AnimationConfig struct {
	FrameWidth,
	FrameHeight,
	FrameCount,
	AnimationFramesPerTick,
	StartTick int
	Freeze bool
}

func NewAnimationConfig(
	frameWidth,
	frameHeight,
	frameCount,
	animationFramesPerTick int,
	freeze bool,

) *AnimationConfig {

	return &AnimationConfig{
		FrameWidth:             frameWidth,
		FrameHeight:            frameHeight,
		FrameCount:             frameCount,
		AnimationFramesPerTick: animationFramesPerTick,
		Freeze:                 freeze,
	}
}
