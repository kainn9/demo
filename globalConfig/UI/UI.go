package UIGlobals

import (
	"github.com/kainn9/demo/components"
	tBokiVec "github.com/kainn9/tteokbokki/math/vec"
)

// Layout Constants.

type Layout = string

const (
	LAYOUT_DEFAULT  Layout = "" // Mouse and Keyboard.
	LAYOUT_NINTENDO Layout = "nintendoSwitch"
	LAYOUT_XBOX     Layout = "xboxOne"
	LAYOUT_PS4      Layout = "ps4"
)

const CurrentLayout = LAYOUT_DEFAULT

// Chat Box Constants.
const (
	CHAT_BOX_POP_UP_SPRITE_NAME   = "chatPopUp"
	CHAT_BOX_POP_DOWN_SPRITE_NAME = "chatPopDown"
	CHAT_BOX_ANIM_SPEED           = 12
)

var (
	CHAT_BOX_FRAME_WIDTH = 480
)

const (
	INDICATOR_JUMP     components.IndicatorType = "jump"
	INDICATOR_LADDER   components.IndicatorType = "ladder"
	INDICATOR_MOVEMENT components.IndicatorType = "movement"
	INDICATOR_INTERACT components.IndicatorType = "interact"
	INDICATOR_DESCEND  components.IndicatorType = "descend"
)

type IndicatorAnimationMap = map[components.IndicatorType]*components.AnimationConfig

var IndicatorAnimationConfigs = map[Layout]IndicatorAnimationMap{
	LAYOUT_DEFAULT: {
		INDICATOR_JUMP:     components.NewAnimationConfig(33, 16, 4, 24, false),
		INDICATOR_LADDER:   components.NewAnimationConfig(16, 34, 7, 24, false),
		INDICATOR_MOVEMENT: components.NewAnimationConfig(50, 33, 4, 24, false),
		INDICATOR_INTERACT: components.NewAnimationConfig(24, 16, 4, 24, false),
		INDICATOR_DESCEND:  components.NewAnimationConfig(62, 16, 4, 24, false),
	},
}

type IndicatorPlayerOffsetMap = map[components.IndicatorType]tBokiVec.Vec2

var IndicatorPlayerOffsets = map[Layout]IndicatorPlayerOffsetMap{
	LAYOUT_DEFAULT: {
		INDICATOR_JUMP:     {X: -15, Y: -73},
		INDICATOR_LADDER:   {X: -8, Y: -93},
		INDICATOR_MOVEMENT: {X: -24, Y: -93},
		INDICATOR_INTERACT: {X: -10, Y: -72},
		INDICATOR_DESCEND:  {X: 0, Y: 0},
	},
}
