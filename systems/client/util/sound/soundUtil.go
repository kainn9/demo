package soundUtil

import (
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	clientGlobals "github.com/kainn9/demo/globalConfig/client"
)

func PlaySound(audContext *audio.Context, sound *components.Sound, tickHandler *coldBrew.TickHandler) {

	if sound.State.Player == nil {
		sound.State.Player = audContext.NewPlayerFromBytes(sound.Bytes)
	}

	usingCustomDur := sound.Config.DurationTicks != -1

	if sound.State.Player.IsPlaying() && !usingCustomDur {
		return
	}

	if tickHandler.TicksSinceNTicks(sound.State.StartTick) < int(sound.Config.DurationTicks) &&
		usingCustomDur &&
		sound.State.StartTick != -1 {
		return
	}

	volume := clientGlobals.SOUND_MAX_VOLUME * sound.Config.VolumeScale

	sound.State.Player.SetVolume(volume)
	sound.State.Player.Rewind()
	sound.State.Player.Play()
	sound.State.StartTick = tickHandler.CurrentTick()

}

func PauseSound(sound *components.Sound) {
	if sound.State.Player == nil {
		return
	}

	sound.State.Player.Pause()
}
