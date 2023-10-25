package clientSystems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kainn9/demo/components"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type InputTrackerSystem struct{}

func NewInputTracker() *InputTrackerSystem {
	return &InputTrackerSystem{}
}

func (*InputTrackerSystem) Query() *donburi.Query {
	return donburi.NewQuery(
		filter.Contains(components.InputsComponent),
	)
}

func (sys *InputTrackerSystem) Sync(entity *donburi.Entry) {
	inputs := components.InputsComponent.Get(entity)

	// Its better to use else if for the horizontal
	// movement keys to avoid weird behavior when players
	// are pressing left/right inputs at the same time.
	// We could also make left/right considered not unique to
	// each other, but this honestly works fine and is simple.
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		addUniqueKey(&inputs.Queue, ebiten.KeyRight)
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		addUniqueKey(&inputs.Queue, ebiten.KeyLeft)
	}

	// For the space key, we want to add it to the queue
	// regardless of the horizontal movement keys.
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		addUniqueKey(&inputs.Queue, ebiten.KeySpace)
	}

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
