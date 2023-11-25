package components

import (
	"log"

	clientGlobals "github.com/kainn9/demo/globalConfig/client"

	audio "github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/yohamta/donburi"
)

var AudioContextComponent = donburi.NewComponentType[AudioContext]()

type AudioContext struct {
	Context *audio.Context
}

func NewAudioContext() *AudioContext {
	log.Println("Creating new AudioContext(only one should exist at a time!).")
	return &AudioContext{
		Context: audio.NewContext(clientGlobals.SOUND_SAMPLE_RATE),
	}
}
