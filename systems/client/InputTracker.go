package clientSystems

import (
	"github.com/hajimehoshi/ebiten/v2"
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
		filter.Contains(components.InputBufferComponent),
	)
}

func (sys *InputTrackerSystem) Sync(entity *donburi.Entry) {
	inputBuffer := components.InputBufferComponent.Get(entity)
	switch {

	case ebiten.IsKeyPressed(ebiten.KeyRight):
		inputBuffer.ActiveInput = ebiten.KeyRight

	case ebiten.IsKeyPressed(ebiten.KeyLeft):
		inputBuffer.ActiveInput = ebiten.KeyLeft
	}

}
