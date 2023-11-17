package simPlayerSystems

import (
	"log"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	playerGlobals "github.com/kainn9/demo/globalConfig/player"
	"github.com/kainn9/demo/queries"
	systemsUtil "github.com/kainn9/demo/systems/util"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
	tBokiVec "github.com/kainn9/tteokbokki/math/vec"
	tBokiPhysics "github.com/kainn9/tteokbokki/physics"
	"github.com/yohamta/donburi"
)

type PlayerNpcHitHandlerSystem struct {
	scene *coldBrew.Scene
}

func NewPlayerNpcHitHandler(scene *coldBrew.Scene) *PlayerNpcHitHandlerSystem {
	return &PlayerNpcHitHandlerSystem{
		scene: scene,
	}
}

func (sys PlayerNpcHitHandlerSystem) Query() *donburi.Query {
	return queries.NpcQuery
}

func (sys PlayerNpcHitHandlerSystem) Run(dt float64, npcEntity *donburi.Entry) {
	world := sys.scene.World
	tickHandler := sys.scene.Manager.TickHandler

	npcBody := components.RigidBodyComponent.Get(npcEntity)
	npcState := components.NpcStateComponent.Get(npcEntity)

	playerEntity := systemsUtil.GetPlayerEntity(world)
	playerBody := components.RigidBodyComponent.Get(playerEntity)
	playerState := components.PlayerStateComponent.Get(playerEntity)

	if !npcState.Combat.Hittable || npcState.Combat.Defeated {
		return
	}

	id := int(npcEntity.Entity().Id())

	isColliding, _ := tBokiPhysics.Detector.Detect(playerBody, npcBody, true)

	if isColliding && playerState.Combat.Hits[id] == 0 {
		sys.processAttack(playerBody, npcBody, playerState, tickHandler, id)
	} else {
		delete(playerState.Combat.Hits, id)
	}

	if tickHandler.TicksSinceNTicks(playerState.Combat.LastHitTick) > playerGlobals.PLAYER_HURT_DURATION_TICKS {
		playerState.Combat.Hit = false
		playerState.Combat.LastHitTick = -1
	}

}

func (sys PlayerNpcHitHandlerSystem) processAttack(
	playerBody, npcBody *tBokiComponents.RigidBody,
	playerState *components.PlayerState,
	tickHandler *coldBrew.TickHandler,
	id int,
) {

	if sys.playerNotHittable(playerState) {
		return
	}

	var xFactor = 1.0

	if playerBody.Pos.X < npcBody.Pos.X {
		xFactor = -1
	}

	playerBody.Vel.X = 0
	playerBody.Vel.Y = 0
	tBokiPhysics.Transformer.ApplyImpulseLinear(playerBody, tBokiVec.Vec2{X: 140 * xFactor, Y: -70})

	playerState.Combat.Hit = true
	playerState.Combat.Health -= 1
	playerState.Combat.Hits[id] = id
	playerState.Combat.LastHitTick = tickHandler.CurrentTick()

	// Starting Iframe(handled in separate system).
	playerState.Combat.Invincible = true

	log.Println("player hit, health:", playerState.Combat.Health)

}

func (sys PlayerNpcHitHandlerSystem) playerNotHittable(playerState *components.PlayerState) bool {
	return playerState.Combat.Defeated || playerState.Combat.Invincible
}
