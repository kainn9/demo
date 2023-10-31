package playerConstants

// RIGID BODY.
const (
	PLAYER_WIDTH           = 18
	PLAYER_HEIGHT          = 55
	PLAYER_SPRITE_OFFSET_X = 70
	PLAYER_SPRITE_OFFSET_Y = 100
)

// ANIMATIONS.

type AnimState string

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
	PLAYER_ANIM_STATE_IDLE                AnimState = "idle"
	PLAYER_ANIM_STATE_RUN                 AnimState = "run"
	PLAYER_ANIM_STATE_JUMP                AnimState = "jump"
	PLAYER_ANIM_STATE_FALL                AnimState = "fall"
	PLAYER_ANIM_STATE_CLIMB_LADDER_IDLE   AnimState = "climbLadderIdle"
	PLAYER_ANIM_STATE_CLIMB_LADDER_ACTIVE AnimState = "climbLadderActive"
)
