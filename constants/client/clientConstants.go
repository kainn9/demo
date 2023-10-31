package client

const (
	SCREEN_HEIGHT     = 360
	SCREEN_WIDTH      = 640
	MAX_TICKS         = 1200000
	SCENE_CACHE_LIMIT = 5

	IMAGE_EXTENSION = ".png"
	ASSET_ROOT_PATH = "./assets/"

	SCENE_ASSETS_SUB_PATH          = "scenes/"                           // ./assets/scenes/
	UI_ASSETS_SUB_PATH             = "UI/"                               // ./assets/UI/
	UI_ASSETS_CHAT_BOX_SUBPATH     = UI_ASSETS_SUB_PATH + "chatBox/"     // ./assets/UI/chatBox/
	UI_ASSETS_FONT_SUBPATH         = UI_ASSETS_SUB_PATH + "font/"        // ./assets/UI/font/
	UI_ASSETS_INDICATORS_SUBPATH   = UI_ASSETS_SUB_PATH + "indicators/"  // ./assets/UI/indicators/
	UI_ASSETS_DEFAULT_FONT_SUBPATH = UI_ASSETS_FONT_SUBPATH + "default/" // ./assets/UI/font/default/

	CHARACTER_ASSETS_SUB_PATH  = "characters/"                         // ./assets/characters/
	PLAYER_ASSETS_SUB_PATH     = CHARACTER_ASSETS_SUB_PATH + "player/" // ./assets/characters/player/
	CHAT_SCENE_ASSETS_SUB_PATH = "chats/"                              // ./assets/scenes/SCENE_SECTION/SCENE_NAME/chat/
)
