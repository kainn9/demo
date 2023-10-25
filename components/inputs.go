package components

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type Inputs struct {
	Queue []ebiten.Key
}

var InputsComponent = donburi.NewComponentType[Inputs]()

func NewInputs() *Inputs {
	return &Inputs{
		Queue: make([]ebiten.Key, 0),
	}
}
