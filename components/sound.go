package components

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/yohamta/donburi"
)

var SoundComponent = donburi.NewComponentType[Sound]()
var SoundsMapComponent = donburi.NewComponentType[map[string]*Sound]()
var SoundCharStateMapComponent = donburi.NewComponentType[map[CharState]*Sound]()

type Sound struct {
	Bytes []byte
	*AssetData

	State  *SoundState
	Config *SoundConfig
}

type SoundState struct {
	Player *audio.Player
}

type SoundConfig struct {
	Duration    float64 // Seconds.
	VolumeScale float64
}

func NewSound(duration float64, volumeScale float64) *Sound {

	if volumeScale > 1 {
		log.Fatalf("VolumeScale must be less than to 1.0. VolumeScale: %f", volumeScale)
	}

	if volumeScale < 0 {
		log.Fatalf("VolumeScale must be greater than to 0.0. VolumeScale: %f", volumeScale)
	}

	return &Sound{
		AssetData: &AssetData{},
		State:     &SoundState{},
		Config: &SoundConfig{
			Duration:    duration,
			VolumeScale: volumeScale,
		},
	}
}
