package simNpcSystems

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	sharedCombatGlobals "github.com/kainn9/demo/globalConfig/sharedCombat"
	"github.com/kainn9/demo/queries"
	"github.com/yohamta/donburi"
)

type NpcClearHitStateHandlerSystem struct {
	scene           *coldBrew.Scene
	DefeatCallbacks []DefeatCallback
}

func NewNpcClearHitStateHandler(scene *coldBrew.Scene) *NpcClearHitStateHandlerSystem {
	return &NpcClearHitStateHandlerSystem{
		scene: scene,
	}
}

func (sys NpcClearHitStateHandlerSystem) Query() *donburi.Query {
	return queries.NpcQuery
}

func (sys NpcClearHitStateHandlerSystem) Run(dt float64, npcEntity *donburi.Entry) {
	npcState := components.NpcStateComponent.Get(npcEntity)

	if !npcState.Combat.Hittable {
		return
	}

	sys.clearHitState(npcState)
}

func (sys NpcClearHitStateHandlerSystem) clearHitState(npcState *components.NpcState) {
	ticksSinceLastHit := sys.scene.Manager.TickHandler.TicksSinceNTicks(npcState.Combat.LastHitTick)

	if ticksSinceLastHit > sharedCombatGlobals.IS_HIT_DURATION_IN_TICKS {
		npcState.Combat.IsHit = false
	}
}
