package simNpcSystems

import (
	"math"
	"math/rand"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	sharedStateGlobals "github.com/kainn9/demo/globalConfig/sharedState"
	scenesUtil "github.com/kainn9/demo/scenes/util"
	systemsUtil "github.com/kainn9/demo/systems/util"
	"github.com/kainn9/demo/tags"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
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
			tags.NpcThugTag,
		),
	)
}

func (sys NpcSimpleAiHandlerSystem) Run(dt float64, npcEntity *donburi.Entry) {
	world := sys.scene.World
	npcBody := components.RigidBodyComponent.Get(npcEntity)
	npcState := components.NpcStateComponent.Get(npcEntity)

	playerEntity := systemsUtil.PlayerEntity(world)
	playerBody := components.RigidBodyComponent.Get(playerEntity)
	playerState := components.PlayerStateComponent.Get(playerEntity)

	chatIsActive, _ := systemsUtil.IsChatActive(sys.scene.World)

	frozenGameState := npcState.Combat.Defeated || playerState.Combat.Defeated || chatIsActive

	if frozenGameState {
		sys.idleNpc(npcState, npcBody)
		return
	}

	npcIsAttacking := npcState.Combat.CurrentAttack != ""

	if npcIsAttacking {
		return
	}

	distanceBetweenNpcAndPlayer := math.Abs(playerBody.Pos.Sub(npcBody.Pos).X)

	playerInsideAttackRange := distanceBetweenNpcAndPlayer < npcState.Combat.AttackRange

	sys.LockNpcDirectionOnPlayer(playerBody, npcBody, npcState)

	if playerInsideAttackRange && rand.Float64() < 0.15 {
		sys.idleNpc(npcState, npcBody)
		sys.triggerPrimaryAttack(npcState, npcEntity)
		return
	}

	sys.patrol(distanceBetweenNpcAndPlayer, npcState, npcBody, playerBody)

}

func (sys NpcSimpleAiHandlerSystem) idleNpc(npcState *components.NpcState, npcBody *tBokiComponents.RigidBody) {
	npcState.Transform.BasicHorizontalMovement = false
	npcBody.Vel.X = 0
}

func (NpcSimpleAiHandlerSystem) LockNpcDirectionOnPlayer(playerBody, npcBody *tBokiComponents.RigidBody, npcState *components.NpcState) {

	if npcBody.Pos.X < playerBody.Pos.X {
		npcState.SetDirectionRight()
	} else {
		npcState.SetDirectionLeft()
	}
}

func (sys NpcSimpleAiHandlerSystem) triggerPrimaryAttack(npcState *components.NpcState, npcEntity *donburi.Entry) {

	npcState.Combat.AttackStartTick = sys.scene.Manager.TickHandler.CurrentTick()
	npcState.Combat.CurrentAttack = sharedStateGlobals.CHAR_STATE_ATTACK_PRIMARY

	attackState := components.NewAttackState(
		string(sharedStateGlobals.CHAR_STATE_ATTACK_PRIMARY),
		npcEntity,
	)

	scenesUtil.AddAttackEntity(sys.scene, *attackState)
}

func (sys NpcSimpleAiHandlerSystem) patrol(
	distanceBetweenNpcAndPlayer float64,
	npcState *components.NpcState,
	npcBody, playerBody *tBokiComponents.RigidBody,
) {
	playerWithinPatrolRange := distanceBetweenNpcAndPlayer < npcState.Combat.PatrolRange

	if playerWithinPatrolRange {
		sys.patrolLeft(npcState, npcBody)
		sys.patrolRight(npcState, npcBody)

	} else {
		sys.idleNpc(npcState, npcBody)
	}
}
func (sys NpcSimpleAiHandlerSystem) patrolLeft(npcState *components.NpcState, npcBody *tBokiComponents.RigidBody) {
	if npcState.IsRight() {
		return
	}

	if npcBody.Pos.X > npcState.Combat.MaxLeft {
		sys.startHorizontalMovement(npcState, npcBody)
	} else {
		sys.idleNpc(npcState, npcBody)
	}
}

func (sys NpcSimpleAiHandlerSystem) patrolRight(npcState *components.NpcState, npcBody *tBokiComponents.RigidBody) {
	if npcState.IsLeft() {
		return
	}

	if npcBody.Pos.X < npcState.Combat.MaxRight {
		sys.startHorizontalMovement(npcState, npcBody)
	} else {
		sys.idleNpc(npcState, npcBody)
	}
}

func (sys NpcSimpleAiHandlerSystem) startHorizontalMovement(npcState *components.NpcState, npcBody *tBokiComponents.RigidBody) {
	npcState.Transform.BasicHorizontalMovement = true
	npcBody.Vel.X = npcState.Transform.Speed * npcState.Direction()
}
