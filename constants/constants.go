package constants

import "github.com/hajimehoshi/ebiten/v2"

// Client Constants.
const (
	SCREEN_HEIGHT     = 360
	SCREEN_WIDTH      = 640
	MAX_TICKS         = 1200000
	SCENE_CACHE_LIMIT = 5

	IMAGE_EXTENSION = ".png"
	ASSET_ROOT_PATH = "./assets/"

	SCENE_ASSETS_SUB_PATH     = "scenes/"     // ./assets/scenes/
	UI_ASSETS_SUB_PATH        = "UI/"         // ./assets/UI/
	CHARACTER_ASSETS_SUB_PATH = "characters/" // ./assets/characters/

	PLAYER_ASSETS_SUB_PATH     = CHARACTER_ASSETS_SUB_PATH + "player/" // ./assets/characters/player/
	CHAT_SCENE_ASSETS_SUB_PATH = "chats/"                              // ./assets/scenes/SCENE_SECTION/SCENE_NAME/chat/
)

// Player  Constants.

const (
	// RIGID BODY.
	PLAYER_WIDTH           = 18
	PLAYER_HEIGHT          = 55
	PLAYER_SPRITE_OFFSET_X = 70
	PLAYER_SPRITE_OFFSET_Y = 100

	// ANIMATIONS.
	PLAYER_IDLE_ANIM_SPEED = 20
	PLAYER_RUN_ANIM_SPEED  = 12
	PLAYER_JUMP_ANIM_SPEED = 8
	PLAYER_FALL_ANIM_SPEED = 12

	PLAYER_IDLE_FRAME_COUNT = 7
	PLAYER_RUN_FRAME_COUNT  = 8
	PLAYER_JUMP_FRAME_COUNT = 5
	PLAYER_FALL_FRAME_COUNT = 2

	PLAYER_ANIMATIONS_SPRITE_WIDTH  = 144
	PLAYER_ANIMATIONS_SPRITE_HEIGHT = 126

	// ANIM STATES.
	PLAYER_ANIM_STATE_IDLE = "idle"
	PLAYER_ANIM_STATE_RUN  = "run"
	PLAYER_ANIM_STATE_JUMP = "jump"
	PLAYER_ANIM_STATE_FALL = "fall"
)

// Camera Constants.
const (
	CAMERA_MIN_SPEED = 2.5
	CAMERA_MAX_SPEED = 3.5
	CAMERA_EPSILON   = SCREEN_WIDTH/2 - 270
)

// Chat Constants.
const (
	CHAT_FRAME_WIDTH       = SCREEN_WIDTH - 21
	CHAT_ANIM_POP_UP_SPEED = 12
	CHAT_ANIM_TEXT_SPEED   = 12
	CHAT_POP_UP_TICKS      = 2 * CHAT_ANIM_POP_UP_SPEED
	CHAT_MAP_SLIDES_KEY    = "slides"
	CHAT_MAP_PORTRAITS_KEY = "portraits"

	CHAT_STATE_POP_UP   = "chatPopUp"
	CHAT_STATE_POP_DOWN = "chatPopDown"
)

// Special Inputs.
const (
	NO_INPUT            = -1
	COMBO_DOWN_SPACE    = -2
	RELEASED_HORIZONTAL = -3
	RELEASED_VERTICAL   = -4
)

// Key Binds.
type KeyBind = string

const (
	KEY_BIND_LEFT     KeyBind = "left"
	KEY_BIND_RIGHT    KeyBind = "right"
	KEY_BIND_JUMP     KeyBind = "jump"
	KEY_BIND_DOWN     KeyBind = "down"
	KEY_BIND_UP       KeyBind = "up"
	KEY_BIND_INTERACT         = "interact"
)

var KEY_BINDS = map[KeyBind]ebiten.Key{
	KEY_BIND_LEFT:     ebiten.KeyA,
	KEY_BIND_RIGHT:    ebiten.KeyD,
	KEY_BIND_JUMP:     ebiten.KeySpace,
	KEY_BIND_UP:       ebiten.KeyW,
	KEY_BIND_DOWN:     ebiten.KeyS,
	KEY_BIND_INTERACT: ebiten.KeyEnter,
}

func AllBinds() (
	left ebiten.Key,
	right ebiten.Key,
	jump ebiten.Key,
	up ebiten.Key,
	down ebiten.Key,
	interact ebiten.Key,
) {
	return KEY_BINDS[KEY_BIND_LEFT],
		KEY_BINDS[KEY_BIND_RIGHT],
		KEY_BINDS[KEY_BIND_JUMP],
		KEY_BINDS[KEY_BIND_UP],
		KEY_BINDS[KEY_BIND_DOWN],
		KEY_BINDS[KEY_BIND_INTERACT]
}
