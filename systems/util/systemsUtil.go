package systemsUtil

import (
	"log"

	assetComponents "github.com/kainn9/demo/components/assets"
	fontConstants "github.com/kainn9/demo/constants/font"

	"github.com/kainn9/demo/queries"
	"github.com/yohamta/donburi"

	UIConstants "github.com/kainn9/demo/constants/UI"
)

func GetCameraEntity(world donburi.World) *donburi.Entry {

	entity, ok := queries.CameraQuery.First(world)

	if !ok {
		log.Fatal("camera query failed.")
	}

	return entity
}

func GetPlayerEntity(world donburi.World) *donburi.Entry {

	entity, ok := queries.PlayerQuery.First(world)

	if !ok {
		log.Fatal("playerQuery query failed.")
	}
	return entity
}

func GetUISingletonEntity(world donburi.World) *donburi.Entry {

	entity, ok := queries.UISingletonQuery.First(world)

	if !ok {
		log.Fatal("UISpritesSingletonQuery query failed.")
	}

	return entity

}

func GetChatPopUpSprite(world donburi.World) *assetComponents.Sprite {
	entity, ok := queries.UISingletonQuery.First(world)

	if !ok {
		log.Fatal("UISpritesSingletonQuery query failed, when getting chat pop up sprite.")
	}

	UISpritesMap := assetComponents.SpritesMapComponent.Get(entity)

	return (*UISpritesMap)[UIConstants.CHAT_BOX_POP_UP_SPRITE_NAME]

}

func GetChatPopDownSprite(world donburi.World) *assetComponents.Sprite {
	entity, ok := queries.UISingletonQuery.First(world)

	if !ok {
		log.Fatal("UISpritesSingletonQuery query failed, when getting chat pop down sprite.")
	}

	UISpritesMap := assetComponents.SpritesMapComponent.Get(entity)

	return (*UISpritesMap)[UIConstants.CHAT_BOX_POP_DOWN_SPRITE_NAME]

}

func GetDefaultFontSpriteMap(world donburi.World) (lower, upper, numbers, special *assetComponents.Sprite) {
	entity, ok := queries.UISingletonQuery.First(world)

	if !ok {
		log.Fatal("UISpritesSingletonQuery query failed.")
	}

	UISpritesMap := assetComponents.SpritesMapComponent.Get(entity)

	keyLower := fontConstants.FONT_DEFAULT_NAME + fontConstants.FONT_LOWER_CASE_SPRITE_NAME
	keyUpper := fontConstants.FONT_DEFAULT_NAME + fontConstants.FONT_UPPER_CASE_SPRITE_NAME

	keyNumbers := fontConstants.FONT_DEFAULT_NAME + fontConstants.FONT_NUMBERS_SPRITE_NAME
	keySpecial := fontConstants.FONT_DEFAULT_NAME + fontConstants.FONT_SPECIAL_SPRITE_NAME

	lower = (*UISpritesMap)[keyLower]
	upper = (*UISpritesMap)[keyUpper]
	numbers = (*UISpritesMap)[keyNumbers]
	special = (*UISpritesMap)[keySpecial]

	return lower, upper, numbers, special
}
