package simNpcSystems

import (
	"log"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	"github.com/kainn9/demo/queries"
	systemsUtil "github.com/kainn9/demo/systems/util"
	"github.com/yohamta/donburi"
)

type NpcHitHandlerSystem struct {
	scene *coldBrew.Scene
}

func NewNpcHitHandler(scene *coldBrew.Scene) *NpcHitHandlerSystem {
	return &NpcHitHandlerSystem{
		scene: scene,
	}
}

func (NpcHitHandlerSystem) Query() *donburi.Query {
	return queries.HitQuery
}

func (sys NpcHitHandlerSystem) Run(dt float64, hitEntity *donburi.Entry) {
	hitState := components.HitStateComponent.Get(hitEntity)

	playerIsNotInitiatorOfHit := hitState.Initiator != systemsUtil.PlayerEntity(sys.scene.World)

	if playerIsNotInitiatorOfHit {
		return
	}

	sys.clearExpiredHits(hitEntity, hitState)
	sys.ApplyHitEffects(hitState)

}

func (sys NpcHitHandlerSystem) clearExpiredHits(hitEntity *donburi.Entry, hitState *components.HitState) {

	th := sys.scene.Manager.TickHandler
	hitNotExpired := hitState.EndTick < th.CurrentTick()

	if hitNotExpired {
		return
	}

	sys.scene.World.Remove(hitEntity.Entity())

	if hitState.HitCachingDisabled() {
		return
	}

	id := systemsUtil.ID(hitState.Target)

	delete(hitState.Hits, id)
}

func (sys NpcHitHandlerSystem) ApplyHitEffects(hitState *components.HitState) {

	targetIsInvalid := !systemsUtil.Valid(sys.scene.World, hitState.Target)
	if targetIsInvalid {
		return
	}

	npcEntity := hitState.Target
	npcState := components.NpcStateComponent.Get(npcEntity)

	if npcState.Combat.Defeated {
		return
	}

	id := systemsUtil.ID(hitState.Initiator)
	atkName := hitState.AttackName

	npcState.Combat.Health -= 1
	npcState.Combat.Hits[id] = true
	npcState.Combat.IsHit = true
	npcState.Combat.LastHitTick = sys.scene.Manager.TickHandler.CurrentTick()
	npcState.Combat.LatestHitAttackName = atkName

	log.Println("NPC hit!")
	log.Println("health:", npcState.Combat.Health)

}
