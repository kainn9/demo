package simPlayerSystems

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	sharedStateGlobals "github.com/kainn9/demo/globalConfig/sharedState"
	"github.com/kainn9/demo/queries"
	combatUtil "github.com/kainn9/demo/systems/sim/combatUtil"
	systemsUtil "github.com/kainn9/demo/systems/util"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
	tBokiPhysics "github.com/kainn9/tteokbokki/physics"
	"github.com/yohamta/donburi"
)

type PlayerContactHitDetectorSystem struct {
	scene *coldBrew.Scene
}

func NewPlayerContactHitDetector(scene *coldBrew.Scene) *PlayerContactHitDetectorSystem {
	return &PlayerContactHitDetectorSystem{
		scene: scene,
	}
}

func (PlayerContactHitDetectorSystem) Query() *donburi.Query {
	return queries.NpcQuery
}

func (sys PlayerContactHitDetectorSystem) Run(dt float64, npcEntity *donburi.Entry) {

	world := sys.scene.World
	tickHandler := sys.scene.Manager.TickHandler

	playerEntity := systemsUtil.PlayerEntity(world)
	playerBody := components.RigidBodyComponent.Get(playerEntity)
	playerState := components.PlayerStateComponent.Get(playerEntity)

	if combatUtil.PlayerIsInvincible(playerState) {
		return
	}

	npcState := components.NpcStateComponent.Get(npcEntity)
	npcBody := components.RigidBodyComponent.Get(npcEntity)

	if !npcState.Combat.Hittable || npcState.Combat.Defeated {
		return
	}

	npcId := systemsUtil.ID(npcEntity)

	playerIsCollidingWithNpc, _ := tBokiPhysics.Detector.Detect(playerBody, npcBody, tBokiComponents.ResolverType)
	playerIsNotAlreadyHit := !playerState.Combat.Hits[npcId]

	if playerIsCollidingWithNpc && playerIsNotAlreadyHit {
		sys.handleContactHit(playerBody, npcBody, playerState, tickHandler, npcEntity, playerEntity)
	}

}

func (sys PlayerContactHitDetectorSystem) handleContactHit(
	playerBody, npcBody *tBokiComponents.RigidBody,
	playerState *components.PlayerState,
	tickHandler *coldBrew.TickHandler,
	initiator, target *donburi.Entry,
) {

	combatUtil.CreateHitEntity(
		sys.scene,
		sharedStateGlobals.CHAR_STATE_CONTACT_ATTACK,
		tickHandler.CurrentTick(), // Expire Immediately.
		nil,
		initiator,
		target,
	)

}
