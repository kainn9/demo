package simSystems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	"github.com/kainn9/demo/constants"
	systemsUtil "github.com/kainn9/demo/systems/util"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type PlayerMovementInputHandlerSystem struct {
	scene *coldBrew.Scene
}

func NewPlayerMovementInputHandler(scene *coldBrew.Scene) *PlayerMovementInputHandlerSystem {
	return &PlayerMovementInputHandlerSystem{
		scene: scene,
	}
}

func (*PlayerMovementInputHandlerSystem) Query() *donburi.Query {
	return donburi.NewQuery(
		filter.And(
			filter.Contains(components.PlayerStateComponent),
			filter.Contains(components.InputsComponent),
		),
	)
}

func (sys *PlayerMovementInputHandlerSystem) Run(dt float64, playerEntity *donburi.Entry) {
	inputs := components.InputsComponent.Get(playerEntity)
	playerState := components.PlayerStateComponent.Get(playerEntity)

	activeInput := inputShift(&inputs.Queue)

	if systemsUtil.PlayerStateHelper.PlayerCannotAcceptInputs(playerState) {
		return
	}

	left, right, jump, up, down, interact := constants.AllBinds()

	// Left/Right movement.
	if activeInput == left {
		handleKeyRightLeft(playerState, false)
	}

	if activeInput == right {
		handleKeyRightLeft(playerState, true)
	}

	if activeInput == constants.RELEASED_HORIZONTAL {
		playerState.BasicHorizontalMovement = false
	}

	// Jumping or descending platform.
	if activeInput == jump {
		handleKeySpace(playerState, sys.scene.Manager)
	}

	if activeInput == constants.COMBO_DOWN_SPACE && playerState.OnGround {
		playerState.PhaseThroughPlatforms = true
	}

	if activeInput == up {
		playerState.Up = true
	}

	if activeInput == down {
		playerState.Down = true
	}

	if activeInput == interact {
		playerState.Interact = true
	}

}

func handleKeyRightLeft(playerState *components.PlayerState, right bool) {

	playerState.BasicHorizontalMovement = true

	if right {
		playerState.SetDirectionRight()

	} else {
		playerState.SetDirectionLeft()
	}
}

func handleKeySpace(playerState *components.PlayerState, m *coldBrew.Manager) {
	tickHandler := m.TickHandler

	playerPreparingToJump := playerState.JumpWindupStart != 0

	if !playerState.OnGround || playerPreparingToJump {
		return
	}

	playerState.JumpWindupStart = tickHandler.CurrentTick()

}

func inputPop(inputQueue *[]ebiten.Key) ebiten.Key {
	if len(*inputQueue) == 0 {
		return constants.NO_INPUT
	}

	popped := (*inputQueue)[len(*inputQueue)-1]
	*inputQueue = (*inputQueue)[:len(*inputQueue)-1]
	return popped
}

func inputShift(inputQueue *[]ebiten.Key) ebiten.Key {
	if len(*inputQueue) == 0 {
		return constants.NO_INPUT
	}

	popped := (*inputQueue)[0]
	*inputQueue = (*inputQueue)[1:]
	return popped
}
