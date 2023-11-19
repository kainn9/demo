package clientNpcSystems

import (
	"log"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	"github.com/kainn9/demo/queries"
	soundUtil "github.com/kainn9/demo/systems/client/util/sound"
	systemsUtil "github.com/kainn9/demo/systems/util"
	"github.com/yohamta/donburi"
)

type NpcHitSoundPlayerSystem struct {
	scene *coldBrew.Scene
}

func (sys NpcHitSoundPlayerSystem) NpcQuery() *donburi.Query {
	return queries.NpcQuery
}

func NewNpcHitSoundPlayer(scene *coldBrew.Scene) *NpcHitSoundPlayerSystem {
	return &NpcHitSoundPlayerSystem{
		scene: scene,
	}
}

func (sys *NpcHitSoundPlayerSystem) Sync(_ *donburi.Entry) {
	world := sys.scene.World
	// playerEntity := systemsUtil.GetPlayerEntity(world)

	sys.NpcQuery().Each(world, func(npcEntity *donburi.Entry) {
		sys.handleHitSounds(npcEntity)
	})
}

func (sys *NpcHitSoundPlayerSystem) handleHitSounds(npcEntity *donburi.Entry) {

	npcState := components.NpcStateComponent.Get(npcEntity)

	if !npcState.Combat.IsHit {
		return
	}

	globalSounds := systemsUtil.GetUISoundsSingletonEntity(sys.scene.World)
	audContext := components.AudioContextComponent.Get(globalSounds)

	playerEntity := systemsUtil.GetPlayerEntity(sys.scene.World)
	playerSoundsMap := components.SoundCharStateMapComponent.Get(playerEntity)

	sound := (*playerSoundsMap)[components.CharState(npcState.Combat.LatestHitAttackName)]
	if sound == nil {
		log.Println("sound is nil")
		log.Println("npcState.Combat.LatestHitAttackName", npcState.Combat.LatestHitAttackName)
		log.Println("playerSoundsMap", playerSoundsMap)
		return
	}

	soundUtil.PlaySound(audContext.Context, sound, sys.scene.Manager.TickHandler)
}
