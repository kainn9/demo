package clientGlobals

var (
	SCREEN_HEIGHT = 360
	SCREEN_WIDTH  = 640
)

const (
	MAX_TICKS         = 1200000
	SCENE_CACHE_LIMIT = 5
	SOUND_SAMPLE_RATE = 32000
	SOUND_MAX_VOLUME  = 1

	IMAGE_EXTENSION = ".png"
	SOUND_EXTENSION = ".mp3"
	ASSET_ROOT_PATH = "./assets/"

	BG_SOUND_NAME = "bgSound"

	CHAT_ASSETS_PREFIX_PATH      = "chatBox/"
	FONT_ASSETS_PREFIX_PATH      = "font/"
	INDICATOR_ASSETS_PREFIX_PATH = "indicators/"
	CHARACTER_ASSETS_PREFIX_PATH = "characters/"
	SOUNDS_ASSETS_PREFIX_PATH    = "sounds/"
	SPRITES_ASSETS_PREFIX_PATH   = "sprites/"
	PLAYER_PREFIX_PATH           = "player/"

	SCENE_ASSETS_SUB_PATH          = "scenes/"                                         // ./assets/scenes/
	UI_ASSETS_SUB_PATH             = "UI/"                                             // ./assets/UI/
	UI_ASSETS_CHAT_BOX_SUBPATH     = UI_ASSETS_SUB_PATH + CHAT_ASSETS_PREFIX_PATH      // ./assets/UI/chatBox/
	UI_ASSETS_FONT_SUBPATH         = UI_ASSETS_SUB_PATH + FONT_ASSETS_PREFIX_PATH      // ./assets/UI/font/
	UI_ASSETS_INDICATORS_SUBPATH   = UI_ASSETS_SUB_PATH + INDICATOR_ASSETS_PREFIX_PATH // ./assets/UI/indicators/
	UI_ASSETS_DEFAULT_FONT_SUBPATH = UI_ASSETS_FONT_SUBPATH + "default/"               // ./assets/UI/font/default/

	CHAT_SCENE_ASSETS_SUB_PATH = "chats/" // ./assets/scenes/SCENE_SECTION/SCENE_NAME/chat/

	DEBUG_HITBOX_PREVIEW_JSON_PATH = "./debugConfigs/hitboxPreview.json"
	SCENE_BUILDER_CLICKER_RADIUS   = 10
)

var DEBUG_MODE = true
