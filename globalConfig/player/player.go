package playerGlobals

import (
	"github.com/kainn9/demo/components"
	sharedAnimationConstants "github.com/kainn9/demo/globalConfig/sharedAnimation"
)

// RIGID BODY.
const (
	PLAYER_WIDTH  = 15
	PLAYER_HEIGHT = 75
)

// Sprite Offsets.
const (
	PLAYER_SPRITE_OFFSET_X = -95
	PLAYER_SPRITE_OFFSET_Y = -133
)

// Combat.
const (
	PLAYER_HURT_DURATION_TICKS   = 20
	PLAYER_DEFEAT_DURATION_TICKS = 200
)

// ANIMATIONS.
const (
	PLAYER_ANIMATIONS_SPRITE_WIDTH  = 192
	PLAYER_ANIMATIONS_SPRITE_HEIGHT = 168

	PLAYER_IDLE_ANIM_SPEED           = 20
	PLAYER_WALK_ANIM_SPEED           = 15
	PLAYER_RUN_ANIM_SPEED            = 12
	PLAYER_JUMP_ANIM_SPEED           = 8
	PLAYER_FALL_ANIM_SPEED           = 12
	PLAYER_CLIMB_LADDER_ANIM_SPEED   = 12
	PLAYER_ATTACK_PRIMARY_ANIM_SPEED = 6
	PLAYER_HURT_ANIM_SPEED           = 60
	PLAYER_DEFEATED_ANIM_SPEED       = 12

	PLAYER_IDLE_FRAME_COUNT                = 7
	PLAYER_WALK_FRAME_COUNT                = 8
	PLAYER_RUN_FRAME_COUNT                 = 8
	PLAYER_JUMP_FRAME_COUNT                = 5
	PLAYER_FALL_FRAME_COUNT                = 3
	PLAYER_CLIMB_LADDER_IDLE_FRAME_COUNT   = 7
	PLAYER_CLIMB_LADDER_ACTIVE_FRAME_COUNT = 8
	PLAYER_ATTACK_PRIMARY_FRAME_COUNT      = 8
	PLAYER_HURT_FRAME_COUNT                = 1
	PLAYER_DEFEATED_FRAME_COUNT            = 9

	PLAYER_CHAR_STATE_CLIMB_LADDER_IDLE   components.CharState = "climbLadderIdle"
	PLAYER_CHAR_STATE_CLIMB_LADDER_ACTIVE components.CharState = "climbLadderActive"
)

var PLAYER_ANIMATION_CONFIGS = map[components.CharState]*components.AnimationConfig{

	sharedAnimationConstants.CHAR_STATE_IDLE: components.NewAnimationConfig(
		PLAYER_ANIMATIONS_SPRITE_WIDTH,
		PLAYER_ANIMATIONS_SPRITE_HEIGHT,
		PLAYER_IDLE_FRAME_COUNT,
		PLAYER_IDLE_ANIM_SPEED,
		false,
	),

	sharedAnimationConstants.CHAR_STATE_WALK: components.NewAnimationConfig(
		PLAYER_ANIMATIONS_SPRITE_WIDTH,
		PLAYER_ANIMATIONS_SPRITE_HEIGHT,
		PLAYER_WALK_FRAME_COUNT,
		PLAYER_WALK_ANIM_SPEED,
		false,
	),

	sharedAnimationConstants.CHAR_STATE_RUN: components.NewAnimationConfig(
		PLAYER_ANIMATIONS_SPRITE_WIDTH,
		PLAYER_ANIMATIONS_SPRITE_HEIGHT,
		PLAYER_RUN_FRAME_COUNT,
		PLAYER_RUN_ANIM_SPEED,
		false,
	),

	sharedAnimationConstants.CHAR_STATE_JUMP: components.NewAnimationConfig(
		PLAYER_ANIMATIONS_SPRITE_WIDTH,
		PLAYER_ANIMATIONS_SPRITE_HEIGHT,
		PLAYER_JUMP_FRAME_COUNT,
		PLAYER_JUMP_ANIM_SPEED,
		true,
	),

	sharedAnimationConstants.CHAR_STATE_FALL: components.NewAnimationConfig(
		PLAYER_ANIMATIONS_SPRITE_WIDTH,
		PLAYER_ANIMATIONS_SPRITE_HEIGHT,
		PLAYER_FALL_FRAME_COUNT,
		PLAYER_FALL_ANIM_SPEED,
		true,
	),

	PLAYER_CHAR_STATE_CLIMB_LADDER_IDLE: components.NewAnimationConfig(
		PLAYER_ANIMATIONS_SPRITE_WIDTH,
		PLAYER_ANIMATIONS_SPRITE_HEIGHT,
		PLAYER_CLIMB_LADDER_IDLE_FRAME_COUNT,
		PLAYER_CLIMB_LADDER_ANIM_SPEED,
		false,
	),

	PLAYER_CHAR_STATE_CLIMB_LADDER_ACTIVE: components.NewAnimationConfig(
		PLAYER_ANIMATIONS_SPRITE_WIDTH,
		PLAYER_ANIMATIONS_SPRITE_HEIGHT,
		PLAYER_CLIMB_LADDER_ACTIVE_FRAME_COUNT,
		PLAYER_CLIMB_LADDER_ANIM_SPEED,
		false,
	),

	sharedAnimationConstants.CHAR_STATE_ATTACK_PRIMARY: components.NewAnimationConfig(
		PLAYER_ANIMATIONS_SPRITE_WIDTH,
		PLAYER_ANIMATIONS_SPRITE_HEIGHT,
		PLAYER_ATTACK_PRIMARY_FRAME_COUNT,
		PLAYER_ATTACK_PRIMARY_ANIM_SPEED,
		false,
	),

	sharedAnimationConstants.CHAR_STATE_HURT: components.NewAnimationConfig(
		PLAYER_ANIMATIONS_SPRITE_WIDTH,
		PLAYER_ANIMATIONS_SPRITE_HEIGHT,
		PLAYER_HURT_FRAME_COUNT,
		PLAYER_HURT_ANIM_SPEED,
		false,
	),

	sharedAnimationConstants.CHAR_STATE_DEFEATED: components.NewAnimationConfig(
		PLAYER_ANIMATIONS_SPRITE_WIDTH,
		PLAYER_ANIMATIONS_SPRITE_HEIGHT,
		PLAYER_DEFEATED_FRAME_COUNT,
		PLAYER_DEFEATED_ANIM_SPEED,
		true,
	),
}

type PlayerAttackData struct {
	TotalTickLength int
	TicksPerFrame   int
}

var PlayerAttackDataMap = map[components.CharState]*PlayerAttackData{
	sharedAnimationConstants.CHAR_STATE_ATTACK_PRIMARY: {
		TotalTickLength: PLAYER_ATTACK_PRIMARY_FRAME_COUNT * PLAYER_ATTACK_PRIMARY_ANIM_SPEED,
		TicksPerFrame:   PLAYER_ATTACK_PRIMARY_ANIM_SPEED,
	},
}

const (
	PLAYER_PORTRAIT_INDEX = "player"
	PLAYER_NAME           = "The Goodrich"
)