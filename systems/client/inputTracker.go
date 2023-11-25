package clientSystems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	inputGlobals "github.com/kainn9/demo/globalConfig/input"
	systemsUtil "github.com/kainn9/demo/systems/util"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type InputTrackerSystem struct {
	scene                   *coldBrew.Scene
	tickLeftKeyLastPressed  int
	tickRightKeyLastPressed int
}

func NewInputTracker(scene *coldBrew.Scene) *InputTrackerSystem {
	return &InputTrackerSystem{
		scene:                   scene,
		tickLeftKeyLastPressed:  0,
		tickRightKeyLastPressed: 0,
	}
}

func (InputTrackerSystem) InputsQuery() *donburi.Query {
	return donburi.NewQuery(
		filter.Contains(components.InputsComponent),
	)
}

func (sys *InputTrackerSystem) Sync(_ *donburi.Entry) {

	world := sys.scene.World
	playerEntity := systemsUtil.PlayerEntity(world)

	sys.processInteractionInput(playerEntity)

	sys.trackLeftRightLastPressed()

	sys.InputsQuery().Each(world, func(inputsEntity *donburi.Entry) {
		sys.processMovementInputs(inputsEntity)
	})

}

func (sys InputTrackerSystem) processInteractionInput(playerEntity *donburi.Entry) {
	playerState := components.PlayerStateComponent.Get(playerEntity)

	if inpututil.IsKeyJustPressed(inputGlobals.KeyInteract()) {
		playerState.IsInteracting = true
	} else {
		playerState.IsInteracting = false
	}

}

func (sys InputTrackerSystem) processMovementInputs(inputsEntity *donburi.Entry) {
	inputs := components.InputsComponent.Get(inputsEntity)
	world := sys.scene.World

	playerEntity := systemsUtil.PlayerEntity(world)
	playerState := components.PlayerStateComponent.Get(playerEntity)

	if sys.playerCannotAcceptMovementInputs(world, playerState) {
		return
	}

	// Left/Right movement.
	leftPriority := sys.tickLeftKeyLastPressed > sys.tickRightKeyLastPressed

	left := inputGlobals.KeyLeft()
	right := inputGlobals.KeyRight()
	dodge := inputGlobals.KeyDodge()
	jump := inputGlobals.KeyJump()
	up := inputGlobals.KeyUp()
	down := inputGlobals.KeyDown()
	attackPrimary := inputGlobals.KeyPrimaryAttack()

	// Dodge.
	if inpututil.IsKeyJustPressed(dodge) {
		sys.addUniqueKey(&inputs.Queue, dodge)
	}

	if playerState.Combat.Attacking {
		return
	}

	if ebiten.IsKeyPressed(left) && leftPriority {
		sys.addUniqueKey(&inputs.Queue, left)
	}

	if ebiten.IsKeyPressed(right) && !leftPriority {
		sys.addUniqueKey(&inputs.Queue, right)
	}

	if !ebiten.IsKeyPressed(left) && !ebiten.IsKeyPressed(right) {
		sys.addUniqueKey(&inputs.Queue, inputGlobals.RELEASED_HORIZONTAL)
	}

	// Jump.
	if inpututil.IsKeyJustPressed(jump) && !ebiten.IsKeyPressed(down) {
		sys.addUniqueKey(&inputs.Queue, jump)
	}

	// Phase through platforms.
	if inpututil.IsKeyJustPressed(jump) && ebiten.IsKeyPressed(down) {
		sys.addUniqueKey(&inputs.Queue, inputGlobals.COMBO_DOWN_SPACE)
	}

	// Climb up.
	if ebiten.IsKeyPressed(up) {
		sys.addUniqueKey(&inputs.Queue, up)
	}

	// Climb down.
	if ebiten.IsKeyPressed(down) {
		sys.addUniqueKey(&inputs.Queue, down)
	}

	if inpututil.IsKeyJustReleased(up) {
		sys.addUniqueKey(&inputs.Queue, inputGlobals.RELEASED_VERTICAL_UP)
	}

	if inpututil.IsKeyJustReleased(down) {
		sys.addUniqueKey(&inputs.Queue, inputGlobals.RELEASED_VERTICAL_DOWN)
	}

	if inpututil.IsKeyJustPressed(attackPrimary) {
		sys.addUniqueKey(&inputs.Queue, attackPrimary)
	}
}

func (sys *InputTrackerSystem) trackLeftRightLastPressed() {

	left := inputGlobals.KeyLeft()
	right := inputGlobals.KeyRight()

	if inpututil.IsKeyJustPressed(left) {
		sys.tickLeftKeyLastPressed = sys.scene.Manager.TickHandler.CurrentTick()
	}

	if inpututil.IsKeyJustPressed(right) {
		sys.tickRightKeyLastPressed = sys.scene.Manager.TickHandler.CurrentTick()
	}

	if inpututil.IsKeyJustReleased(left) {
		sys.tickLeftKeyLastPressed = 0
	}

	if inpututil.IsKeyJustReleased(right) {
		sys.tickRightKeyLastPressed = 0
	}

}

func (sys InputTrackerSystem) addUniqueKey(slice *[]ebiten.Key, element ebiten.Key) bool {
	for _, existing := range *slice {
		if existing == element {
			return false // Element is not unique
		}
	}
	*slice = append(*slice, element)
	return true // Element added (unique)
}

func (sys InputTrackerSystem) playerCannotAcceptMovementInputs(world donburi.World, playerState *components.PlayerState) bool {

	isChatActive, _ := systemsUtil.IsChatActive(world)
	return isChatActive ||
		playerState.Combat.IsHit ||
		playerState.Combat.Defeated
}
