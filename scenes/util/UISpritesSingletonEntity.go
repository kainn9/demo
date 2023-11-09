package scenesUtil

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	UIGlobals "github.com/kainn9/demo/globalConfig/UI"
	fontGlobals "github.com/kainn9/demo/globalConfig/font"

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
	keyDefault := fontGlobals.FONT_DEFAULT_NAME
	keyLower := keyDefault + fontGlobals.FONT_LOWER_CASE_SPRITE_NAME
	keyUpper := keyDefault + fontGlobals.FONT_UPPER_CASE_SPRITE_NAME
	keyNumbers := keyDefault + fontGlobals.FONT_NUMBERS_SPRITE_NAME
	keySpecial := keyDefault + fontGlobals.FONT_SPECIAL_SPRITE_NAME

	UISingletonSprites[keyLower] = components.NewSprite(0, 0)
	UISingletonSprites[keyUpper] = components.NewSprite(0, 0)
	UISingletonSprites[keyNumbers] = components.NewSprite(0, 0)
	UISingletonSprites[keySpecial] = components.NewSprite(0, 0)

	// --------------------------------------------------------------------------------
	// Chat Box Sprites.
	UISingletonSprites[UIGlobals.CHAT_BOX_POP_UP_SPRITE_NAME] = components.NewSprite(0, 0)
	UISingletonSprites[UIGlobals.CHAT_BOX_POP_DOWN_SPRITE_NAME] = components.NewSprite(0, 0)

	// --------------------------------------------------------------------------------
	// Indicator Sprites.
	UISingletonSprites[string(UIGlobals.INDICATOR_JUMP)] = components.NewSprite(0, 0)
	UISingletonSprites[string(UIGlobals.INDICATOR_LADDER)] = components.NewSprite(0, 0)
	UISingletonSprites[string(UIGlobals.INDICATOR_MOVEMENT)] = components.NewSprite(0, 0)
	UISingletonSprites[string(UIGlobals.INDICATOR_INTERACT)] = components.NewSprite(0, 0)
	UISingletonSprites[string(UIGlobals.INDICATOR_DESCEND)] = components.NewSprite(0, 0)

	// --------------------------------------------------------------------------------

	// Register.
	components.SpritesMapComponent.SetValue(UISingletonSpritesMapEntity, UISingletonSprites)

}
