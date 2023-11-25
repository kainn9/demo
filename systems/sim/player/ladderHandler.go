package simPlayerSystems

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	"github.com/kainn9/demo/queries"
	systemsUtil "github.com/kainn9/demo/systems/util"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
	tBokiPhysics "github.com/kainn9/tteokbokki/physics"
	"github.com/yohamta/donburi"
)

type LadderHandlerSystem struct {
	scene *coldBrew.Scene
}

func NewLadderHandler(scene *coldBrew.Scene) *LadderHandlerSystem {
	return &LadderHandlerSystem{
		scene: scene,
	}
}

func (sys LadderHandlerSystem) LadderQuery() *donburi.Query {
	return queries.LadderQuery
}

func (sys LadderHandlerSystem) Run(dt float64, _ *donburi.Entry) {

	world := sys.scene.World

	playerEntity := systemsUtil.PlayerEntity(world)
	playerState := components.PlayerStateComponent.Get(playerEntity)
	playerBody := components.RigidBodyComponent.Get(playerEntity)

	noLadderCollisions := true

	sys.LadderQuery().Each(world, func(ladderEntity *donburi.Entry) {

		ladderBody := components.RigidBodyComponent.Get(ladderEntity)

		if isColliding, _ := tBokiPhysics.Detector.Detect(playerBody, ladderBody, true); isColliding {
			noLadderCollisions = false
			sys.handleClimb(playerBody, ladderBody, playerState)
		}
	})

	if noLadderCollisions {
		playerState.Collision.Climbing = false
	}

}

func (sys LadderHandlerSystem) handleClimb(
	playerBody *tBokiComponents.RigidBody,
	ladderBody *tBokiComponents.RigidBody,
	playerState *components.PlayerState,
) {

	// Return if the player is not pressing up or down and is not already climbing.
	if !playerState.Transform.Up && !playerState.Transform.Down && !playerState.Collision.Climbing {
		return
	}

	// Return if the player is jumping.
	if playerState.Transform.Jumping {
		playerState.Collision.Climbing = false
		return
	}

	playerState.Collision.Climbing = true

	if playerState.Transform.Up {
		playerBody.Vel.Y = -100
		return
	}

	if playerState.Transform.Down {
		playerBody.Vel.Y = 100
		return
	}

	// Disable gravity when climbing.
	playerBody.Vel.Y = 0
}
