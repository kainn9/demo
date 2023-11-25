package simNpcSystems

import (
	"log"

	"github.com/kainn9/coldBrew"
	components "github.com/kainn9/demo/components"
	npcGlobals "github.com/kainn9/demo/globalConfig/npc"
	"github.com/kainn9/demo/queries"
	combatUtil "github.com/kainn9/demo/systems/sim/combatUtil"
	systemsUtil "github.com/kainn9/demo/systems/util"
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

	npcState := components.NpcStateComponent.Get(npcEntity)

	if combatUtil.NpcIsInvincible(npcState) || npcState.Combat.CurrentAttack == "" {
		return
	}

	world := sys.scene.World
	tickHandler := sys.scene.Manager.TickHandler

	npcConfig := components.NpcConfigComponent.Get(npcEntity)
	npcBody := components.RigidBodyComponent.Get(npcEntity)

	currentAttackData := npcGlobals.NPCAttackDataMaps[npcConfig.Name][npcState.Combat.CurrentAttack]

	if currentAttackData == nil {
		log.Println("curr attack state:", npcState.Combat.CurrentAttack)
		panic("currentAttackData is nil!")
	}

	ticksSinceAttackStart := tickHandler.TicksSinceNTicks(npcState.Combat.AttackStartTick)
	attackIsFinished := ticksSinceAttackStart > currentAttackData.TotalTickLength

	if attackIsFinished {
		combatUtil.RemoveAttackEntity(world, systemsUtil.ID(npcEntity))
		sys.clearAttackState(npcState)
		return
	}

	currentAttackFrame := ticksSinceAttackStart / currentAttackData.TicksPerFrame
	sys.processAttack(currentAttackFrame, npcBody, npcConfig, npcState, npcEntity)

}

func (sys NpcMeleeAttackHandlerSystem) clearAttackState(npcState *components.NpcState) {
	npcState.Combat.AttackStartTick = -1
	npcState.Combat.CurrentAttack = ""
}

func (sys NpcMeleeAttackHandlerSystem) processAttack(
	frame int,
	npcBody *tBokiComponents.RigidBody,
	npcConfig *components.NpcConfig,
	npcState *components.NpcState,
	npcEntity *donburi.Entry,
) {

	attackHitboxesData, ok := npcGlobals.NPCAttackHitboxesDataMaps[npcConfig.Name][npcState.Combat.CurrentAttack]
	if !ok {
		log.Println("key:", npcState.Combat.CurrentAttack)
		panic("attackHitboxesData is nil!")
	}

	if frame > len(attackHitboxesData.Hitboxes)-1 {
		return
	}

	currentFrameHitboxesData := attackHitboxesData.Hitboxes[frame]
	currentFrameHitboxes := make([]*tBokiComponents.RigidBody, 0)

	for _, boxData := range currentFrameHitboxesData {

		xPos := npcBody.Pos.X + (boxData.OffsetX * npcState.Direction())
		yPos := npcBody.Pos.Y + boxData.OffsetY

		isAngular := boxData.Rotation != 0

		box := tBokiComponents.NewRigidBodyBox(xPos, yPos, boxData.Width, boxData.Height, 0, isAngular)
		box.Rotation = (boxData.Rotation * npcState.Direction())
		box.UpdateVertices()

		currentFrameHitboxes = append(currentFrameHitboxes, box)
	}

	queries.AttackQuery.Each(sys.scene.World, func(attackEntity *donburi.Entry) {
		attackState := components.AttackDataComponent.Get(attackEntity)

		if attackState.Initiator == npcEntity {
			components.AttackHitboxesComponent.SetValue(attackEntity, currentFrameHitboxes)
		}
	})

}
