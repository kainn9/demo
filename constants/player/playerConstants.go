package playerConstants

import "github.com/kainn9/demo/components"

// RIGID BODY.
const (
	PLAYER_WIDTH  = 18
	PLAYER_HEIGHT = 55
)

// Sprite Offsets.
const (
	PLAYER_SPRITE_OFFSET_X = -70
	PLAYER_SPRITE_OFFSET_Y = -100
)

// ANIMATIONS.
const (
	PLAYER_IDLE_ANIM_SPEED         = 20
	PLAYER_RUN_ANIM_SPEED          = 12
	PLAYER_JUMP_ANIM_SPEED         = 8
	PLAYER_FALL_ANIM_SPEED         = 12
	PLAYER_CLIMB_LADDER_ANIM_SPEED = 12

	PLAYER_IDLE_FRAME_COUNT                = 7
	PLAYER_RUN_FRAME_COUNT                 = 8
	PLAYER_JUMP_FRAME_COUNT                = 5
	PLAYER_FALL_FRAME_COUNT                = 2
	PLAYER_CLIMB_LADDER_IDLE_FRAME_COUNT   = 7
	PLAYER_CLIMB_LADDER_ACTIVE_FRAME_COUNT = 8

	PLAYER_ANIMATIONS_SPRITE_WIDTH  = 144
	PLAYER_ANIMATIONS_SPRITE_HEIGHT = 126

	// ANIM STATES.
	PLAYER_ANIM_STATE_IDLE                components.AnimState = "idle"
	PLAYER_ANIM_STATE_RUN                 components.AnimState = "run"
	PLAYER_ANIM_STATE_JUMP                components.AnimState = "jump"
	PLAYER_ANIM_STATE_FALL                components.AnimState = "fall"
	PLAYER_ANIM_STATE_CLIMB_LADDER_IDLE   components.AnimState = "climbLadderIdle"
	PLAYER_ANIM_STATE_CLIMB_LADDER_ACTIVE components.AnimState = "climbLadderActive"
)

var PLAYER_ANIMATION_CONFIGS = map[components.AnimState]*components.AnimationConfig{
	PLAYER_ANIM_STATE_IDLE:                components.NewAnimationConfig(PLAYER_ANIMATIONS_SPRITE_WIDTH, PLAYER_ANIMATIONS_SPRITE_HEIGHT, PLAYER_IDLE_FRAME_COUNT, PLAYER_IDLE_ANIM_SPEED, false),
	PLAYER_ANIM_STATE_RUN:                 components.NewAnimationConfig(PLAYER_ANIMATIONS_SPRITE_WIDTH, PLAYER_ANIMATIONS_SPRITE_HEIGHT, PLAYER_RUN_FRAME_COUNT, PLAYER_RUN_ANIM_SPEED, false),
	PLAYER_ANIM_STATE_JUMP:                components.NewAnimationConfig(PLAYER_ANIMATIONS_SPRITE_WIDTH, PLAYER_ANIMATIONS_SPRITE_HEIGHT, PLAYER_JUMP_FRAME_COUNT, PLAYER_JUMP_ANIM_SPEED, true),
	PLAYER_ANIM_STATE_FALL:                components.NewAnimationConfig(PLAYER_ANIMATIONS_SPRITE_WIDTH, PLAYER_ANIMATIONS_SPRITE_HEIGHT, PLAYER_FALL_FRAME_COUNT, PLAYER_FALL_ANIM_SPEED, true),
	PLAYER_ANIM_STATE_CLIMB_LADDER_IDLE:   components.NewAnimationConfig(PLAYER_ANIMATIONS_SPRITE_WIDTH, PLAYER_ANIMATIONS_SPRITE_HEIGHT, PLAYER_CLIMB_LADDER_IDLE_FRAME_COUNT, PLAYER_CLIMB_LADDER_ANIM_SPEED, false),
	PLAYER_ANIM_STATE_CLIMB_LADDER_ACTIVE: components.NewAnimationConfig(PLAYER_ANIMATIONS_SPRITE_WIDTH, PLAYER_ANIMATIONS_SPRITE_HEIGHT, PLAYER_CLIMB_LADDER_ACTIVE_FRAME_COUNT, PLAYER_CLIMB_LADDER_ANIM_SPEED, false),
}
