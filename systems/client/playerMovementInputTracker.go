package clientSystems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kainn9/demo/components"

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

func chatIsActive(inputEntity *donburi.Entry) bool {
	// A bit of a hacky way to do this, but it works.
	// Could also just use the scene to get the world(like we do in other systems)...
	// Not sure which I like better yet.
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
