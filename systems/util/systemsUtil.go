package systemsUtil

import (
	"log"

	"github.com/kainn9/demo/components"
	"github.com/kainn9/demo/queries"
	"github.com/yohamta/donburi"
)

func GetCamera(world donburi.World) *components.Camera {

	entry, ok := queries.CameraQuery.FirstEntity(world)

	if !ok {
		log.Fatal("camera query failed.")
	}

	return components.CameraComponent.Get(entry)
}

func GetPlayerPos(world donburi.World) (x, y float64) {

	entry, ok := queries.PlayerQuery.FirstEntity(world)

	if !ok {
		log.Fatal("playerQuery query failed.")
	}

	playerBody := components.RigidBodyComponent.Get(entry)

	return playerBody.X, playerBody.Y
}
