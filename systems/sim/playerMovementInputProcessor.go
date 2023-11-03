package simSystems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	inputConstants "github.com/kainn9/demo/constants/input"

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

	left, right, jump, up, down, _ := inputConstants.ALL_BINDS()

	// Left/Right movement.
	if activeInput == left {
		sys.handleKeyRightLeft(playerState, false)
	}

	if activeInput == right {
		sys.handleKeyRightLeft(playerState, true)
	}

	if activeInput == inputConstants.RELEASED_HORIZONTAL {
		playerState.Transform.BasicHorizontalMovement = false
	}

	// Jumping or descending platform.
	if activeInput == jump {
		sys.handleKeySpace(playerState, sys.scene.Manager)
	}

	if activeInput == inputConstants.COMBO_DOWN_SPACE && playerState.Collision.OnGround {
		playerState.Transform.PhaseThroughPlatforms = true
	}

	if activeInput == up {
		playerState.Transform.Up = true
	}

	if activeInput == down {
		playerState.Transform.Down = true
	}

	if activeInput == inputConstants.RELEASED_VERTICAL {
		playerState.Transform.Up = false
		playerState.Transform.Down = false
	}

}

func (sys PlayerMovementInputProcessorSystem) handleKeyRightLeft(playerState *components.PlayerState, right bool) {

	playerState.Transform.BasicHorizontalMovement = true

	if right {
		playerState.SetDirectionRight()

	} else {
		playerState.SetDirectionLeft()
	}
}

func (sys PlayerMovementInputProcessorSystem) handleKeySpace(playerState *components.PlayerState, m *coldBrew.Manager) {
	tickHandler := m.TickHandler

	playerPreparingToJump := playerState.Transform.JumpWindupStart != 0

	if (!playerState.Collision.OnGround && !playerState.Collision.Climbing) || playerPreparingToJump {
		return
	}

	playerState.Transform.JumpWindupStart = tickHandler.CurrentTick()

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
