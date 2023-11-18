package components

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

var InputsComponent = donburi.NewComponentType[Inputs]()

type Inputs struct {
	Queue []ebiten.Key
}

func NewInputs() *Inputs {
	return &Inputs{
		Queue: make([]ebiten.Key, 0),
	}
}
