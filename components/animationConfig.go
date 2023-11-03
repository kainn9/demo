package components

type AnimationConfig struct {
	FrameWidth,
	FrameHeight,
	FrameCount,
	AnimationFramesPerTick int
	Freeze bool

	// This is technically "state", but it's used to track the animation.
	// It is the only that is allowed to be mutated in the render phase,
	// to begin/start the animation.
	// -1 is the default value, and means the animation is not active.
	StartTick int
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
		StartTick:              -1,
	}
}
