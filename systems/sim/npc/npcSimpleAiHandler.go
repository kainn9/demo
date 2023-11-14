package simNpcSystems

import (
	"math"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	systemsUtil "github.com/kainn9/demo/systems/util"
	"github.com/kainn9/demo/tags"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type NpcSimpleAiHandlerSystem struct {
	scene *coldBrew.Scene
}

func NewNpcSimpleAiHandler(scene *coldBrew.Scene) *NpcSimpleAiHandlerSystem {
	return &NpcSimpleAiHandlerSystem{
		scene: scene,
	}
}

func (sys NpcSimpleAiHandlerSystem) Query() *donburi.Query {
	return donburi.NewQuery(
		filter.Contains(
			tags.NpcBigBoiTag,
		),
	)
}

func (sys NpcSimpleAiHandlerSystem) Run(dt float64, npcEntity *donburi.Entry) {
	world := sys.scene.World
	npcBody := components.RigidBodyComponent.Get(npcEntity)
	npcState := components.NpcStateComponent.Get(npcEntity)

	playerEntity := systemsUtil.GetPlayerEntity(world)
	playerBody := components.RigidBodyComponent.Get(playerEntity)
	playerState := components.PlayerStateComponent.Get(playerEntity)

	if npcState.Combat.Defeated || playerState.Combat.Defeated {
		return
	}

	patrolRange := 200.0

	tempFactor := 50.0

	if math.Abs(playerBody.Pos.Sub(npcBody.Pos).X) < patrolRange {

		if npcBody.Pos.X < playerBody.Pos.X {
			npcState.Transform.BasicHorizontalMovement = true
			npcBody.Vel.X = tempFactor
			npcState.SetDirectionRight()
		} else {
			npcState.Transform.BasicHorizontalMovement = true
			npcBody.Vel.X = -tempFactor
			npcState.SetDirectionLeft()
		}
	} else {
		npcState.Transform.BasicHorizontalMovement = false
		npcBody.Vel.X = 0
	}
}
