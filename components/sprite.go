package components

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

var SpriteComponent = donburi.NewComponentType[Sprite]()
var SpritesMapComponent = donburi.NewComponentType[map[string]*Sprite]()
var SpritesSliceComponent = donburi.NewComponentType[[]*Sprite]()
var SpritesCharStateMapComponent = donburi.NewComponentType[map[CharState]*Sprite]()

type Sprite struct {
	OffSetX, OffSetY float64
	*ebiten.Image
	*AssetData

	// This is optional, only needed for sprites that are animated(aka spriteSheets).
	*AnimationConfig
}

type CharState string

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

func NewSprite(offX, offY float64) *Sprite {
	return &Sprite{
		OffSetX:   offX,
		OffSetY:   offY,
		Image:     &ebiten.Image{},
		AssetData: &AssetData{},
	}
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
