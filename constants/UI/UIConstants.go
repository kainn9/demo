package UIConstants

import (
	clientConstants "github.com/kainn9/demo/constants/client"
)

// Chat Box Constants.
const (
	CHAT_BOX_FRAME_WIDTH          = clientConstants.SCREEN_WIDTH - 21
	CHAT_BOX_ANIM_SPEED           = 12
	CHAT_BOX_POP_UP_SPRITE_NAME   = "chatPopUp"
	CHAT_BOX_POP_DOWN_SPRITE_NAME = "chatPopDown"
)

// Indicator Constants.

type IndicatorType string

const (
	INDICATOR_JUMP     IndicatorType = "jump"
	INDICATOR_LADDER   IndicatorType = "ladder"
	INDICATOR_MOVEMENT IndicatorType = "movement"
	INDICATOR_INTERACT IndicatorType = "interact"
	INDICATOR_DESCEND  IndicatorType = "descend"
)
