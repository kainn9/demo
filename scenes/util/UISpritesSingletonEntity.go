package scenesUtil

import (
	"github.com/kainn9/coldBrew"
	assetComponents "github.com/kainn9/demo/components/assets"
	UIConstants "github.com/kainn9/demo/constants/UI"
	fontConstants "github.com/kainn9/demo/constants/font"

	"github.com/kainn9/demo/tags"
)

func AddUISpritesSingletonEntity(scene *coldBrew.Scene) {

	// Initialize UI Singleton SpritesMap.
	UISingletonSpritesMapEntity := scene.AddEntity(
		assetComponents.SpritesMapComponent,
		tags.UISingletonTag,
	)

	UISingletonSprites := make(map[string]*assetComponents.Sprite)

	// --------------------------------------------------------------------------------
	// Default Font
	keyDefault := fontConstants.FONT_DEFAULT_NAME
	keyLower := keyDefault + fontConstants.FONT_LOWER_CASE_SPRITE_NAME
	keyUpper := keyDefault + fontConstants.FONT_UPPER_CASE_SPRITE_NAME
	keyNumbers := keyDefault + fontConstants.FONT_NUMBERS_SPRITE_NAME
	keySpecial := keyDefault + fontConstants.FONT_SPECIAL_SPRITE_NAME

	UISingletonSprites[keyLower] = assetComponents.NewSprite(0, 0)
	UISingletonSprites[keyUpper] = assetComponents.NewSprite(0, 0)
	UISingletonSprites[keyNumbers] = assetComponents.NewSprite(0, 0)
	UISingletonSprites[keySpecial] = assetComponents.NewSprite(0, 0)

	// --------------------------------------------------------------------------------
	// Chat Box Sprites.
	UISingletonSprites[UIConstants.CHAT_BOX_POP_UP_SPRITE_NAME] = assetComponents.NewSprite(0, 0)
	UISingletonSprites[UIConstants.CHAT_BOX_POP_DOWN_SPRITE_NAME] = assetComponents.NewSprite(0, 0)

	// --------------------------------------------------------------------------------
	// Indicator Sprites.
	UISingletonSprites[string(UIConstants.INDICATOR_JUMP)] = assetComponents.NewSprite(0, 0)
	UISingletonSprites[string(UIConstants.INDICATOR_LADDER)] = assetComponents.NewSprite(0, 0)
	UISingletonSprites[string(UIConstants.INDICATOR_MOVEMENT)] = assetComponents.NewSprite(0, 0)
	UISingletonSprites[string(UIConstants.INDICATOR_INTERACT)] = assetComponents.NewSprite(0, 0)
	UISingletonSprites[string(UIConstants.INDICATOR_DESCEND)] = assetComponents.NewSprite(0, 0)

	// --------------------------------------------------------------------------------

	// Register.
	assetComponents.SpritesMapComponent.SetValue(UISingletonSpritesMapEntity, UISingletonSprites)

}
