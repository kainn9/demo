package simPlayerSystems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	inputGlobals "github.com/kainn9/demo/globalConfig/input"
	sharedAnimationGlobals "github.com/kainn9/demo/globalConfig/sharedAnimation"

	scenesUtil "github.com/kainn9/demo/scenes/util"
	systemsUtil "github.com/kainn9/demo/systems/util"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type PlayerPhysicsInputProcessorSystem struct {
	scene *coldBrew.Scene
}

func NewPlayerPhysicsInputProcessor(scene *coldBrew.Scene) *PlayerPhysicsInputProcessorSystem {
	return &PlayerPhysicsInputProcessorSystem{
		scene: scene,
	}
}

func (PlayerPhysicsInputProcessorSystem) Query() *donburi.Query {
	return donburi.NewQuery(
		filter.And(
			filter.Contains(components.PlayerStateComponent),
			filter.Contains(components.InputsComponent),
		),
	)
}

func (sys PlayerPhysicsInputProcessorSystem) Run(dt float64, playerEntity *donburi.Entry) {
	inputs := components.InputsComponent.Get(playerEntity)
	playerState := components.PlayerStateComponent.Get(playerEntity)

	if systemsUtil.IsChatActive(sys.scene.World) {
		playerState.Transform.BasicHorizontalMovement = false
	}

	if playerState.Combat.Hit {
		return
	}

	activeInput := sys.inputShift(&inputs.Queue)

	left, right, jump, up, down, _, attackPrimary := inputGlobals.ALL_BINDS()

	if activeInput == attackPrimary {
		sys.handleKeyPrimaryAtk(playerState, playerEntity)

	}

	// Left/Right movement.
	if activeInput == left {
		sys.handleKeyRightLeft(playerState, false)
	}

	if activeInput == right {
		sys.handleKeyRightLeft(playerState, true)
	}

	if activeInput == inputGlobals.RELEASED_HORIZONTAL {
		playerState.Transform.BasicHorizontalMovement = false
	}

	// Jumping or descending platform.
	if activeInput == jump {
		sys.handleKeySpace(playerState)
	}

	if activeInput == inputGlobals.COMBO_DOWN_SPACE && playerState.Collision.OnGround {
		playerState.Transform.PhaseThroughPlatforms = true
	}

	if activeInput == up {
		playerState.Transform.Up = true
	}

	if activeInput == down {
		playerState.Transform.Down = true
	}

	if activeInput == inputGlobals.RELEASED_VERTICAL_UP {
		playerState.Transform.Up = false
	}

	if activeInput == inputGlobals.RELEASED_VERTICAL_DOWN {
		playerState.Transform.Down = false
	}

}

func (sys PlayerPhysicsInputProcessorSystem) handleKeyRightLeft(playerState *components.PlayerState, right bool) {

	playerState.Transform.BasicHorizontalMovement = true

	if right {
		playerState.SetDirectionRight()

	} else {
		playerState.SetDirectionLeft()
	}
}

func (sys PlayerPhysicsInputProcessorSystem) handleKeySpace(playerState *components.PlayerState) {

	if !playerState.Collision.OnGround && !playerState.Collision.Climbing {
		return
	}

	playerState.Transform.JumpTriggered = true

}

func (sys PlayerPhysicsInputProcessorSystem) inputPop(inputQueue *[]ebiten.Key) ebiten.Key {
	if len(*inputQueue) == 0 {
		return inputGlobals.NO_INPUT
	}

	popped := (*inputQueue)[len(*inputQueue)-1]
	*inputQueue = (*inputQueue)[:len(*inputQueue)-1]
	return popped
}

func (sys PlayerPhysicsInputProcessorSystem) inputShift(inputQueue *[]ebiten.Key) ebiten.Key {
	if len(*inputQueue) == 0 {
		return inputGlobals.NO_INPUT
	}

	popped := (*inputQueue)[0]
	*inputQueue = (*inputQueue)[1:]
	return popped
}

func (sys PlayerPhysicsInputProcessorSystem) handleKeyPrimaryAtk(playerState *components.PlayerState, playerEntity *donburi.Entry) {
	playerState.Combat.Attacking = true
	playerState.Combat.AttackStartTick = sys.scene.Manager.TickHandler.CurrentTick()
	playerState.Combat.CurrentAttack = sharedAnimationGlobals.CHAR_STATE_ATTACK_PRIMARY

	attackState := components.NewAttackState(
		sys.scene.Manager.TickHandler.CurrentTick(),
		int(playerEntity.Entity().Id()),
		true,
	)

	scenesUtil.AddAttackEntity(sys.scene, *attackState)
}
