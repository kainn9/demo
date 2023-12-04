package simPlayerSystems

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	npcGlobals "github.com/kainn9/demo/globalConfig/npc"
	sharedStateGlobals "github.com/kainn9/demo/globalConfig/sharedState"
	"github.com/kainn9/demo/queries"
	combatUtil "github.com/kainn9/demo/systems/sim/combatUtil"
	systemsUtil "github.com/kainn9/demo/systems/util"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
	tBokiPhysics "github.com/kainn9/tteokbokki/physics"
	"github.com/yohamta/donburi"
)

type PlayerAttackHitDetectorSystem struct {
	scene *coldBrew.Scene
}

func NewPlayerAttackHitDetector(scene *coldBrew.Scene) *PlayerAttackHitDetectorSystem {
	return &PlayerAttackHitDetectorSystem{
		scene: scene,
	}
}

func (PlayerAttackHitDetectorSystem) Query() *donburi.Query {
	return queries.PlayerQuery
}

func (sys PlayerAttackHitDetectorSystem) Run(dt float64, playerEntity *donburi.Entry) {

	world := sys.scene.World
	playerState := components.PlayerStateComponent.Get(playerEntity)
	playerBody := components.RigidBodyComponent.Get(playerEntity)

	if combatUtil.PlayerIsInvincible(playerState) {
		return
	}

	sys.checkForHit(world, playerEntity, playerState, playerBody)
}

func (sys PlayerAttackHitDetectorSystem) checkForHit(world donburi.World, playerEntity *donburi.Entry, playerState *components.PlayerState, playerBody *tBokiComponents.RigidBody) {
	queries.AttackQuery.Each(world, func(attackEntity *donburi.Entry) {

		attackHitboxes := components.AttackHitboxesComponent.Get(attackEntity)
		attackData := components.AttackDataComponent.Get(attackEntity)

		for _, attackHitbox := range *attackHitboxes {

			if isColliding, _ := tBokiPhysics.Detector.Detect(playerBody, attackHitbox, tBokiComponents.ResolverType); isColliding {
				sys.handleAttackHit(playerEntity, playerState, attackData, sys.scene.Manager.TickHandler)
			}
		}
	})
}

func (sys PlayerAttackHitDetectorSystem) handleAttackHit(
	playerEntity *donburi.Entry,
	playerState *components.PlayerState,
	attackData *components.AttackData,
	tickHandler *coldBrew.TickHandler,
) {

	world := sys.scene.World

	if attackData.Initiator == playerEntity {
		return
	}

	npcEntity := attackData.Initiator

	initiatorIsInvalid := !systemsUtil.Valid(world, npcEntity)
	if initiatorIsInvalid {
		return
	}

	npcConfig := components.NpcConfigComponent.Get(npcEntity)

	npcName := npcConfig.Name

	id := systemsUtil.ID(npcEntity)

	if playerState.Combat.Hits[id] {
		return
	}

	attackTickLength := npcGlobals.NPCAttackDataMaps[npcName][sharedStateGlobals.CHAR_STATE_ATTACK_PRIMARY].TotalTickLength

	combatUtil.CreateHitEntity(
		sys.scene,
		attackData.Name,
		tickHandler.CurrentTick()+attackTickLength,
		playerState.Combat.Hits,
		npcEntity,
		playerEntity,
	)

}
