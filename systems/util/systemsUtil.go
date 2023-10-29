package systemsUtil

import (
	"log"

	"github.com/kainn9/demo/queries"
	"github.com/yohamta/donburi"
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

func GetChatPopUpEntity(world donburi.World) *donburi.Entry {
	entity, ok := queries.ChatPopUpEntityQuery.First(world)

	if !ok {
		log.Fatal("chatPopUpEntity query failed.")
	}

	return entity
}

func GetChatPopDownEntity(world donburi.World) *donburi.Entry {
	entity, ok := queries.ChatPopDownEntityQuery.First(world)

	if !ok {
		log.Fatal("chatPopDownEntity query failed.")
	}

	return entity
}
