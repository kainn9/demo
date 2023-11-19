package soundUtil

import (
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	clientGlobals "github.com/kainn9/demo/globalConfig/client"
)

func PlaySound(audContext *audio.Context, sound *components.Sound, tickHandler *coldBrew.TickHandler) {

	if sound.State.StartTick == -1 || tickHandler.CurrentTick() >= sound.State.StartTick+sound.Config.TickLength {
		sound.State.StartTick = tickHandler.CurrentTick()
		sound.State.Player = audContext.NewPlayerFromBytes(sound.Bytes)

		volume := clientGlobals.SOUND_MAX_VOLUME * sound.Config.VolumeScale
		sound.State.Player.SetVolume(volume)
		sound.State.Player.Play()
	}

}

func PauseSound(sound *components.Sound) {

	if sound.State.Player != nil {
		sound.State.Player.Pause()
		return
	}

	//log.Println("sound.State.Player is nil")
}

func ResetSound(sound *components.Sound) {

	if sound.State.Player != nil {
		sound.State.Player.Rewind()
		return
	}

	//log.Println("sound.State.Player is nil")

}
