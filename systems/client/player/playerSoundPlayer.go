package clientPlayerSystems

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	"github.com/kainn9/demo/queries"
	soundUtil "github.com/kainn9/demo/systems/client/util/sound"
	systemsUtil "github.com/kainn9/demo/systems/util"

	sharedStateGlobals "github.com/kainn9/demo/globalConfig/sharedState"
	"github.com/yohamta/donburi"
)

type PlayerSoundPlayerSystem struct {
	scene *coldBrew.Scene
}

func (sys PlayerSoundPlayerSystem) Query() *donburi.Query {
	return queries.PlayerQuery
}

func NewPlayerSoundPlayer(scene *coldBrew.Scene) *PlayerSoundPlayerSystem {
	return &PlayerSoundPlayerSystem{
		scene: scene,
	}
}

func (sys *PlayerSoundPlayerSystem) Sync(playerEntity *donburi.Entry) {

	soundMap := components.SoundCharStateMapComponent.Get(playerEntity)
	playerState := components.PlayerStateComponent.Get(playerEntity)

	globalSoundsSingleton := systemsUtil.UISoundsSingletonEntity(sys.scene.World)
	audContext := components.AudioContextComponent.Get(globalSoundsSingleton)

	th := sys.scene.Manager.TickHandler

	if playerState.Combat.IsHit {

		sound := (*soundMap)[sharedStateGlobals.CHAR_STATE_HURT]
		soundUtil.PlaySound(audContext.Context, sound, th)
	}

	if playerState.Collision.OnGround && playerState.Transform.BasicHorizontalMovement {
		sound := (*soundMap)[sharedStateGlobals.CHAR_STATE_RUN]
		soundUtil.PlaySound(audContext.Context, sound, th)
	} else {
		sound := (*soundMap)[sharedStateGlobals.CHAR_STATE_RUN]
		soundUtil.PauseSound(sound)
	}

}
