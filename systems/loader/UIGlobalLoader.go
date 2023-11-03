package loaderSystems

import (
	"log"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	UIConstants "github.com/kainn9/demo/constants/UI"
	clientConstants "github.com/kainn9/demo/constants/client"
	fontConstants "github.com/kainn9/demo/constants/font"
	"github.com/kainn9/demo/queries"
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

func (sys UIGlobalLoaderSystem) CustomQuery() *donburi.Query {
	return queries.UISingletonQuery
}

func (sys UIGlobalLoaderSystem) Load(_ *donburi.Entry) {

	query := sys.CustomQuery()
	world := sys.scene.World

	UISpritesMapEntity, _ := query.First(world)

	UISpritesMap := components.SpritesMapComponent.Get(UISpritesMapEntity)

	// Todo: Break this up into sub functions.
	sys.loadAll(UISpritesMap)

}

func (sys UIGlobalLoaderSystem) loadAll(UISpritesMap *map[string]*components.Sprite) {

	// Default Font.
	lowerCaseSpriteSheet := (*UISpritesMap)[fontConstants.FONT_DEFAULT_NAME+fontConstants.FONT_LOWER_CASE_SPRITE_NAME]

	// Assuming if lowerCaseSpriteSheet(first global asset) is loaded,
	// then all global assets are loaded(for now, at least).
	if lowerCaseSpriteSheet.Loaded {
		return
	}

	// --------------------------------------------------------------------------------
	// Default Font.

	log.Println("Loading default font.")

	upperCaseSpriteSheet := (*UISpritesMap)[fontConstants.FONT_DEFAULT_NAME+fontConstants.FONT_UPPER_CASE_SPRITE_NAME]
	numbersSpriteSheet := (*UISpritesMap)[fontConstants.FONT_DEFAULT_NAME+fontConstants.FONT_NUMBERS_SPRITE_NAME]
	specialSpriteSheet := (*UISpritesMap)[fontConstants.FONT_DEFAULT_NAME+fontConstants.FONT_SPECIAL_SPRITE_NAME]

	path := clientConstants.UI_ASSETS_DEFAULT_FONT_SUBPATH

	log.Println("Loading", path+fontConstants.FONT_LOWER_CASE_SPRITE_NAME)
	LoadImage(path+fontConstants.FONT_LOWER_CASE_SPRITE_NAME, lowerCaseSpriteSheet)

	log.Println("Loading", path+fontConstants.FONT_UPPER_CASE_SPRITE_NAME)
	LoadImage(path+fontConstants.FONT_UPPER_CASE_SPRITE_NAME, upperCaseSpriteSheet)

	log.Println("Loading", path+fontConstants.FONT_NUMBERS_SPRITE_NAME)
	LoadImage(path+fontConstants.FONT_NUMBERS_SPRITE_NAME, numbersSpriteSheet)

	log.Println("Loading", path+fontConstants.FONT_SPECIAL_SPRITE_NAME)
	LoadImage(path+fontConstants.FONT_SPECIAL_SPRITE_NAME, specialSpriteSheet)

	// --------------------------------------------------------------------------------
	// Chat Box Sprites.
	log.Println("Loading chat box sprites.")
	chatBoxPopUpSprite := (*UISpritesMap)[UIConstants.CHAT_BOX_POP_UP_SPRITE_NAME]
	chatBoxPopDownSprite := (*UISpritesMap)[UIConstants.CHAT_BOX_POP_DOWN_SPRITE_NAME]

	spriteNames := []string{
		UIConstants.CHAT_BOX_POP_UP_SPRITE_NAME,
		UIConstants.CHAT_BOX_POP_DOWN_SPRITE_NAME,
	}

	sprites := []*components.Sprite{
		chatBoxPopUpSprite,
		chatBoxPopDownSprite,
	}

	for i, sprite := range sprites {

		if sprite.Loaded {
			continue
		}

		path = clientConstants.UI_ASSETS_CHAT_BOX_SUBPATH

		log.Println("Loading", path+spriteNames[i])
		LoadImage(path+spriteNames[i], sprite)

		frameWidth, frameHeight, frameCount := sys.getAnimData(sprite)
		sprite.AnimationConfig = components.NewAnimationConfig(frameWidth, frameHeight, frameCount, UIConstants.CHAT_BOX_ANIM_SPEED, true)
	}

	// --------------------------------------------------------------------------------
	// Indicator Sprites.
	log.Println("Loading indicator sprites.")

	descKey := string(UIConstants.CurrentLayout) + string(UIConstants.INDICATOR_DESCEND)
	interactKey := string(UIConstants.CurrentLayout) + string(UIConstants.INDICATOR_INTERACT)
	jumpKey := string(UIConstants.CurrentLayout) + string(UIConstants.INDICATOR_JUMP)
	ladderKey := string(UIConstants.CurrentLayout) + string(UIConstants.INDICATOR_LADDER)
	movementKey := string(UIConstants.CurrentLayout) + string(UIConstants.INDICATOR_MOVEMENT)

	indicatorDescendSprite := (*UISpritesMap)[descKey]
	indicatorInteractSprite := (*UISpritesMap)[interactKey]
	indicatorJumpSprite := (*UISpritesMap)[string(jumpKey)]
	indicatorLadderSprite := (*UISpritesMap)[string(ladderKey)]
	indicatorMovementSprite := (*UISpritesMap)[string(movementKey)]

	path = clientConstants.UI_ASSETS_INDICATORS_SUBPATH
	path += UIConstants.CurrentLayout + ""

	log.Println("Loading", path+descKey)
	LoadImage(path+descKey, indicatorDescendSprite)

	log.Println("Loading", path+interactKey)
	LoadImage(path+interactKey, indicatorInteractSprite)

	log.Println("Loading", path+jumpKey)
	LoadImage(path+jumpKey, indicatorJumpSprite)

	log.Println("Loading", path+ladderKey)
	LoadImage(path+ladderKey, indicatorLadderSprite)

	log.Println("Loading", path+movementKey)
	LoadImage(path+movementKey, indicatorMovementSprite)

	// Register Animation Configs.
	indicatorDescendSprite.AnimationConfig = UIConstants.IndicatorAnimationConfigs[UIConstants.CurrentLayout][UIConstants.INDICATOR_DESCEND]
	indicatorInteractSprite.AnimationConfig = UIConstants.IndicatorAnimationConfigs[UIConstants.CurrentLayout][UIConstants.INDICATOR_INTERACT]
	indicatorJumpSprite.AnimationConfig = UIConstants.IndicatorAnimationConfigs[UIConstants.CurrentLayout][UIConstants.INDICATOR_JUMP]
	indicatorLadderSprite.AnimationConfig = UIConstants.IndicatorAnimationConfigs[UIConstants.CurrentLayout][UIConstants.INDICATOR_LADDER]
	indicatorMovementSprite.AnimationConfig = UIConstants.IndicatorAnimationConfigs[UIConstants.CurrentLayout][UIConstants.INDICATOR_MOVEMENT]

	// --------------------------------------------------------------------------------

}

func (sys UIGlobalLoaderSystem) getAnimData(spriteComponent *components.Sprite) (frameWidth, frameHeight, frameCount int) {
	totalFrameWidth := spriteComponent.Image.Bounds().Size().X

	frameWidth = UIConstants.CHAT_BOX_FRAME_WIDTH
	frameHeight = spriteComponent.Image.Bounds().Size().Y
	frameCount = totalFrameWidth / UIConstants.CHAT_BOX_FRAME_WIDTH

	return frameWidth, frameHeight, frameCount
}
