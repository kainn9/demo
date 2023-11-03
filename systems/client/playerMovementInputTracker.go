package clientSystems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	inputConstants "github.com/kainn9/demo/constants/input"
	systemsUtil "github.com/kainn9/demo/systems/util"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type PlayerMovementInputTrackerSystem struct {
	scene *coldBrew.Scene
}

func NewPlayerMovementInputTracker(scene *coldBrew.Scene) *PlayerMovementInputTrackerSystem {
	return &PlayerMovementInputTrackerSystem{
		scene: scene,
	}
}

func (PlayerMovementInputTrackerSystem) Query() *donburi.Query {
	return donburi.NewQuery(
		filter.Contains(components.InputsComponent),
	)
}

func (sys PlayerMovementInputTrackerSystem) Sync(inputsEntity *donburi.Entry) {

	// Block movement inputs unless chat is not active.
	if systemsUtil.IsChatActive(sys.scene.World) {
		return
	}

	inputs := components.InputsComponent.Get(inputsEntity)

	left, right, jump, up, down, interact := inputConstants.ALL_BINDS()

	// Left/Right movement.
	// Else if, is intentional here.
	if ebiten.IsKeyPressed(left) {
		sys.addUniqueKey(&inputs.Queue, left)
	} else if ebiten.IsKeyPressed(right) {
		sys.addUniqueKey(&inputs.Queue, right)
	}

	if !ebiten.IsKeyPressed(left) && !ebiten.IsKeyPressed(right) {
		sys.addUniqueKey(&inputs.Queue, inputConstants.RELEASED_HORIZONTAL)
	}

	// Jump.
	if inpututil.IsKeyJustPressed(jump) && !ebiten.IsKeyPressed(down) {
		sys.addUniqueKey(&inputs.Queue, jump)
	}

	// Phase through platforms.
	if inpututil.IsKeyJustPressed(jump) && ebiten.IsKeyPressed(down) {
		sys.addUniqueKey(&inputs.Queue, inputConstants.COMBO_DOWN_SPACE)
	}

	// Climb up.
	if ebiten.IsKeyPressed(up) {
		sys.addUniqueKey(&inputs.Queue, up)
	}

	// Climb down.
	if ebiten.IsKeyPressed(down) {
		sys.addUniqueKey(&inputs.Queue, down)
	}

	if !ebiten.IsKeyPressed(up) && !ebiten.IsKeyPressed(down) {
		sys.addUniqueKey(&inputs.Queue, inputConstants.RELEASED_VERTICAL)
	}

	// Interact.
	if inpututil.IsKeyJustPressed(interact) {
		sys.addUniqueKey(&inputs.Queue, interact)
	}

}

func (sys PlayerMovementInputTrackerSystem) addUniqueKey(slice *[]ebiten.Key, element ebiten.Key) bool {
	for _, existing := range *slice {
		if existing == element {
			return false // Element is not unique
		}
	}
	*slice = append(*slice, element)
	return true // Element added (unique)
}
