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

type LadderHandlerSystem struct {
	scene *coldBrew.Scene
}

func NewLadderHandler(scene *coldBrew.Scene) *LadderHandlerSystem {
	return &LadderHandlerSystem{
		scene: scene,
	}
}

func (sys LadderHandlerSystem) CustomQuery() *donburi.Query {
	return queries.LadderQuery
}

func (sys LadderHandlerSystem) Run(dt float64, _ *donburi.Entry) {

	world := sys.scene.World
	query := sys.CustomQuery()

	noLadderCollisions := true

	playerEntity := systemsUtil.GetPlayerEntity(world)
	playerState := components.PlayerStateComponent.Get(playerEntity)
	playerBody := components.RigidBodyComponent.Get(playerEntity)

	query.Each(world, func(ladderEntity *donburi.Entry) {

		ladderBody := components.RigidBodyComponent.Get(ladderEntity)

		if isColliding, _ := tBokiPhysics.Detector.Detect(playerBody, ladderBody, true); isColliding {
			noLadderCollisions = false
			sys.handleClimb(playerBody, ladderBody, playerState)
		}
	})

	if noLadderCollisions {
		playerState.Climbing = false
	}

}

func (sys LadderHandlerSystem) handleClimb(
	playerBody *tBokiComponents.RigidBody,
	ladderBody *tBokiComponents.RigidBody,
	playerState *components.PlayerState,
) {

	if !playerState.Up && !playerState.Down && !playerState.Climbing {
		return
	}

	if playerState.Jumping || playerState.JumpWindupStart != 0 {
		playerState.Climbing = false
		return
	}

	playerState.Climbing = true

	if playerState.Up {
		playerBody.Vel.Y = -100
		return
	}

	if playerState.Down {
		playerBody.Vel.Y = 100
		return
	}

	playerBody.Vel.Y = 0
}
