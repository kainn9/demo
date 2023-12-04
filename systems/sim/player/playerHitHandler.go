package simPlayerSystems

import (
	"log"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	"github.com/kainn9/demo/queries"
	systemsUtil "github.com/kainn9/demo/systems/util"
	tBokiVec "github.com/kainn9/tteokbokki/math/vec"
	tBokiPhysics "github.com/kainn9/tteokbokki/physics"
	"github.com/yohamta/donburi"
)

type PlayerHitHandlerSystem struct {
	scene *coldBrew.Scene
}

func NewPlayerHitHandler(scene *coldBrew.Scene) *PlayerHitHandlerSystem {
	return &PlayerHitHandlerSystem{
		scene: scene,
	}
}

func (PlayerHitHandlerSystem) Query() *donburi.Query {
	return queries.HitQuery
}

func (sys PlayerHitHandlerSystem) Run(dt float64, hitEntity *donburi.Entry) {

	hitState := components.HitStateComponent.Get(hitEntity)

	playerIsNotTargetOfHit := hitState.Target != systemsUtil.PlayerEntity(sys.scene.World)

	if playerIsNotTargetOfHit {
		return
	}

	sys.clearExpiredHits(hitEntity, hitState)
	sys.ApplyHitEffects(hitEntity, hitState)

}

func (sys PlayerHitHandlerSystem) clearExpiredHits(hitEntity *donburi.Entry, hitState *components.HitState) {

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

func (sys PlayerHitHandlerSystem) ApplyHitEffects(hitEntity *donburi.Entry, hitState *components.HitState) {

	id := systemsUtil.ID(hitState.Target)

	if hitState.HitCachingEnabled() && hitState.Hits[id] {
		return
	}

	sys.handlePlayerHitDisplacement(hitState)
	sys.handlePlayerHitState(hitState)
}

func (sys PlayerHitHandlerSystem) handlePlayerHitDisplacement(hitState *components.HitState) {

	targetOrInitiatorIsInvalid := !systemsUtil.Valid(sys.scene.World, hitState.Target) || !systemsUtil.Valid(sys.scene.World, hitState.Initiator)
	if targetOrInitiatorIsInvalid {
		return
	}

	playerBody := components.RigidBodyComponent.Get(hitState.Target)
	npcBody := components.RigidBodyComponent.Get(hitState.Initiator)

	var direction = 1.0
	if playerBody.Pos.X < npcBody.Pos.X {
		direction = -1
	}

	playerBody.Vel.X = 0
	playerBody.Vel.Y = 0

	tBokiPhysics.Transformer.ApplyImpulseLinear(playerBody, tBokiVec.Vec2{X: 70 * direction, Y: -35})

}

func (sys PlayerHitHandlerSystem) handlePlayerHitState(hitState *components.HitState) {

	playerState := components.PlayerStateComponent.Get(hitState.Target)

	if hitState.HitCachingEnabled() {
		id := systemsUtil.ID(hitState.Initiator)
		playerState.Combat.Hits[id] = true
	}

	playerState.Combat.IsHit = true
	playerState.Combat.Health -= 1
	playerState.Combat.IsInRecoveryIframe = true
	playerState.Combat.LastHitTick = sys.scene.Manager.TickHandler.CurrentTick()

	log.Println("Player hit!")
	log.Println("Health:", playerState.Combat.Health)

}
