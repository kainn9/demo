package clientSystems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kainn9/demo/components"
	"github.com/kainn9/demo/constants"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type PlayerMovementInputTrackerSystem struct{}

func NewPlayerMovementInputTracker() *PlayerMovementInputTrackerSystem {
	return &PlayerMovementInputTrackerSystem{}
}

func (*PlayerMovementInputTrackerSystem) Query() *donburi.Query {
	return donburi.NewQuery(
		filter.Contains(components.InputsComponent),
	)
}

func (sys *PlayerMovementInputTrackerSystem) Sync(entity *donburi.Entry) {

	// Block movement inputs unless chat is not active.
	if chatIsActive(entity) {
		return
	}

	inputs := components.InputsComponent.Get(entity)

	left, right, jump, up, down, interact := constants.AllBinds()

	// Left/Right movement.
	// Else if, is intentional here.
	if ebiten.IsKeyPressed(left) {
		addUniqueKey(&inputs.Queue, left)
	} else if ebiten.IsKeyPressed(right) {
		addUniqueKey(&inputs.Queue, right)
	}

	if !ebiten.IsKeyPressed(left) && !ebiten.IsKeyPressed(right) {
		addUniqueKey(&inputs.Queue, constants.RELEASED_HORIZONTAL)
	}

	// Jump.
	if inpututil.IsKeyJustPressed(jump) && !ebiten.IsKeyPressed(down) {
		addUniqueKey(&inputs.Queue, jump)
	}

	// Phase through platforms.
	if inpututil.IsKeyJustPressed(jump) && ebiten.IsKeyPressed(down) {
		addUniqueKey(&inputs.Queue, constants.COMBO_DOWN_SPACE)
	}

	// Climb up.
	if ebiten.IsKeyPressed(up) {
		addUniqueKey(&inputs.Queue, up)
	}

	// Climb down.
	if ebiten.IsKeyPressed(down) {
		addUniqueKey(&inputs.Queue, down)
	}

	// Interact.
	if inpututil.IsKeyJustPressed(interact) {
		addUniqueKey(&inputs.Queue, interact)
	}

}

func chatIsActive(inputEntity *donburi.Entry) bool {
	world := inputEntity.World

	var isChatActive bool

	query := donburi.NewQuery(
		filter.Contains(components.ChatStateComponent),
	)

	query.Each(world, func(chatEntity *donburi.Entry) {

		config := components.ChatStateComponent.Get(chatEntity)
		if config.Active {
			isChatActive = true
		}
	})

	return isChatActive
}

func addUniqueKey(slice *[]ebiten.Key, element ebiten.Key) bool {
	for _, existing := range *slice {
		if existing == element {
			return false // Element is not unique
		}
	}
	*slice = append(*slice, element)
	return true // Element added (unique)
}
