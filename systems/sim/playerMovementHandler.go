package simSystems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kainn9/demo/components"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type PlayerMovementHandler struct{}

func NewPlayerMovementHandler() *PlayerMovementHandler {
	return &PlayerMovementHandler{}
}

func (*PlayerMovementHandler) Query() *donburi.Query {
	return donburi.NewQuery(
		filter.And(
			filter.Contains(components.InputBufferComponent),
			filter.Contains(components.RigidBodyComponent),
		),
	)
}

func (sys *PlayerMovementHandler) Run(_ float64, entry *donburi.Entry) {
	inputBuffer := components.InputBufferComponent.Get(entry)
	playerBody := components.RigidBodyComponent.Get(entry)

	switch {

	case inputBuffer.ActiveInput == ebiten.KeyRight:
		playerBody.X += 3

	case inputBuffer.ActiveInput == ebiten.KeyLeft:
		playerBody.X -= 3

	}

	inputBuffer.ActiveInput = -1

}
