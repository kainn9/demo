package systemsUtil

import (
	"log"

	"github.com/kainn9/demo/components"
	"github.com/kainn9/demo/queries"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
	tBokiVec "github.com/kainn9/tteokbokki/math/vec"
	"github.com/yohamta/donburi"
)

func GetCamera(world donburi.World) *components.Camera {

	entity, ok := queries.CameraQuery.FirstEntity(world)

	if !ok {
		log.Fatal("camera query failed.")
	}

	return components.CameraComponent.Get(entity)
}

func GetPlayerPos(world donburi.World) tBokiVec.Vec2 {

	entity, ok := queries.PlayerQuery.FirstEntity(world)

	if !ok {
		log.Fatal("playerQuery query failed.")
	}

	playerBody := components.RigidBodyComponent.Get(entity)

	return playerBody.Pos
}

func GetPlayerRigidBody(world donburi.World) *tBokiComponents.RigidBody {

	entity, ok := queries.PlayerQuery.FirstEntity(world)

	if !ok {
		log.Fatal("playerQuery query failed.")
	}

	playerBody := components.RigidBodyComponent.Get(entity)

	return playerBody
}

func GetPlayerState(world donburi.World) *components.PlayerState {

	entity, ok := queries.PlayerQuery.FirstEntity(world)

	if !ok {
		log.Fatal("playerQuery query failed.")
	}

	state := components.PlayerStateComponent.Get(entity)

	return state
}
