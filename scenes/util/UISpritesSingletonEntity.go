package scenesUtil

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	UIConstants "github.com/kainn9/demo/constants/UI"
	fontConstants "github.com/kainn9/demo/constants/font"

	"github.com/kainn9/demo/tags"
)

func AddUISpritesSingletonEntity(scene *coldBrew.Scene) {

	// Initialize UI Singleton SpritesMap.
	UISingletonSpritesMapEntity := scene.AddEntity(
		components.SpritesMapComponent,
		tags.UISingletonTag,
	)

	UISingletonSprites := make(map[string]*components.Sprite)

	// --------------------------------------------------------------------------------
	// Default Font
	keyDefault := fontConstants.FONT_DEFAULT_NAME
	keyLower := keyDefault + fontConstants.FONT_LOWER_CASE_SPRITE_NAME
	keyUpper := keyDefault + fontConstants.FONT_UPPER_CASE_SPRITE_NAME
	keyNumbers := keyDefault + fontConstants.FONT_NUMBERS_SPRITE_NAME
	keySpecial := keyDefault + fontConstants.FONT_SPECIAL_SPRITE_NAME

	UISingletonSprites[keyLower] = components.NewSprite(0, 0)
	UISingletonSprites[keyUpper] = components.NewSprite(0, 0)
	UISingletonSprites[keyNumbers] = components.NewSprite(0, 0)
	UISingletonSprites[keySpecial] = components.NewSprite(0, 0)

	// --------------------------------------------------------------------------------
	// Chat Box Sprites.
	UISingletonSprites[UIConstants.CHAT_BOX_POP_UP_SPRITE_NAME] = components.NewSprite(0, 0)
	UISingletonSprites[UIConstants.CHAT_BOX_POP_DOWN_SPRITE_NAME] = components.NewSprite(0, 0)

	// --------------------------------------------------------------------------------
	// Indicator Sprites.
	UISingletonSprites[string(UIConstants.INDICATOR_JUMP)] = components.NewSprite(0, 0)
	UISingletonSprites[string(UIConstants.INDICATOR_LADDER)] = components.NewSprite(0, 0)
	UISingletonSprites[string(UIConstants.INDICATOR_MOVEMENT)] = components.NewSprite(0, 0)
	UISingletonSprites[string(UIConstants.INDICATOR_INTERACT)] = components.NewSprite(0, 0)
	UISingletonSprites[string(UIConstants.INDICATOR_DESCEND)] = components.NewSprite(0, 0)

	// --------------------------------------------------------------------------------

	// Register.
	components.SpritesMapComponent.SetValue(UISingletonSpritesMapEntity, UISingletonSprites)

}
