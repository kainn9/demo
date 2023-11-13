package simNpcSystems

import (
	npcGlobals "github.com/kainn9/demo/globalConfig/npc"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	"github.com/kainn9/demo/queries"
	"github.com/yohamta/donburi"
)

type NpcDefeatedHandlerSystem struct {
	scene *coldBrew.Scene
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

	}

	if !state.Combat.Defeated {
		return
	}

	ticksSinceDefeated := ticksHandler.TicksSinceNTicks(state.Combat.DefeatedStartTick)
	if ticksSinceDefeated > npcGlobals.NPC_DEFEATED_DURATIONS[config.Name] {
		sys.scene.World.Remove(npcEntity.Entity())
	}

}
