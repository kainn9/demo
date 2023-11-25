package simNpcSystems

import (
	npcGlobals "github.com/kainn9/demo/globalConfig/npc"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	"github.com/kainn9/demo/queries"
	"github.com/yohamta/donburi"
)

type NpcDefeatedHandlerSystem struct {
	scene           *coldBrew.Scene
	DefeatCallbacks []DefeatCallback
}

type DefeatCallback interface {
	Npc() *donburi.Entry
	OnDefeat(scene *coldBrew.Scene, npcEntity *donburi.Entry)
}

func NewNpcDefeatedHandler(scene *coldBrew.Scene) *NpcDefeatedHandlerSystem {
	return &NpcDefeatedHandlerSystem{
		scene: scene,
	}
}

func (sys NpcDefeatedHandlerSystem) Query() *donburi.Query {
	return queries.NpcQuery
}

func (sys NpcDefeatedHandlerSystem) Run(dt float64, npcEntity *donburi.Entry) {
	ticksHandler := sys.scene.Manager.TickHandler

	state := components.NpcStateComponent.Get(npcEntity)
	body := components.RigidBodyComponent.Get(npcEntity)
	config := components.NpcConfigComponent.Get(npcEntity)

	if state.Combat.Health == 0 && !state.Combat.Defeated {
		state.Combat.Defeated = true
		state.Combat.DefeatedStartTick = ticksHandler.CurrentTick()
		body.Vel.X = 0

		sys.handleCallbacks(npcEntity)
	}

	if !state.Combat.Defeated {
		return
	}

	ticksSinceDefeated := ticksHandler.TicksSinceNTicks(state.Combat.DefeatedStartTick)
	if ticksSinceDefeated > npcGlobals.NPC_DEFEATED_DURATIONS[config.Name] {
		sys.scene.World.Remove(npcEntity.Entity())
	}

}

func (sys *NpcDefeatedHandlerSystem) handleCallbacks(npcEntity *donburi.Entry) {
	for _, callback := range sys.DefeatCallbacks {

		if callback.Npc() == npcEntity {
			callback.OnDefeat(sys.scene, callback.Npc())
		}

	}
}
