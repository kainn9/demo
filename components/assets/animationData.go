package assetComponents

type AnimationData struct {
	FrameWidth,
	FrameHeight,
	FrameCount,
	AnimationFramesPerTick,
	StartTick int
	Freeze bool
}

func NewAnimationData(
	frameWidth,
	frameHeight,
	frameCount,
	animationFramesPerTick int,
	freeze bool,

) *AnimationData {

	return &AnimationData{
		FrameWidth:             frameWidth,
		FrameHeight:            frameHeight,
		FrameCount:             frameCount,
		AnimationFramesPerTick: animationFramesPerTick,
		Freeze:                 freeze,
	}
}
