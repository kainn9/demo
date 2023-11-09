package npcGlobals

import (
	"github.com/kainn9/demo/components"
	sharedAnimationGlobals "github.com/kainn9/demo/globalConfig/sharedAnimation"
	tBokiVec "github.com/kainn9/tteokbokki/math/vec"
)

const (
	NPC_NAME_BIG_BOI components.NpcName = "bigBoi"

	BIG_BOI_WIDTH  = 18
	BIG_BOI_HEIGHT = 55

	BIG_BOI_SPRITE_OFFSET_X = -45
	BIG_BOI_SPRITE_OFFSET_Y = -68

	BIG_BOI_ANIMATIONS_SPRITE_WIDTH  = 96
	BIG_BOI_ANIMATIONS_SPRITE_HEIGHT = 96

	BIG_BOI_IDLE_FRAME_COUNT = 4
	BIG_BOI_IDLE_ANIM_SPEED  = 20

	BIG_BOI_HURT_FRAME_COUNT = 1
	BIG_BOI_HURT_ANIM_SPEED  = 60

	BIG_BOI_DEFEATED_FRAME_COUNT = 4
	BIG_BOI_DEFEATED_ANIM_SPEED  = 12

	BIG_BOI_DEFEATED_DURATION = 240
)

// RigidBody dimensions
var NPC_DIMENSIONS = map[components.NpcName]tBokiVec.Vec2{
	NPC_NAME_BIG_BOI: {X: BIG_BOI_WIDTH, Y: BIG_BOI_HEIGHT},
}

// Sprite offsets
var NPC_SPRITE_OFFSETS = map[components.NpcName]tBokiVec.Vec2{
	NPC_NAME_BIG_BOI: {X: BIG_BOI_SPRITE_OFFSET_X, Y: BIG_BOI_SPRITE_OFFSET_Y},
}

// Anim Configs
var BIG_BOI_ANIMATION_CONFIGS = map[components.CharState]components.AnimationConfig{
	sharedAnimationGlobals.CHAR_STATE_IDLE: *components.NewAnimationConfig(
		BIG_BOI_ANIMATIONS_SPRITE_WIDTH,
		BIG_BOI_ANIMATIONS_SPRITE_HEIGHT,
		BIG_BOI_IDLE_FRAME_COUNT,
		BIG_BOI_IDLE_ANIM_SPEED,
		false,
	),

	sharedAnimationGlobals.CHAR_STATE_HURT: *components.NewAnimationConfig(
		BIG_BOI_ANIMATIONS_SPRITE_WIDTH,
		BIG_BOI_ANIMATIONS_SPRITE_HEIGHT,
		BIG_BOI_HURT_FRAME_COUNT,
		BIG_BOI_HURT_ANIM_SPEED,
		false,
	),

	sharedAnimationGlobals.CHAR_STATE_DEFEATED: *components.NewAnimationConfig(
		BIG_BOI_ANIMATIONS_SPRITE_WIDTH,
		BIG_BOI_ANIMATIONS_SPRITE_HEIGHT,
		BIG_BOI_DEFEATED_FRAME_COUNT,
		BIG_BOI_DEFEATED_ANIM_SPEED,
		true,
	),
}

var NPC_ANIMATION_CONFIGS = map[components.NpcName]map[components.CharState]components.AnimationConfig{
	NPC_NAME_BIG_BOI: BIG_BOI_ANIMATION_CONFIGS,
}

var NPC_DEFEATED_DURATIONS = map[components.NpcName]int{
	NPC_NAME_BIG_BOI: BIG_BOI_DEFEATED_DURATION,
}
