package simNpcSystems

import (
	"log"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	npcGlobals "github.com/kainn9/demo/globalConfig/npc"
	"github.com/kainn9/demo/queries"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
	"github.com/yohamta/donburi"
)

type NpcMeleeAttackHandlerSystem struct {
	scene *coldBrew.Scene
}

func NewNpcMeleeAttackHandler(scene *coldBrew.Scene) *NpcMeleeAttackHandlerSystem {
	return &NpcMeleeAttackHandlerSystem{
		scene: scene,
	}
}

func (sys NpcMeleeAttackHandlerSystem) Query() *donburi.Query {
	return queries.NpcQuery
}

func (sys NpcMeleeAttackHandlerSystem) Run(dt float64, npcEntity *donburi.Entry) {
	world := sys.scene.World
	tickHandler := sys.scene.Manager.TickHandler

	npcState := components.NpcStateComponent.Get(npcEntity)

	if !npcState.Combat.Hittable {
		return
	}

	npcConfig := components.NpcConfigComponent.Get(npcEntity)
	npcBody := components.RigidBodyComponent.Get(npcEntity)

	hitboxesData := components.AttackHitboxConfigComponent.Get(npcEntity)

	if npcState.Combat.CurrentAttack == "" {
		return
	}

	currentAttackData := npcGlobals.NpcAttackDataMap[npcConfig.Name]

	if currentAttackData == nil {
		log.Println("currant attack state:", npcState.Combat.CurrentAttack)
		panic("currentAttackData is nil!")
	}

	ticksSinceAttackStart := tickHandler.TicksSinceNTicks(npcState.Combat.AttackStartTick)

	if ticksSinceAttackStart >= currentAttackData.TotalTickLength {
		sys.endAttack(world, npcState)
		return
	}

	frame := ticksSinceAttackStart / currentAttackData.TicksPerFrame
	sys.processAttack(frame, hitboxesData, npcBody, npcState)

}

func (sys NpcMeleeAttackHandlerSystem) endAttack(world donburi.World, npcState *components.NpcState) {

	// Remove attack hitboxes for the matching attack id/entity.
	sys.removeAttackEntityFromWorld(world, npcState)

	// Update player combat state.
	npcState.Combat.AttackStartTick = -1
	npcState.Combat.CurrentAttack = ""
}

func (sys NpcMeleeAttackHandlerSystem) removeAttackEntityFromWorld(world donburi.World, npcState *components.NpcState) {

	queries.AttackQuery.Each(world, func(attackEntity *donburi.Entry) {

		emptyBoxes := []*tBokiComponents.RigidBody{
			tBokiComponents.NewRigidBodyBox(0, 0, 0, 0, 0, false),
		}

		attackState := components.AttackStateComponent.Get(attackEntity)
		attackId := attackState.ID

		if int(attackId) == npcState.Combat.AttackStartTick {
			components.AttackBoxesComponent.SetValue(attackEntity, emptyBoxes)
			sys.clearAttackFromNpcHits(world, attackId)
			world.Remove(attackEntity.Entity())

		}
	})

}

func (sys NpcMeleeAttackHandlerSystem) clearAttackFromNpcHits(world donburi.World, attackId int) {
	// Remove the attack id from all NPC that were hit by this attack.
	queries.PlayerQuery.Each(world, func(playerEntity *donburi.Entry) {
		playerState := components.PlayerStateComponent.Get(playerEntity)

		if playerState.Combat.Hits[attackId] != 0 {
			delete(playerState.Combat.Hits, attackId)
		}

	})
}

func (sys NpcMeleeAttackHandlerSystem) processAttack(
	frame int,
	hitboxesData *components.AttackHitboxConfig,
	npcBody *tBokiComponents.RigidBody,
	npcState *components.NpcState,
) {

	if frame > len(hitboxesData.Hitboxes)-1 {
		return
	}

	// Generate hitboxes for the current frame.
	// Each frame can have multiple hitboxes to
	// create more complex attack shapes.
	currBoxesData := hitboxesData.Hitboxes[frame]
	hitboxes := make([]*tBokiComponents.RigidBody, 0)

	for _, boxData := range currBoxesData {

		xPos := npcBody.Pos.X + (boxData.OffsetX * npcState.Direction())
		yPos := npcBody.Pos.Y + boxData.OffsetY

		isAngular := boxData.Rotation != 0

		box := tBokiComponents.NewRigidBodyBox(xPos, yPos, boxData.Width, boxData.Height, 0, isAngular)
		box.Rotation = (boxData.Rotation * npcState.Direction())
		box.UpdateVertices()

		hitboxes = append(hitboxes, box)
	}

	// Associate the hitboxes with the attack entity & record the attack id.
	queries.AttackQuery.Each(sys.scene.World, func(attackEntity *donburi.Entry) {
		attackState := components.AttackStateComponent.Get(attackEntity)
		attackId := attackState.ID

		if int(attackId) == npcState.Combat.AttackStartTick {
			components.AttackBoxesComponent.SetValue(attackEntity, hitboxes)
		}
	})

}
