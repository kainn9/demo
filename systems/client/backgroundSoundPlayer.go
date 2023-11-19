package clientSystems

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	"github.com/kainn9/demo/queries"
	soundUtil "github.com/kainn9/demo/systems/client/util/sound"
	systemsUtil "github.com/kainn9/demo/systems/util"

	"github.com/yohamta/donburi"
)

type BackgroundSoundPlayerSystem struct {
	scene                   *coldBrew.Scene
	tickLeftKeyLastPressed  int
	tickRightKeyLastPressed int
}

func NewBackgroundSoundPlayer(scene *coldBrew.Scene) *BackgroundSoundPlayerSystem {
	return &BackgroundSoundPlayerSystem{
		scene:                   scene,
		tickLeftKeyLastPressed:  0,
		tickRightKeyLastPressed: 0,
	}
}

func (BackgroundSoundPlayerSystem) Query() *donburi.Query {
	return queries.BackgroundSoundQuery
}

func (sys *BackgroundSoundPlayerSystem) Sync(soundEntity *donburi.Entry) {

	sound := components.SoundComponent.Get(soundEntity)
	globalSounds := systemsUtil.GetUISoundsSingletonEntity(sys.scene.World)
	context := components.AudioContextComponent.Get(globalSounds)

	soundUtil.PlaySound(context.Context, sound, sys.scene.Manager.TickHandler)
}