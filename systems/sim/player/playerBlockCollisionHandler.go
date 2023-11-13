package simPlayerSystems

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	"github.com/kainn9/demo/queries"
	systemsUtil "github.com/kainn9/demo/systems/util"
	tBokiPhysics "github.com/kainn9/tteokbokki/physics"
	"github.com/yohamta/donburi"
)

type PlayerBlockCollisionHandlerSystem struct {
	scene *coldBrew.Scene
}

func NewPlayerBlockCollisionHandler(scene *coldBrew.Scene) *PlayerBlockCollisionHandlerSystem {
	return &PlayerBlockCollisionHandlerSystem{
		scene: scene,
	}
}

func (sys PlayerBlockCollisionHandlerSystem) Query() *donburi.Query {
	return queries.BlockQuery
}

func (sys PlayerBlockCollisionHandlerSystem) Run(dt float64, blockEntity *donburi.Entry) {

	world := sys.scene.World

	playerEntity := systemsUtil.GetPlayerEntity(world)
	playerState := components.PlayerStateComponent.Get(playerEntity)
	playerBody := components.RigidBodyComponent.Get(playerEntity)

	blockBody := components.RigidBodyComponent.Get(blockEntity)

	if playerState.Collision.Climbing {
		return
	}

	if isColliding, contacts := tBokiPhysics.Detector.Detect(playerBody, blockBody, true); isColliding {

		tBokiPhysics.Resolver.Resolve(playerBody, blockBody, contacts[0])

		playerBottom := playerBody.Pos.Y + playerBody.Polygon.Box.Height/2
		blockop := blockBody.Pos.Y - blockBody.Polygon.Box.Height/2

		if playerBottom <= blockop {
			playerState.Collision.OnGround = true
		}

	}

}
