package simSystems

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	"github.com/kainn9/demo/queries"
	systemsUtil "github.com/kainn9/demo/systems/util"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
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

func (sys PlayerBlockCollisionHandlerSystem) CustomQuery() *donburi.Query {
	return queries.Block
}

func (sys PlayerBlockCollisionHandlerSystem) Run(dt float64, _ *donburi.Entry) {
	query := sys.CustomQuery()
	world := sys.scene.World

	playerEntity := systemsUtil.GetPlayerEntity(world)
	playerState := components.PlayerStateComponent.Get(playerEntity)
	playerBody := components.RigidBodyComponent.Get(playerEntity)

	playerState.OnGround = false

	query.Each(world, func(blockEntity *donburi.Entry) {

		blockBody := components.RigidBodyComponent.Get(blockEntity)
		runHelper(playerBody, blockBody, playerState)

	})

}

func runHelper(playerBody, blockBody *tBokiComponents.RigidBody, playerState *components.PlayerState) {
	if isColliding, contacts := tBokiPhysics.Detector.Detect(playerBody, blockBody, true); isColliding {
		tBokiPhysics.Resolver.Resolve(playerBody, blockBody, contacts[0])
		playerState.OnGround = true
	}
}
