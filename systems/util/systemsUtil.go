package systemsUtil

import (
	"log"

	"github.com/kainn9/demo/components"
	UIGlobals "github.com/kainn9/demo/globalConfig/UI"
	fontGlobals "github.com/kainn9/demo/globalConfig/font"
	"github.com/kainn9/demo/queries"
	"github.com/yohamta/donburi"
)

func Valid(world donburi.World, entry *donburi.Entry) bool {
	return world.Valid(entry.Entity())
}

func ID(entry *donburi.Entry) int {
	return int(entry.Entity().Id())
}

func CameraEntity(world donburi.World) *donburi.Entry {

	entity, ok := queries.CameraQuery.First(world)

	if !ok {
		log.Fatal("camera query failed.")
	}

	return entity
}

func PlayerEntity(world donburi.World) *donburi.Entry {

	entity, ok := queries.PlayerQuery.First(world)

	if !ok {
		log.Fatal("playerQuery query failed.")
	}
	return entity
}

func UISpritesSingletonEntity(world donburi.World) *donburi.Entry {

	entity, ok := queries.UISingletonSpritesQuery.First(world)

	if !ok {
		log.Fatal("UISpritesSingletonQuery query failed.")
	}

	return entity

}

func UISoundsSingletonEntity(world donburi.World) *donburi.Entry {

	entity, ok := queries.UISingletonSoundsQuery.First(world)

	if !ok {
		log.Fatal("UISpritesSingletonQuery query failed.")
	}

	return entity

}

func ChatPopUpSprite(world donburi.World) *components.Sprite {
	entity, ok := queries.UISingletonSpritesQuery.First(world)

	if !ok {
		log.Fatal("UISpritesSingletonQuery query failed, when getting chat pop up sprite.")
	}

	UISpritesMap := components.SpritesMapComponent.Get(entity)

	return (*UISpritesMap)[UIGlobals.CHAT_BOX_POP_UP_SPRITE_NAME]

}

func ChatPopDownSprite(world donburi.World) *components.Sprite {
	entity, ok := queries.UISingletonSpritesQuery.First(world)

	if !ok {
		log.Fatal("UISpritesSingletonQuery query failed, when getting chat pop down sprite.")
	}

	UISpritesMap := components.SpritesMapComponent.Get(entity)

	return (*UISpritesMap)[UIGlobals.CHAT_BOX_POP_DOWN_SPRITE_NAME]

}

func DefaultFontSpriteMap(world donburi.World) (lower, upper, numbers, special *components.Sprite) {
	entity, ok := queries.UISingletonSpritesQuery.First(world)

	if !ok {
		log.Fatal("UISpritesSingletonQuery query failed.")
	}

	UISpritesMap := components.SpritesMapComponent.Get(entity)

	keyLower := fontGlobals.FONT_DEFAULT_NAME + fontGlobals.FONT_LOWER_CASE_SPRITE_NAME
	keyUpper := fontGlobals.FONT_DEFAULT_NAME + fontGlobals.FONT_UPPER_CASE_SPRITE_NAME

	keyNumbers := fontGlobals.FONT_DEFAULT_NAME + fontGlobals.FONT_NUMBERS_SPRITE_NAME
	keySpecial := fontGlobals.FONT_DEFAULT_NAME + fontGlobals.FONT_SPECIAL_SPRITE_NAME

	lower = (*UISpritesMap)[keyLower]
	upper = (*UISpritesMap)[keyUpper]
	numbers = (*UISpritesMap)[keyNumbers]
	special = (*UISpritesMap)[keySpecial]

	return lower, upper, numbers, special
}

func IsChatActive(world donburi.World) (bool, *donburi.Entry) {

	var isChatActive bool
	var matchedChatEntity *donburi.Entry

	query := queries.ChatQuery

	if query.Count(world) == 0 {
		return false, matchedChatEntity
	}

	query.Each(world, func(chatEntity *donburi.Entry) {

		configAndState := components.ChatStateAndConfigComponent.Get(chatEntity)

		if configAndState.State.Active {
			matchedChatEntity = chatEntity
			isChatActive = true
		}
	})

	return isChatActive, matchedChatEntity
}
