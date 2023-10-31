package simSystems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	inputConstants "github.com/kainn9/demo/constants/input"
	systemsUtil "github.com/kainn9/demo/systems/util"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type PlayerMovementInputProcessorSystem struct {
	scene *coldBrew.Scene
}

func NewPlayerMovementInputProcessor(scene *coldBrew.Scene) *PlayerMovementInputProcessorSystem {
	return &PlayerMovementInputProcessorSystem{
		scene: scene,
	}
}

func (PlayerMovementInputProcessorSystem) Query() *donburi.Query {
	return donburi.NewQuery(
		filter.And(
			filter.Contains(components.PlayerStateComponent),
			filter.Contains(components.InputsComponent),
		),
	)
}

func (sys PlayerMovementInputProcessorSystem) Run(dt float64, playerEntity *donburi.Entry) {
	inputs := components.InputsComponent.Get(playerEntity)
	playerState := components.PlayerStateComponent.Get(playerEntity)

	activeInput := sys.inputShift(&inputs.Queue)

	if systemsUtil.PlayerStateHelper.PlayerCannotAcceptInputs(playerState) {
		return
	}

	left, right, jump, up, down, interact := inputConstants.ALL_BINDS()

	// Left/Right movement.
	if activeInput == left {
		sys.handleKeyRightLeft(playerState, false)
	}

	if activeInput == right {
		sys.handleKeyRightLeft(playerState, true)
	}

	if activeInput == inputConstants.RELEASED_HORIZONTAL {
		playerState.BasicHorizontalMovement = false
	}

	// Jumping or descending platform.
	if activeInput == jump {
		sys.handleKeySpace(playerState, sys.scene.Manager)
	}

	if activeInput == inputConstants.COMBO_DOWN_SPACE && playerState.OnGround {
		playerState.PhaseThroughPlatforms = true
	}

	if activeInput == up {
		playerState.Up = true
	}

	if activeInput == down {
		playerState.Down = true
	}

	if activeInput == inputConstants.RELEASED_VERTICAL {
		playerState.Up = false
		playerState.Down = false
	}

	if activeInput == interact {
		playerState.Interact = true
	}

}

func (sys PlayerMovementInputProcessorSystem) handleKeyRightLeft(playerState *components.PlayerState, right bool) {

	playerState.BasicHorizontalMovement = true

	if right {
		playerState.SetDirectionRight()

	} else {
		playerState.SetDirectionLeft()
	}
}

func (sys PlayerMovementInputProcessorSystem) handleKeySpace(playerState *components.PlayerState, m *coldBrew.Manager) {
	tickHandler := m.TickHandler

	playerPreparingToJump := playerState.JumpWindupStart != 0

	if (!playerState.OnGround && !playerState.Climbing) || playerPreparingToJump {
		return
	}

	playerState.JumpWindupStart = tickHandler.CurrentTick()

}

func (sys PlayerMovementInputProcessorSystem) inputPop(inputQueue *[]ebiten.Key) ebiten.Key {
	if len(*inputQueue) == 0 {
		return inputConstants.NO_INPUT
	}

	popped := (*inputQueue)[len(*inputQueue)-1]
	*inputQueue = (*inputQueue)[:len(*inputQueue)-1]
	return popped
}

func (sys PlayerMovementInputProcessorSystem) inputShift(inputQueue *[]ebiten.Key) ebiten.Key {
	if len(*inputQueue) == 0 {
		return inputConstants.NO_INPUT
	}

	popped := (*inputQueue)[0]
	*inputQueue = (*inputQueue)[1:]
	return popped
}
