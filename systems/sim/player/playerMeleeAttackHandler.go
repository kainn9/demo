package simPlayerSystems

import (
	"log"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	playerGlobals "github.com/kainn9/demo/globalConfig/player"
	"github.com/kainn9/demo/queries"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
	"github.com/yohamta/donburi"
)

type PlayerMeleeAttackHandlerSystem struct {
	scene *coldBrew.Scene
}

func NewPlayerMeleeAttackHandler(scene *coldBrew.Scene) *PlayerMeleeAttackHandlerSystem {
	return &PlayerMeleeAttackHandlerSystem{
		scene: scene,
	}
}

func (sys PlayerMeleeAttackHandlerSystem) Query() *donburi.Query {
	return queries.PlayerQuery
}

func (sys PlayerMeleeAttackHandlerSystem) Run(dt float64, playerEntity *donburi.Entry) {
	world := sys.scene.World

	playerState := components.PlayerStateComponent.Get(playerEntity)
	playerBody := components.RigidBodyComponent.Get(playerEntity)
	hitboxesData := components.AttackHitboxConfigComponent.Get(playerEntity)

	tickHandler := sys.scene.Manager.TickHandler

	if playerState.Combat.CurrentAttack == "" {
		return
	}

	currentAttackData := playerGlobals.PlayerAttackDataMap[playerState.Combat.CurrentAttack]

	if currentAttackData == nil {
		log.Println("currant attack state:", playerState.Combat.CurrentAttack)
		panic("currentAttackData is nil!")
	}

	ticksSinceAttackStart := tickHandler.TicksSinceNTicks(playerState.Combat.AttackStartTick)

	if ticksSinceAttackStart >= currentAttackData.TotalTickLength {
		sys.endAttack(world, playerState)
		return
	}

	frame := ticksSinceAttackStart / currentAttackData.TicksPerFrame
	sys.processAttack(frame, hitboxesData, playerBody, playerState)

}

func (sys PlayerMeleeAttackHandlerSystem) endAttack(world donburi.World, playerState *components.PlayerState) {

	// Remove attack hitboxes for the matching attack id/entity.
	sys.removeAttackEntityFromWorld(world, playerState)

	// Update player combat state.
	playerState.Combat.Attacking = false
	playerState.Combat.AttackStartTick = -1
	playerState.Combat.CurrentAttack = ""
}

func (sys PlayerMeleeAttackHandlerSystem) removeAttackEntityFromWorld(world donburi.World, playerState *components.PlayerState) {

	queries.AttackQuery.Each(world, func(attackEntity *donburi.Entry) {

		emptyBoxes := []*tBokiComponents.RigidBody{
			tBokiComponents.NewRigidBodyBox(0, 0, 0, 0, 0, false),
		}

		attackState := components.AttackStateComponent.Get(attackEntity)
		attackId := attackState.ID

		if int(attackId) == playerState.Combat.AttackStartTick {
			components.AttackBoxesComponent.SetValue(attackEntity, emptyBoxes)
			sys.clearAttackFromNpcHits(world, attackId)
			world.Remove(attackEntity.Entity())

		}
	})

}

func (sys PlayerMeleeAttackHandlerSystem) clearAttackFromNpcHits(world donburi.World, attackId int) {
	// Remove the attack id from all NPC that were hit by this attack.
	queries.NpcQuery.Each(world, func(npcEntity *donburi.Entry) {
		npcState := components.NpcStateComponent.Get(npcEntity)

		if npcState.Combat.Hits[attackId] != 0 {
			delete(npcState.Combat.Hits, attackId)
		}

	})
}

func (sys PlayerMeleeAttackHandlerSystem) processAttack(
	frame int,
	hitboxesData *components.AttackHitboxConfig,
	playerBody *tBokiComponents.RigidBody,
	playerState *components.PlayerState,
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

		xPos := playerBody.Pos.X + (boxData.OffsetX * playerState.Direction())
		yPos := playerBody.Pos.Y + boxData.OffsetY

		isAngular := boxData.Rotation != 0

		box := tBokiComponents.NewRigidBodyBox(xPos, yPos, boxData.Width, boxData.Height, 0, isAngular)
		box.Rotation = (boxData.Rotation * playerState.Direction())
		box.UpdateVertices()

		hitboxes = append(hitboxes, box)
	}

	// Associate the hitboxes with the attack entity & record the attack id.
	queries.AttackQuery.Each(sys.scene.World, func(attackEntity *donburi.Entry) {
		attackState := components.AttackStateComponent.Get(attackEntity)
		attackId := attackState.ID

		if int(attackId) == playerState.Combat.AttackStartTick {
			components.AttackBoxesComponent.SetValue(attackEntity, hitboxes)
		}
	})

}
