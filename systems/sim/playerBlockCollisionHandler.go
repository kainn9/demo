package simSystems

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
	return queries.Block
}

func (sys PlayerBlockCollisionHandlerSystem) Run(dt float64, blockEntity *donburi.Entry) {

	player := systemsUtil.GetPlayerRigidBody(sys.scene.World)
	playerState := systemsUtil.GetPlayerState(sys.scene.World)
	blockBody := components.RigidBodyComponent.Get(blockEntity)

	if isColliding, contacts := tBokiPhysics.Detector.Detect(player, blockBody, true); isColliding {
		tBokiPhysics.Resolver.Resolve(player, blockBody, contacts[0])
		playerState.OnGround = true
	} else {
		playerState.OnGround = false
	}

}
