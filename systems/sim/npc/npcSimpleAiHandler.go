package simNpcSystems

import (
	"math"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	sharedAnimationGlobals "github.com/kainn9/demo/globalConfig/sharedAnimation"
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

	chatIsActive, _ := systemsUtil.IsChatActive(sys.scene.World)

	if npcState.Combat.CurrentAttack != "" {
		return
	}

	if npcState.Combat.Defeated || playerState.Combat.Defeated || chatIsActive {
		sys.idleNpc(npcState, npcBody)
		return
	}

	dist := math.Abs(playerBody.Pos.Sub(npcBody.Pos).X)

	if dist < npcState.Combat.AttackRange {
		sys.idleNpc(npcState, npcBody)
		sys.primaryAttack(npcState, npcEntity)
		return
	}

	if dist < npcState.Combat.PatrolRange {

		if npcBody.Pos.X < playerBody.Pos.X {
			npcState.Transform.BasicHorizontalMovement = true
			npcBody.Vel.X = npcState.Transform.Speed
			npcState.SetDirectionRight()
		} else {
			npcState.Transform.BasicHorizontalMovement = true
			npcBody.Vel.X = -npcState.Transform.Speed
			npcState.SetDirectionLeft()
		}
	} else {
		sys.idleNpc(npcState, npcBody)
	}
}

func (sys NpcSimpleAiHandlerSystem) idleNpc(npcState *components.NpcState, npcBody *tBokiComponents.RigidBody) {
	npcState.Transform.BasicHorizontalMovement = false
	npcBody.Vel.X = 0
}

// func (sys NpcSimpleAiHandlerSystem) primaryAttack(npcState *components.NpcState) {
// 	npcState.Combat.CurrentAttack = sharedAnimationGlobals.CHAR_STATE_ATTACK_PRIMARY
// 	npcState.Combat.AttackStartTick = sys.scene.Manager.TickHandler.CurrentTick()
// }

func (sys NpcSimpleAiHandlerSystem) primaryAttack(npcState *components.NpcState, npcEntity *donburi.Entry) {
	npcState.Combat.AttackStartTick = sys.scene.Manager.TickHandler.CurrentTick()
	npcState.Combat.CurrentAttack = sharedAnimationGlobals.CHAR_STATE_ATTACK_PRIMARY

	attackState := components.NewAttackState(
		sys.scene.Manager.TickHandler.CurrentTick(),
		int(npcEntity.Entity().Id()),
		string(sharedAnimationGlobals.CHAR_STATE_ATTACK_PRIMARY),
		false,
	)

	scenesUtil.AddAttackEntity(sys.scene, *attackState)
}
