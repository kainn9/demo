package simNpcSystems

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	"github.com/kainn9/demo/queries"
	combatUtil "github.com/kainn9/demo/systems/sim/combatUtil"
	systemsUtil "github.com/kainn9/demo/systems/util"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
	tBokiPhysics "github.com/kainn9/tteokbokki/physics"
	"github.com/yohamta/donburi"
)

type NpcHitDetectorSystem struct {
	scene *coldBrew.Scene
}

func NewNpcHitDetector(scene *coldBrew.Scene) *NpcHitDetectorSystem {
	return &NpcHitDetectorSystem{
		scene: scene,
	}
}

func (sys NpcHitDetectorSystem) Query() *donburi.Query {
	return queries.NpcQuery
}

func (sys NpcHitDetectorSystem) Run(dt float64, npcEntity *donburi.Entry) {

	world := sys.scene.World
	npcBody := components.RigidBodyComponent.Get(npcEntity)
	npcState := components.NpcStateComponent.Get(npcEntity)

	if !npcState.Combat.Hittable {
		return
	}

	// For each attack entity, check if it's colliding with the npc.
	queries.AttackQuery.Each(world, func(attackEntity *donburi.Entry) {

		attackHitboxes := components.AttackHitboxesComponent.Get(attackEntity)
		attackData := components.AttackDataComponent.Get(attackEntity)

		for _, attackHitbox := range *attackHitboxes {

			if isColliding, _ := tBokiPhysics.Detector.Detect(npcBody, attackHitbox, tBokiComponents.ResolverType); isColliding {
				sys.handleHit(npcEntity, npcState, attackHitbox, attackData)
			}
		}
	})

}

func (sys NpcHitDetectorSystem) handleHit(
	npcEntity *donburi.Entry,
	npcState *components.NpcState,
	attackHitbox *tBokiComponents.RigidBody,
	attackData *components.AttackData,
) {

	attackDoesNotBelongToPlayer := attackData.Initiator != systemsUtil.PlayerEntity(sys.scene.World)

	if attackDoesNotBelongToPlayer {
		return
	}

	id := systemsUtil.ID(attackData.Initiator)

	if npcState.Combat.Hits[id] {
		return
	}
	th := sys.scene.Manager.TickHandler
	combatUtil.CreateHitEntity(
		sys.scene,
		attackData.Name,
		th.CurrentTick(),
		npcState.Combat.Hits,
		attackData.Initiator,
		npcEntity,
	)

}
