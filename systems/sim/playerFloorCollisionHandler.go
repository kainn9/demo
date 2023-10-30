package simSystems

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	"github.com/kainn9/demo/queries"
	systemsUtil "github.com/kainn9/demo/systems/util"
	tBokiPhysics "github.com/kainn9/tteokbokki/physics"
	"github.com/yohamta/donburi"
)

type PlayerFloorCollisionHandlerSystem struct {
	scene *coldBrew.Scene
}

func NewPlayerFloorCollisionHandler(scene *coldBrew.Scene) *PlayerFloorCollisionHandlerSystem {
	return &PlayerFloorCollisionHandlerSystem{
		scene: scene,
	}
}

func (sys PlayerFloorCollisionHandlerSystem) Query() *donburi.Query {
	return queries.FloorQuery
}

func (sys PlayerFloorCollisionHandlerSystem) Run(dt float64, floorEntity *donburi.Entry) {

	world := sys.scene.World

	playerEntity := systemsUtil.GetPlayerEntity(world)
	playerState := components.PlayerStateComponent.Get(playerEntity)
	playerBody := components.RigidBodyComponent.Get(playerEntity)

	floorBody := components.RigidBodyComponent.Get(floorEntity)

	if isColliding, contacts := tBokiPhysics.Detector.Detect(playerBody, floorBody, true); isColliding {
		tBokiPhysics.Resolver.Resolve(playerBody, floorBody, contacts[0])
		playerState.OnGround = true
	}

}
