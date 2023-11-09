package simNpcSystems

import (
	"log"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	sharedCombatGlobals "github.com/kainn9/demo/globalConfig/sharedCombat"
	"github.com/kainn9/demo/queries"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
	tBokiPhysics "github.com/kainn9/tteokbokki/physics"
	"github.com/yohamta/donburi"
)

type NpcAttackedHandlerSystem struct {
	scene *coldBrew.Scene
}

func NewNpcAttackedHandler(scene *coldBrew.Scene) *NpcAttackedHandlerSystem {
	return &NpcAttackedHandlerSystem{
		scene: scene,
	}
}

func (sys NpcAttackedHandlerSystem) Query() *donburi.Query {
	return queries.NpcQuery
}

func (sys NpcAttackedHandlerSystem) Run(dt float64, npcEntity *donburi.Entry) {

	world := sys.scene.World
	npcBody := components.RigidBodyComponent.Get(npcEntity)
	npcState := components.NpcStateComponent.Get(npcEntity)

	if !npcState.Combat.Hittable {
		return
	}

	ticksSinceLastHit := sys.scene.Manager.TickHandler.TicksSinceNTicks(npcState.Combat.LastHitTick)

	if ticksSinceLastHit > sharedCombatGlobals.IS_HIT_DURATION_IN_TICKS {
		npcState.Combat.IsHit = false
	}

	if npcState.Combat.Defeated {
		return
	}

	// For each attack entity, check if it's colliding with the npc.
	queries.AttackQuery.Each(world, func(attackEntity *donburi.Entry) {

		attackHitboxes := components.AttackBoxesComponent.Get(attackEntity)
		attackState := components.AttackStateComponent.Get(attackEntity)
		id := attackState.ID

		for _, attackHitbox := range *attackHitboxes {
			if isColliding, _ := tBokiPhysics.Detector.Detect(npcBody, attackHitbox, true); isColliding {
				sys.handleHit(*npcState, attackHitbox, id)
			}
		}
	})

}

func (sys NpcAttackedHandlerSystem) handleHit(npcState components.NpcState, attackHitbox *tBokiComponents.RigidBody, id int) {

	if npcState.Combat.Hits[id] != 0 {
		return
	}

	npcState.Combat.Health -= 1
	npcState.Combat.Hits[id] = id
	npcState.Combat.IsHit = true
	npcState.Combat.LastHitTick = sys.scene.Manager.TickHandler.CurrentTick()

	log.Println("npc hit! health:", npcState.Combat.Health)
}
