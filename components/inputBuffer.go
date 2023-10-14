package components

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type InputBuffer struct {
	ActiveInput         ebiten.Key
	BufferedActiveInput ebiten.Key
}

var InputBufferComponent = donburi.NewComponentType[InputBuffer]()

func NewInputBuffer() *InputBuffer {
	return &InputBuffer{
		ActiveInput:         -1,
		BufferedActiveInput: -1,
	}
}
