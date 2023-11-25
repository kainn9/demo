package soundUtil

import (
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/kainn9/demo/components"
	clientGlobals "github.com/kainn9/demo/globalConfig/client"
)

func PlaySound(audContext *audio.Context, sound *components.Sound) {

	if sound.State.Player == nil {
		sound.State.Player = audContext.NewPlayerFromBytes(sound.Bytes)
	}

	if sound.State.Player.IsPlaying() && sound.Config.Duration != -1 {
		return
	}

	if sound.State.Player.Position().Seconds() >= sound.Config.Duration {
		sound.State.Player.Rewind()
	}

	volume := clientGlobals.SOUND_MAX_VOLUME * sound.Config.VolumeScale
	sound.State.Player.SetVolume(volume)
	sound.State.Player.Play()

}

func PauseSound(sound *components.Sound) {
	if sound.State.Player == nil {
		return
	}

	sound.State.Player.Pause()
}
