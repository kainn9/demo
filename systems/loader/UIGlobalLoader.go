package loaderSystems

import (
	"log"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	UIGlobals "github.com/kainn9/demo/globalConfig/UI"
	clientGlobals "github.com/kainn9/demo/globalConfig/client"
	fontGlobals "github.com/kainn9/demo/globalConfig/font"
	"github.com/kainn9/demo/queries"
	loaderUtil "github.com/kainn9/demo/systems/loader/util"
	"github.com/yohamta/donburi"
)

// Todo: add some kind of sorting based on ZIndex.
type UIGlobalLoaderSystem struct {
	scene *coldBrew.Scene
}

func NewUIGlobalLoader(scene *coldBrew.Scene) *UIGlobalLoaderSystem {
	return &UIGlobalLoaderSystem{
		scene: scene,
	}
}

func (sys UIGlobalLoaderSystem) SpritesSingletonQuery() *donburi.Query {
	return queries.UISingletonSpritesQuery
}
func (sys UIGlobalLoaderSystem) SoundsSingletonQuery() *donburi.Query {
	return queries.UISingletonSoundsQuery
}

func (sys UIGlobalLoaderSystem) Load(_ *donburi.Entry) {

	spritesSingletonQuery := sys.SpritesSingletonQuery()
	soundsSingletonQuery := sys.SoundsSingletonQuery()

	world := sys.scene.World

	UISpritesMapEntity, _ := spritesSingletonQuery.First(world)
	UISoundsMapEntity, _ := soundsSingletonQuery.First(world)

	UISpritesMap := components.SpritesMapComponent.Get(UISpritesMapEntity)
	UISoundsMap := components.SoundsMapComponent.Get(UISoundsMapEntity)

	// Todo: Break this up into sub functions.
	sys.loadAllSprites(UISpritesMap)
	sys.loadAllSounds(UISoundsMap)

}

func (sys UIGlobalLoaderSystem) loadAllSprites(UISpritesMap *map[string]*components.Sprite) {

	// Default Font.
	lowerCaseSpriteSheet := (*UISpritesMap)[fontGlobals.FONT_DEFAULT_NAME+fontGlobals.FONT_LOWER_CASE_SPRITE_NAME]

	// Assuming if lowerCaseSpriteSheet(first global asset) is loaded,
	// then all global assets are loaded(for now, at least).
	if lowerCaseSpriteSheet.Loaded {
		return
	}

	// --------------------------------------------------------------------------------
	// Default Font.

	log.Println("Loading default font.")

	upperCaseSpriteSheet := (*UISpritesMap)[fontGlobals.FONT_DEFAULT_NAME+fontGlobals.FONT_UPPER_CASE_SPRITE_NAME]
	numbersSpriteSheet := (*UISpritesMap)[fontGlobals.FONT_DEFAULT_NAME+fontGlobals.FONT_NUMBERS_SPRITE_NAME]
	specialSpriteSheet := (*UISpritesMap)[fontGlobals.FONT_DEFAULT_NAME+fontGlobals.FONT_SPECIAL_SPRITE_NAME]

	path := clientGlobals.UI_ASSETS_DEFAULT_FONT_SUBPATH

	log.Println("Loading", path+fontGlobals.FONT_LOWER_CASE_SPRITE_NAME)
	loaderUtil.LoadImage(path+fontGlobals.FONT_LOWER_CASE_SPRITE_NAME, lowerCaseSpriteSheet)

	log.Println("Loading", path+fontGlobals.FONT_UPPER_CASE_SPRITE_NAME)
	loaderUtil.LoadImage(path+fontGlobals.FONT_UPPER_CASE_SPRITE_NAME, upperCaseSpriteSheet)

	log.Println("Loading", path+fontGlobals.FONT_NUMBERS_SPRITE_NAME)
	loaderUtil.LoadImage(path+fontGlobals.FONT_NUMBERS_SPRITE_NAME, numbersSpriteSheet)

	log.Println("Loading", path+fontGlobals.FONT_SPECIAL_SPRITE_NAME)
	loaderUtil.LoadImage(path+fontGlobals.FONT_SPECIAL_SPRITE_NAME, specialSpriteSheet)

	// --------------------------------------------------------------------------------
	// Chat Box Sprites.
	log.Println("Loading chat box sprites.")
	chatBoxPopUpSprite := (*UISpritesMap)[UIGlobals.CHAT_BOX_POP_UP_SPRITE_NAME]
	chatBoxPopDownSprite := (*UISpritesMap)[UIGlobals.CHAT_BOX_POP_DOWN_SPRITE_NAME]

	spriteNames := []string{
		UIGlobals.CHAT_BOX_POP_UP_SPRITE_NAME,
		UIGlobals.CHAT_BOX_POP_DOWN_SPRITE_NAME,
	}

	sprites := []*components.Sprite{
		chatBoxPopUpSprite,
		chatBoxPopDownSprite,
	}

	for i, sprite := range sprites {

		if sprite.Loaded {
			continue
		}

		path = clientGlobals.UI_ASSETS_CHAT_BOX_SUBPATH

		log.Println("Loading", path+spriteNames[i])
		loaderUtil.LoadImage(path+spriteNames[i], sprite)

		frameWidth, frameHeight, frameCount := sys.getAnimData(sprite)
		sprite.AnimationConfig = components.NewAnimationConfig(frameWidth, frameHeight, frameCount, UIGlobals.CHAT_BOX_ANIM_SPEED, true)
	}

	// --------------------------------------------------------------------------------
	// Indicator Sprites.
	log.Println("Loading indicator sprites.")

	descKey := string(UIGlobals.CurrentLayout) + string(UIGlobals.INDICATOR_DESCEND)
	interactKey := string(UIGlobals.CurrentLayout) + string(UIGlobals.INDICATOR_INTERACT)
	jumpKey := string(UIGlobals.CurrentLayout) + string(UIGlobals.INDICATOR_JUMP)
	ladderKey := string(UIGlobals.CurrentLayout) + string(UIGlobals.INDICATOR_LADDER)
	movementKey := string(UIGlobals.CurrentLayout) + string(UIGlobals.INDICATOR_MOVEMENT)

	indicatorDescendSprite := (*UISpritesMap)[descKey]
	indicatorInteractSprite := (*UISpritesMap)[interactKey]
	indicatorJumpSprite := (*UISpritesMap)[string(jumpKey)]
	indicatorLadderSprite := (*UISpritesMap)[string(ladderKey)]
	indicatorMovementSprite := (*UISpritesMap)[string(movementKey)]

	path = clientGlobals.UI_ASSETS_INDICATORS_SUBPATH
	path += UIGlobals.CurrentLayout + ""

	log.Println("Loading", path+descKey)
	loaderUtil.LoadImage(path+descKey, indicatorDescendSprite)

	log.Println("Loading", path+interactKey)
	loaderUtil.LoadImage(path+interactKey, indicatorInteractSprite)

	log.Println("Loading", path+jumpKey)
	loaderUtil.LoadImage(path+jumpKey, indicatorJumpSprite)

	log.Println("Loading", path+ladderKey)
	loaderUtil.LoadImage(path+ladderKey, indicatorLadderSprite)

	log.Println("Loading", path+movementKey)
	loaderUtil.LoadImage(path+movementKey, indicatorMovementSprite)

	// Register Animation Configs.
	indicatorDescendSprite.AnimationConfig = UIGlobals.IndicatorAnimationConfigs[UIGlobals.CurrentLayout][UIGlobals.INDICATOR_DESCEND]
	indicatorInteractSprite.AnimationConfig = UIGlobals.IndicatorAnimationConfigs[UIGlobals.CurrentLayout][UIGlobals.INDICATOR_INTERACT]
	indicatorJumpSprite.AnimationConfig = UIGlobals.IndicatorAnimationConfigs[UIGlobals.CurrentLayout][UIGlobals.INDICATOR_JUMP]
	indicatorLadderSprite.AnimationConfig = UIGlobals.IndicatorAnimationConfigs[UIGlobals.CurrentLayout][UIGlobals.INDICATOR_LADDER]
	indicatorMovementSprite.AnimationConfig = UIGlobals.IndicatorAnimationConfigs[UIGlobals.CurrentLayout][UIGlobals.INDICATOR_MOVEMENT]

	// --------------------------------------------------------------------------------

}

func (sys UIGlobalLoaderSystem) getAnimData(spriteComponent *components.Sprite) (frameWidth, frameHeight, frameCount int) {
	totalFrameWidth := spriteComponent.Image.Bounds().Size().X

	frameWidth = UIGlobals.CHAT_BOX_FRAME_WIDTH
	frameHeight = spriteComponent.Image.Bounds().Size().Y
	frameCount = totalFrameWidth / UIGlobals.CHAT_BOX_FRAME_WIDTH

	return frameWidth, frameHeight, frameCount
}

func (sys UIGlobalLoaderSystem) loadAllSounds(UISoundsMap *map[string]*components.Sound) {

	// Assuming if chatActiveSound(first global asset) is loaded,
	// then all global assets are loaded(for now, at least).
	if (*UISoundsMap)[UIGlobals.CHAT_BOX_NEW_SOUND_NAME].Loaded {
		return
	}

	log.Println("Loading chat box sounds.")

	log.Println("Loading", clientGlobals.UI_ASSETS_CHAT_BOX_SUBPATH+UIGlobals.CHAT_BOX_NEW_SOUND_NAME)
	chatNewSound := (*UISoundsMap)[UIGlobals.CHAT_BOX_NEW_SOUND_NAME]
	loaderUtil.LoadSound(clientGlobals.UI_ASSETS_CHAT_BOX_SUBPATH+UIGlobals.CHAT_BOX_NEW_SOUND_NAME, chatNewSound)
}
