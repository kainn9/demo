package simSystems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	systemsUtil "github.com/kainn9/demo/systems/util"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type PlayerInputHandlerSystem struct {
	scene *coldBrew.Scene
}

func NewPlayerInputHandler(scene *coldBrew.Scene) *PlayerInputHandlerSystem {
	return &PlayerInputHandlerSystem{
		scene: scene,
	}
}

func (*PlayerInputHandlerSystem) Query() *donburi.Query {
	return donburi.NewQuery(
		filter.And(
			filter.Contains(components.PlayerStateComponent),
			filter.Contains(components.InputsComponent),
		),
	)
}

func (sys *PlayerInputHandlerSystem) Run(dt float64, playerEntity *donburi.Entry) {
	inputs := components.InputsComponent.Get(playerEntity)
	playerState := components.PlayerStateComponent.Get(playerEntity)

	activeInput := inputPop(&inputs.Queue)

	if systemsUtil.PlayerStateHelper.PlayerCannotAcceptInputs(playerState) {
		return
	}

	// Handle active input.
	if activeInput == ebiten.KeyRight {
		handleKeyRightLeft(playerState, true)
	}

	if activeInput == ebiten.KeyLeft {
		handleKeyRightLeft(playerState, false)
	}

	if activeInput == ebiten.KeySpace {
		handleKeySpace(playerState, sys.scene.Manager)
	}

	if activeInput == -1 {
		handleNoHorizontalInput(playerState)
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

func handleNoHorizontalInput(playerState *components.PlayerState) {
	playerState.BasicHorizontalMovement = false
}

func inputPop(inputQueue *[]ebiten.Key) ebiten.Key {
	if len(*inputQueue) == 0 {
		return -1
	}

	popped := (*inputQueue)[len(*inputQueue)-1]
	*inputQueue = (*inputQueue)[:len(*inputQueue)-1]
	return popped
}
